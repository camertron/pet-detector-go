package main

import "fmt"
import "sync"
import "strconv"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "encoding/json"
import "path/filepath"
// import "github.com/davecgh/go-spew/spew"

func main() {
  var solutions map[string][]*SolutionStep
  solutionsMutex := &sync.Mutex{}
  solutionData, _ := ioutil.ReadFile("./solutions.json")
  json.Unmarshal([]byte(solutionData), &solutions)

  files, _ := filepath.Glob("/Users/cameron/Desktop/game_results/*-*")

  for _, file := range files {
    fmt.Println(file)
    processFile(file, &solutions, solutionsMutex)
  }

  marshalled, _ := json.MarshalIndent(solutions, "", " ")
  _ = ioutil.WriteFile("./solutions.json", marshalled, 0644)
}

func processFile(file string, solutions *map[string][]*SolutionStep, mutex *sync.Mutex) {
  data, _ := ioutil.ReadFile(file)
  rows := strings.Split(string(data), "\n")

  for _, row := range rows {
    gameResultFields := strings.Split(row, "\t")

    // not sure why this happens?
    if len(gameResultFields) < 10 {
      continue
    }

    go processJson(gameResultFields[9], solutions, mutex)
  }
}

func processJson(gameResultJson string, solutions *map[string][]*SolutionStep, mutex *sync.Mutex) {
  var gameResultData GameResultData
  json.Unmarshal([]byte(gameResultJson), &gameResultData)

  roundReader := csv.NewReader(strings.NewReader(strings.Replace(gameResultData.Round_csv, "/", "\n", -1)))
  rounds, _ := roundReader.ReadAll()

  // trialReader := csv.NewReader(strings.NewReader(strings.Replace(gameResultData.Trial_csv, "/", "\n", -1)))
  // trials, _ := trialReader.ReadAll()

  // firstMoves := make(map[string]*SolutionStep)

  // for i, trial := range trials {
  //   // skip header, make sure there are two numbers in the address
  //   if i == 0 || len(trial[0]) < 2 {
  //     continue
  //   }

  //   round := trial[2]

  //   if _, ok := firstMoves[round]; !ok {
  //     rowAndCol := strings.Split(trial[0], "")
  //     rRaw, cRaw := rowAndCol[0], rowAndCol[1]
  //     r, _ := strconv.Atoi(rRaw)
  //     c, _ := strconv.Atoi(cRaw)
  //     firstMoves[round] = &SolutionStep{R: r, C: c}
  //   }
  // }

  for i, round := range rounds {
    // skip header
    if i == 0 {
      continue
    }

    gas, _ := strconv.Atoi(round[2])  // 2nd column is fuel
    matrix := ParseLevel(round[10])   // 10th column is trial_csv
    signature := matrix.GetSignature()
    var solution []*SolutionStep

    mutex.Lock()
    foundSolution, foundOk := (*solutions)[signature]
    mutex.Unlock()

    if foundOk {
      solution = foundSolution
    } else {
      solution = Solve(matrix, gas)

      mutex.Lock()
      (*solutions)[signature] = solution
      mutex.Unlock()
    }
  }
}
