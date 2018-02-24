package main

import "fmt"
import "strings"
import "strconv"
// import "github.com/davecgh/go-spew/spew"
import "encoding/json"

var data = `[[{"track":"bottom right","name":"houseCockatiel"},{"track":"bottom left right","name":"petSiamese"},{"track":"bottom left right","name":""},{"track":"bottom left","name":""}],[{"track":"top bottom right","name":"petTurtle"},{"track":"top left right","name":"petDacshund"},{"track":"top bottom left","name":"houseHedgehog"},{"track":"top bottom","name":""}],[{"track":"top bottom","name":""},{"track":"bottom right","name":"petTabby"},{"track":"top bottom left right","name":"petHusky"},{"track":"top bottom left","name":"petHedgehog"}],[{"track":"top bottom","name":"houseTabby"},{"track":"top bottom right","name":"car"},{"track":"top left right","name":""},{"track":"top bottom left","name":""}],[{"track":"top bottom","name":"petCockatiel"},{"track":"top right","name":""},{"track":"bottom left right","name":""},{"track":"top bottom left","name":"houseTurtle"}],[{"track":"top right","name":"houseSiamese"},{"track":"left right","name":"houseDacshund"},{"track":"top left right","name":""},{"track":"top left","name":"houseHusky"}]]`;

type Entity struct {
  Track string;
  Name string;
}

type Matrix [][]Entity;

func main() {
  var matrix Matrix;
  var vertices = make(map[Entity]Vertex);
  json.Unmarshal([]byte(data), &matrix);

  for r := range matrix {
    for c := range matrix[r] {
      vertices[matrix[r][c]] = Vertex{id: matrix[r][c], dist: 0, arcs: make(map[Entity]int)};
    }
  }

  for r := range matrix {
    for c := range matrix[r] {
      matrix.AddTopNeighbors(vertices, &matrix[r][c], c, r);
      matrix.AddBottomNeighbors(vertices, &matrix[r][c], c, r);
      matrix.AddLeftNeighbors(vertices, &matrix[r][c], c, r);
      matrix.AddRightNeighbors(vertices, &matrix[r][c], c, r);
    }
  }

  var graph = NewGraph(vertices);
  var distances = make(map[Entity]map[Entity]int);

  for fromEntity, _ := range graph.vertices {
    for toEntity, _ := range graph.vertices {
      if fromEntity != toEntity && fromEntity.Name != "" && toEntity.Name != "" {
        if distances[fromEntity] == nil {
          distances[fromEntity] = make(map[Entity]int)
        }

        var distance = graph.ShortestPath(fromEntity, toEntity)
        fmt.Println(fromEntity.Name + " -> " + toEntity.Name + " = " + strconv.Itoa(distance));
        distances[fromEntity][toEntity] = distance;
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

func (mRef *Matrix) AddTopNeighbors(vertices map[Entity]Vertex, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "top") {
    neighbor := m.TopNeighbor(x, y);

    if neighbor != nil {
      vertices[*entity].arcs[*neighbor] = 1;
    }
  }
}

func (mRef *Matrix) AddBottomNeighbors(vertices map[Entity]Vertex, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "bottom") {
    neighbor := m.BottomNeighbor(x, y);

    if neighbor != nil {
      vertices[*entity].arcs[*neighbor] = 1;
    }
  }
}

func (mRef *Matrix) AddLeftNeighbors(vertices map[Entity]Vertex, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "left") {
    neighbor := m.LeftNeighbor(x, y);

    if neighbor != nil {
      vertices[*entity].arcs[*neighbor] = 1;
    }
  }
}

func (mRef *Matrix) AddRightNeighbors(vertices map[Entity]Vertex, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "right") {
    neighbor := m.RightNeighbor(x, y);

    if neighbor != nil {
      vertices[*entity].arcs[*neighbor] = 1;
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
