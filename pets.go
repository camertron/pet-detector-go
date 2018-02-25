package main

import "fmt"
import "strings"
import "strconv"
// import "github.com/davecgh/go-spew/spew"
import "encoding/json"

var data = `[[{"track":"bottom right","name":"houseCockatiel"},{"track":"bottom left right","name":"petSiamese"},{"track":"bottom left right","name":"track-0"},{"track":"bottom left","name":"track-1"}],[{"track":"top bottom right","name":"petTurtle"},{"track":"top left right","name":"petDacshund"},{"track":"top bottom left","name":"houseHedgehog"},{"track":"top bottom","name":"track-2"}],[{"track":"top bottom","name":"track-3"},{"track":"bottom right","name":"petTabby"},{"track":"top bottom left right","name":"petHusky"},{"track":"top bottom left","name":"petHedgehog"}],[{"track":"top bottom","name":"houseTabby"},{"track":"top bottom right","name":"car"},{"track":"top left right","name":"track-4"},{"track":"top bottom left","name":"track-5"}],[{"track":"top bottom","name":"petCockatiel"},{"track":"top right","name":"track-6"},{"track":"bottom left right","name":"track-7"},{"track":"top bottom left","name":"houseTurtle"}],[{"track":"top right","name":"houseSiamese"},{"track":"left right","name":"houseDacshund"},{"track":"top left right","name":"track-8"},{"track":"top left","name":"houseHusky"}]]`;

type Entity struct {
  Track string;
  Name string;
}

type Matrix [][]Entity;

func main() {
  var matrix Matrix;
  json.Unmarshal([]byte(data), &matrix);

  graph := NewGraph();

  for r := range matrix {
    for c := range matrix[r] {
      graph.AddVertex(&matrix[r][c])
      matrix.AddLeftNeighbors(graph, &matrix[r][c], c, r);
      matrix.AddRightNeighbors(graph, &matrix[r][c], c, r);
      matrix.AddTopNeighbors(graph, &matrix[r][c], c, r);
      matrix.AddBottomNeighbors(graph, &matrix[r][c], c, r);
    }
  }

  var distances = make(map[string]map[string]int);

  for fromName, fromVertex := range graph.vertices {
    for toName, toVertex := range graph.vertices {
      if fromVertex != toVertex && !strings.Contains(fromName, "track") && !strings.Contains(toName, "track") {
        if distances[fromName] == nil {
          distances[fromName] = make(map[string]int)
        }

        shortestPath := graph.ShortestPath(fromVertex.value, toVertex.value);
        distance := len(shortestPath);
        fmt.Println(fromName + " -> " + toName + " = " + strconv.Itoa(distance));
        distances[fromName][toName] = distance;
      }
    }
  }

  // spew.Dump(distances);
}

func (m *Matrix) width() int {
  return len((*m)[0]);
}

func (m *Matrix) height() int {
  return len(*m);
}

func (mRef *Matrix) AddTopNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "top") {
    neighbor := m.TopNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) AddBottomNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "bottom") {
    neighbor := m.BottomNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) AddLeftNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "left") {
    neighbor := m.LeftNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) AddRightNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "right") {
    neighbor := m.RightNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) TopNeighbor(x, y int) *Entity {
  var m = *mRef;

  if y > 0 {
    return &(m[y - 1][x]);
  }

  return nil;
}

func (mRef *Matrix) BottomNeighbor(x, y int) *Entity {
  var m = *mRef;

  if y < m.height() - 1 {
    return &(m[y + 1][x]);
  }

  return nil;
}

func (mRef *Matrix) LeftNeighbor(x, y int) *Entity {
  var m = *mRef;

  if x > 0 {
    return &(m[y][x - 1]);
  }

  return nil;
}

func (mRef *Matrix) RightNeighbor(x, y int) *Entity {
  var m = *mRef;

  if x < m.width() - 1 {
    return &(m[y][x + 1]);
  }

  return nil;
}
