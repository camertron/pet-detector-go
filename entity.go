package main

import "strings"

type Entity struct {
  Track string   // top, bottom, etc
  Name string    // human-readable name
  R int          // i.e. Y
  C int          // i.e. X
  Abbrev string  // abbreviation from game result data
  baseName string
}

func (entity *Entity) IsPet() bool {
  return strings.Contains(entity.Name, "pet")
}

func (entity *Entity) IsHouse() bool {
  return strings.Contains(entity.Name, "house")
}

func (entity *Entity) IsCar() bool {
  return entity.Name == "car"
}

func (entity *Entity) GetBaseName() string {
  if entity.baseName == "" {
    if entity.IsHouse() {
      entity.baseName = strings.ToLower(strings.Replace(entity.Name, "house", "", 1))
    } else if entity.IsPet() {
      entity.baseName = strings.ToLower(strings.Replace(entity.Name, "pet", "", 1))
    } else {
      entity.baseName = entity.Name
    }
  }

  return entity.baseName
}
