package main

import "fmt"
import "strings"
import "strconv"

type DistanceMap map[string]map[string]int;

func (dm *DistanceMap) Print() {
  for startName, rest := range *dm {
    for finishName, distance := range rest {
      fmt.Println(startName + " -> " + finishName + " = " + strconv.Itoa(distance));
    }
  }
}

type Entity struct {
  Track string;
  Name string;
}

type Vertex struct {
  value *Entity;
  neighbors map[string]*Vertex;
}

func (parent *Vertex) AddNeighbor(neighbor *Vertex) {
  parent.neighbors[neighbor.value.Name] = neighbor;
}

type Graph struct {
  vertices map[string]*Vertex;
}

func NewGraph() *Graph {
  return &Graph{vertices: make(map[string]*Vertex)};
}

func (g *Graph) AddVertex(value *Entity) {
  if _, ok := g.vertices[value.Name]; !ok {
    g.vertices[value.Name] = &Vertex{value: value, neighbors: make(map[string]*Vertex)};
  }
}

func (g *Graph) AddEdge(value1 *Entity, value2 *Entity) {
  g.vertices[value1.Name].AddNeighbor(g.vertices[value2.Name]);
}

func (g *Graph) ShortestPath(source, target *Entity) []*Entity {
  distances := make(map[string]int);
  previouses := make(map[string]*Vertex);

  for name, _ := range g.vertices {
    previouses[name] = nil;
  }

  distances[source.Name] = 0;
  verts := make(map[string]*Vertex);

  // copy vertex map
  for name, vertex := range g.vertices {
    verts[name] = vertex;
  }

  var nearestVertex *Vertex;

  for len(verts) > 0 {
    nearestVertex = nil;

    for name, v := range verts {
      if distance, ok := distances[name]; ok {
        if nearestVertex == nil || distance < distances[nearestVertex.value.Name] {
          nearestVertex = v;
        }
      }
    }

    if nearestVertex == nil {
      break;
    }

    if _, ok := distances[nearestVertex.value.Name]; !ok {
      break;
    }

    if target != nil && nearestVertex.value == target {
      return g.composePath(target, distances[target.Name], previouses);
    }

    alt := distances[nearestVertex.value.Name] + 1;

    for name, _ := range nearestVertex.neighbors {
      if _, ok := distances[name]; !ok || alt < distances[name] {
        distances[name] = alt;
      }

      previouses[name] = nearestVertex;
    }

    delete(verts, nearestVertex.value.Name);
  }

  return make([]*Entity, 0, 0);
}

func (g *Graph) composePath(target *Entity, distance int, previouses map[string]*Vertex) []*Entity {
  result := make([]*Entity, distance);

  for i := distance - 1; i >= 0; i -= 1 {
    result[i] = target;
    target = previouses[target.Name].value;
  }

  return result;
}

func (g *Graph) GetDistanceMap() DistanceMap {
  distances := make(DistanceMap);

  for fromName, fromVertex := range g.vertices {
    for toName, toVertex := range g.vertices {
      if fromVertex != toVertex && !strings.Contains(fromName, "track") && !strings.Contains(toName, "track") {
        if distances[fromName] == nil {
          distances[fromName] = make(map[string]int)
        }

        shortestPath := g.ShortestPath(fromVertex.value, toVertex.value);
        distances[fromName][toName] = len(shortestPath);
      }
    }
  }

  return distances;
}
