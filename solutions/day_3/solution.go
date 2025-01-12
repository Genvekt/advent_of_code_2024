package day_3

import (
  "fmt"
  "regexp"
  "sort"
  "strconv"

  "advent_of_code_2024/common"
)

const (
  typeMul  = "MUL"
  typeDo   = "DO"
  typeDont = "DONT"
)

type Operation struct {
  operationType string
  position      int
  a             int64
  b             int64
}

func readInput() ([]Operation, error) {
  var inputs []Operation
  var activatedOperations []Operation
  err := common.ReadInput("input/day_3.txt", func(lines []string) error {
    rData := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
    rDo := regexp.MustCompile(`do\(\)`)
    rDont := regexp.MustCompile(`don't\(\)`)

    for _, line := range lines {
      paramIds := rData.FindAllStringSubmatchIndex(line, -1)
      params := rData.FindAllStringSubmatch(line, -1)

      doIds := rDo.FindAllStringSubmatchIndex(line, -1)
      dontIds := rDont.FindAllStringSubmatchIndex(line, -1)

      for paramId := range paramIds {
        a, _ := strconv.ParseInt(params[paramId][1], 10, 64)
        b, _ := strconv.ParseInt(params[paramId][2], 10, 64)
        inputs = append(inputs, Operation{operationType: typeMul, a: a, b: b, position: paramIds[paramId][0]})
      }

      for doId := range doIds {
        inputs = append(inputs, Operation{operationType: typeDo, position: doIds[doId][0]})
      }
      for dontId := range dontIds {
        inputs = append(inputs, Operation{operationType: typeDont, position: dontIds[dontId][0]})
      }
    }

    sort.Slice(inputs, func(i, j int) bool {
      return inputs[i].position < inputs[j].position
    })

    fmt.Println(inputs)

    isActivated := true
    for _, op := range inputs {
      switch op.operationType {
      case typeMul:
        if isActivated {
          activatedOperations = append(activatedOperations, op)
        }
      case typeDo:
        isActivated = true
      case typeDont:
        isActivated = false
      }
    }
    return nil
  })

  if err != nil {
    return nil, err
  }

  fmt.Println(activatedOperations)
  return activatedOperations, nil
}

func Solve() error {
  data, err := readInput()
  if err != nil {
    return err
  }

  var result int64
  for _, entry := range data {
    result += entry.a * entry.b
  }
  fmt.Println(result)
  return nil
}
