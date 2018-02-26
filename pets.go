package main

import "fmt"
import "encoding/json"

var data = `[[{"track":"bottom right","name":"houseCockatiel"},{"track":"bottom left right","name":"petSiamese"},{"track":"bottom left right","name":"track-0"},{"track":"bottom left","name":"track-1"}],[{"track":"top bottom right","name":"petTurtle"},{"track":"top left right","name":"petDacshund"},{"track":"top bottom left","name":"houseHedgehog"},{"track":"top bottom","name":"track-2"}],[{"track":"top bottom","name":"track-3"},{"track":"bottom right","name":"petTabby"},{"track":"top bottom left right","name":"petHusky"},{"track":"top bottom left","name":"petHedgehog"}],[{"track":"top bottom","name":"houseTabby"},{"track":"top bottom right","name":"car"},{"track":"top left right","name":"track-4"},{"track":"top bottom left","name":"track-5"}],[{"track":"top bottom","name":"petCockatiel"},{"track":"top right","name":"track-6"},{"track":"bottom left right","name":"track-7"},{"track":"top bottom left","name":"houseTurtle"}],[{"track":"top right","name":"houseSiamese"},{"track":"left right","name":"houseDacshund"},{"track":"top left right","name":"track-8"},{"track":"top left","name":"houseHusky"}]]`;

const MAX_CAR_CAPACITY = 4

func main() {
  var matrix Matrix;
  json.Unmarshal([]byte(data), &matrix);

  result := solve(&matrix, 25);

  for _, entity := range result {
    fmt.Println(entity.Name);
  }
}
