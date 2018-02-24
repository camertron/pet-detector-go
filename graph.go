package main

type Vertex struct {
    id Entity
    dist int
    arcs map[Entity]int        // arcs[vertex id] = weight
}

type Candidates []Vertex
func (h Candidates) Len() int           { return len(h) }
func (h Candidates) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h Candidates) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Candidates) IsEmpty() bool      { return len(h) == 0 }

func (h *Candidates) Push(v Vertex) {
    var changed bool
    old := *h
    updated := *h
    // insert Vertex at the correct position (keyed by distance)
    for i, w := range old {
        if v.id == w.id {
            if changed {
                if i + 1 < len(updated) {
                    updated = append(updated[:i], updated[i+1:]...)
                } else {
                    updated = updated[:i]
                }
            } else if v.dist < w.dist {
                updated[i] = v
            }
            changed = true
        } else if v.dist < w.dist {
            changed = true
            updated = append(old[:i], v)
            updated = append(updated, w)
            updated = append(updated, old[i + 1:]...)
        }
    }
    if !changed {
        updated = append(old, v)
    }
    *h = updated
}

func (h *Candidates) Pop() (v Vertex) {
    old := *h
    v = old[0]
    *h = old[1:]
    return
}

type Graph struct {
    visited map[Entity]bool
    vertices map[Entity]Vertex
}

func NewGraph(vs map[Entity]Vertex) *Graph {
    g := new(Graph)
    g.visited = make(map[Entity]bool)
    g.vertices = make(map[Entity]Vertex)
    for i, v := range vs {
        v.dist = 1000000
        g.vertices[i] = v
    }
    return g
}

func (g *Graph) Len() int { return len(g.vertices)  }
func (g *Graph) visit(v Entity) { g.visited[v] = true }

func (g *Graph) ShortestPath(src, dest Entity) (x int) {
    for k := range g.visited {
        delete(g.visited, k)
    }

    for _, v := range g.vertices {
        v.dist = 1000000
    }

    g.visit(src)
    v := g.vertices[src]
    h := make(Candidates, len(v.arcs))
    // initialize the heap with out edges from src
    for id, y := range v.arcs {
        v := g.vertices[id]
        // update the vertices being pointed to with the distance.
        v.dist = y
        g.vertices[id] = v
        h.Push(v)
    }
    for src != dest {
        if h.IsEmpty() {
            return 1000000
        }
        v = h.Pop()
        src = v.id
        if g.visited[src] {
            continue
        }
        g.visit(src)
        for w, d := range v.arcs {
            if g.visited[w] {
                continue
            }
            c := g.vertices[w]
            distance := d + v.dist
            if distance < c.dist {
                c.dist = distance
                g.vertices[w] = c
            }
            h.Push(c)
        }
    }
    v = g.vertices[dest]
    return v.dist
}
