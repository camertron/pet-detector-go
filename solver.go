package main

const MAX_CAR_CAPACITY = 4

type SolutionStep struct {
  R int
  C int
}

func Solve(matrixRef *Matrix, gas int) []*SolutionStep {
  matrix := *matrixRef;
  graph := matrix.ToGraph();
  distances := graph.GetDistanceMap();

  housesAndPets := make([]*Entity, 0, matrix.GetWidth() * matrix.GetHeight());

  for r := range matrix {
    for c := range matrix[r] {
      if matrix[r][c].IsHouse() || matrix[r][c].IsPet() {
        housesAndPets = append(housesAndPets, &matrix[r][c]);
      }
    }
  }

  housePetMap := make(map[string]*Entity);

  for _, entity := range housesAndPets {
    if !entity.IsHouse() {
      continue;
    }

    for _, potentialPet := range housesAndPets {
      if potentialPet.IsPet() && entity.GetBaseName() == potentialPet.GetBaseName() {
        housePetMap[entity.Name] = potentialPet;
      }
    }
  }

  var car *Entity;

  for r := range matrix {
    for c := range matrix[r] {
      if matrix[r][c].IsCar() {
        car = &matrix[r][c];
        break;
      }
    }
  }

  solutionEntities := solveHelper(
    car, make([]*Entity, 0), housesAndPets, gas, make([]*Entity, 0), housePetMap, &distances,
  )

  solutionSteps := make([]*SolutionStep, 0, len(solutionEntities))
  for _, entity := range solutionEntities {
    solutionSteps = append(solutionSteps, &SolutionStep{R: entity.R, C: entity.C})
  }

  return solutionSteps
}

func solveHelper(
  lastEntity *Entity,
  carPets []*Entity,
  remainingEntities []*Entity,
  currentGas int,
  currentPath []*Entity,
  housePetMap map[string]*Entity,
  distances *DistanceMap) []*Entity {

  if len(carPets) == 0 && len(remainingEntities) == 0 {
    return currentPath;
  }

  candidates := getCandidates(carPets, remainingEntities, housePetMap);

  for _, candidate := range candidates {
    gas := currentGas - (*distances)[lastEntity.Name][candidate.Name];

    if gas < 0 {
      return nil;
    }

    var foundPath []*Entity;

    if candidate.IsHouse() {
      // drop off a pet
      correspondingPet := housePetMap[candidate.Name];

      for _, carPet := range carPets {
        if carPet == correspondingPet {
          newCarPets := make([]*Entity, 0, len(carPets) - 1);

          for _, newCarPet := range carPets {
            if newCarPet != correspondingPet {
              newCarPets = append(newCarPets, newCarPet);
            }
          }

          newRemaining := make([]*Entity, 0, len(remainingEntities) - 1);

          for _, remaining := range remainingEntities {
            if remaining != candidate {
              newRemaining = append(newRemaining, remaining);
            }
          }

          foundPath = solveHelper(
            candidate,
            newCarPets,
            newRemaining,
            gas,
            append(currentPath, candidate),
            housePetMap,
            distances);

          break;
        }
      }
    } else if candidate.IsPet() {
      newRemaining := make([]*Entity, 0, len(remainingEntities) - 1);

      for _, remaining := range remainingEntities {
        if remaining != candidate {
          newRemaining = append(newRemaining, remaining);
        }
      }

      foundPath = solveHelper(
        candidate,
        append(carPets, candidate),
        newRemaining,
        gas,
        append(currentPath, candidate),
        housePetMap,
        distances);
    }

    if foundPath != nil {
      return foundPath;
    }
  }

  return nil;
}

func getCandidates(carPets []*Entity, remainingEntities []*Entity, housePetMap map[string]*Entity) []*Entity {
  candidates := make([]*Entity, 0);

  if len(carPets) == 0 {
    // no pets in car means we can only visit pets
    for _, entity := range remainingEntities {
      if entity.IsPet() {
        candidates = append(candidates, entity);
      }
    }
  } else if len(carPets) == MAX_CAR_CAPACITY {
    // max pets means we can only visit houses
    for _, entity := range remainingEntities {
      if entity.IsHouse() {
        candidates = append(candidates, entity);
      }
    }
  } else {
    // can visit either a house or a pet, but only houses that correspond
    // to pets we currently have in the car
    for _, entity := range remainingEntities {
      if entity.IsHouse() {
        for _, carPet := range carPets {
          if carPet == housePetMap[entity.Name] {
            candidates = append(candidates, entity);
            break;
          }
        }
      } else if entity.IsPet() {
        candidates = append(candidates, entity)
      }
    }
  }

  return candidates;
}
