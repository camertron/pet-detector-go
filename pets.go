package main

import "fmt"
import "strconv"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "encoding/json"
// import "github.com/davecgh/go-spew/spew"

func main() {
  data, _ := ioutil.ReadFile("/Users/cameron/Desktop/game_result.json")
  var gameResult GameResult
  json.Unmarshal([]byte(data), &gameResult)
  gameResultData := gameResult.Game_result.Game_result_data

  var solutions map[string][]*SolutionStep
  solutionData, _ := ioutil.ReadFile("./solutions.json")
  json.Unmarshal([]byte(solutionData), &solutions)

  roundReader := csv.NewReader(strings.NewReader(strings.Replace(gameResultData.Round_csv, "/", "\n", -1)))
  rounds, _ := roundReader.ReadAll()

  trialReader := csv.NewReader(strings.NewReader(strings.Replace(gameResultData.Trial_csv, "/", "\n", -1)))
  trials, _ := trialReader.ReadAll()

  firstMoves := make(map[string]*SolutionStep)

  for i, trial := range trials {
    // skip header, make sure there are two numbers in the address
    if i == 0 || len(trial[0]) < 2 {
      continue
    }

    round := trial[2]

    if _, ok := firstMoves[round]; !ok {
      rowAndCol := strings.Split(trial[0], "")
      fmt.Println(rowAndCol)
      rRaw, cRaw := rowAndCol[0], rowAndCol[1]
      r, _ := strconv.Atoi(rRaw)
      c, _ := strconv.Atoi(cRaw)
      firstMoves[round] = &SolutionStep{R: r, C: c}
    }
  }

  fmt.Println("")

  for i, round := range rounds {
    // skip header
    if i == 0 {
      continue
    }

    gas, _ := strconv.Atoi(round[2])  // 2nd column is fuel
    matrix := ParseLevel(round[10])   // 10th column is trial_csv
    signature := matrix.GetSignature()
    var solution []*SolutionStep

    if foundSolution, ok := solutions[signature]; ok {
      solution = foundSolution
    } else {
      solution = Solve(matrix, gas)
      solutions[signature] = solution
    }

    for _, entity := range solution {
      fmt.Println("(" + strconv.Itoa(entity.C) + ", " + strconv.Itoa(entity.R) + ")")
    }

    fmt.Println("")
  }

  marshalled, _ := json.MarshalIndent(solutions, "", " ")
  _ = ioutil.WriteFile("./solutions.json", marshalled, 0644)
}
