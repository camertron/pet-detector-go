package main

import "strings"
import "strconv"
import "regexp"

var animals = []string{
  "Husky", "Dachsund", "Turtle", "Ferret", "Tabby", "Siamese",
  "Hedgehog", "Cockatiel", "Chameleon", "Pug", "Rabbit",
}

func ParseLevel(levelData string) *Matrix {
  rows := strings.Split(levelData, ";")
  matrix := NewMatrix((len(rows) / 2) + 1, (len(rows[0]) / 2) + 1)
  entityAbbrevMap := make(map[string]string)
  carRe, _ := regexp.Compile("(\\d+)")
  trackCounter := 0

  for r := 0; r < len(rows); r += 2 {
    for c := 0; c < len(rows[r]); c += 2 {
      track := make([]string, 0, 4)
      entityAbbrev := string([]rune(rows[r])[c])
      entityAbbrevLower := strings.ToLower(entityAbbrev)
      var entityName string;

      if entityAbbrev == "." {
        entityName = "track-" + strconv.Itoa(trackCounter)
        trackCounter ++
      } else if carRe.MatchString(entityAbbrev) {
        // the match here is a number indicating the car's capacity, which is always 4
        // even in the earlier levels
        entityName = "car"
      } else {
        if _, ok := entityAbbrevMap[entityAbbrevLower]; !ok {
          // use the next available animal
          entityAbbrevMap[entityAbbrevLower] = animals[len(entityAbbrevMap)]
        }

        animal := entityAbbrevMap[entityAbbrevLower]
        var entityType string

        // if the entity is equal to the lowercased version of itself, then
        // that means it's a house
        if entityAbbrev == entityAbbrevLower {
          entityType = "house"
        } else {
          entityType = "pet"
        }

        entityName = entityType + animal
      }

      if c < len(rows[r]) - 1 && rows[r][c + 1] != ' ' {
        track = append(track, "right")
      }

      if c > 0 && rows[r][c - 1] != ' ' {
        track = append(track, "left")
      }

      if r > 0 && rows[r - 1][c] != ' ' {
        track = append(track, "top");
      }

      if r < len(rows) - 1 && rows[r + 1][c] != ' ' {
        track = append(track, "bottom")
      }

      matrix[r / 2][c / 2] = Entity{Name: entityName, Track: strings.Join(track, " ")}
    }
  }

  return &matrix;
}
