package main

import "strings"

type Matrix [][]Entity;

func (mRef *Matrix) ToGraph() *Graph {
  matrix := *mRef;
  graph := NewGraph();

  for r := range matrix {
    for c := range matrix[r] {
      graph.AddVertex(&matrix[r][c])
      matrix.addLeftNeighbors(graph, &matrix[r][c], c, r);
      matrix.addRightNeighbors(graph, &matrix[r][c], c, r);
      matrix.addTopNeighbors(graph, &matrix[r][c], c, r);
      matrix.addBottomNeighbors(graph, &matrix[r][c], c, r);
    }
  }

  return graph;
}

func (m *Matrix) GetWidth() int {
  return len((*m)[0]);
}

func (m *Matrix) GetHeight() int {
  return len(*m);
}

func (mRef *Matrix) addTopNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "top") {
    neighbor := m.topNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) addBottomNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "bottom") {
    neighbor := m.bottomNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) addLeftNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "left") {
    neighbor := m.leftNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) addRightNeighbors(graph *Graph, entity *Entity, x, y int) {
  var m = *mRef;

  if strings.Contains(entity.Track, "right") {
    neighbor := m.rightNeighbor(x, y);

    if neighbor != nil {
      graph.AddVertex(neighbor);
      graph.AddEdge(entity, neighbor);
    }
  }
}

func (mRef *Matrix) topNeighbor(x, y int) *Entity {
  var m = *mRef;

  if y > 0 {
    return &(m[y - 1][x]);
  }

  return nil;
}

func (mRef *Matrix) bottomNeighbor(x, y int) *Entity {
  var m = *mRef;

  if y < m.GetHeight() - 1 {
    return &(m[y + 1][x]);
  }

  return nil;
}

func (mRef *Matrix) leftNeighbor(x, y int) *Entity {
  var m = *mRef;

  if x > 0 {
    return &(m[y][x - 1]);
  }

  return nil;
}

func (mRef *Matrix) rightNeighbor(x, y int) *Entity {
  var m = *mRef;

  if x < m.GetWidth() - 1 {
    return &(m[y][x + 1]);
  }

  return nil;
}
