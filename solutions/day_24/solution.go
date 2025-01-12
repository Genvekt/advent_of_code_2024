package day_24

import (
  "fmt"
  "regexp"
  "sort"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

const dataFile = "input/day_24_solved.txt"

type Register struct {
  name  string
  value bool
}

type Expression struct {
  kind   string
  input1 string
  input2 string
  result string
}

func (e *Expression) String() string {
  switch e.kind {
  case "AND":
    return fmt.Sprintf("assign %s = %s & %s; ", e.result, e.input1, e.input2)
  case "OR":
    return fmt.Sprintf("assign %s = %s | %s; ", e.result, e.input1, e.input2)
  case "XOR":
    return fmt.Sprintf("assign %s = %s ^ %s; ", e.result, e.input1, e.input2)
  }
  return ""
}

func (e *Expression) evaluate(value1 bool, value2 bool) bool {
  switch e.kind {
  case "AND":
    return value1 && value2
  case "OR":
    return value1 || value2
  case "XOR":
    return value1 != value2
  }

  return false
}

func readInput() (map[string]bool, []*Expression, error) {
  valies := map[string]bool{}
  var expressions []*Expression

  err := common.ReadInput(dataFile, func(lines []string) error {
    initialValueRegex := regexp.MustCompile(`^(\w+): (0|1)$`)
    expressionRegex := regexp.MustCompile(`^(\w+) (XOR|OR|AND) (\w+) -> (\w+)$`)
    lastLineId := 0
    for lineId, line := range lines {
      if line == "" {
        lastLineId = lineId + 1
        break
      }

      parts := initialValueRegex.FindAllStringSubmatch(line, -1)
      valies[parts[0][1]] = parts[0][2] == "1"
    }

    for _, line := range lines[lastLineId:] {
      parts := expressionRegex.FindAllStringSubmatch(line, -1)
      expression := &Expression{
        kind:   parts[0][2],
        input1: parts[0][1],
        input2: parts[0][3],
        result: parts[0][4],
      }
      expressions = append(expressions, expression)
    }

    return nil
  })

  return valies, expressions, err
}

func Solve() error {
  values, expressions, err := readInput()
  if err != nil {
    return err
  }

  findAdditionFlowMistake(expressions)
  return nil

  inputs := []string{}
  outputs := []string{}
  wires := []string{}

  for inputName := range values {
    inputs = append(inputs, inputName)
  }

  for _, expression := range expressions {
    if strings.HasPrefix(expression.result, "z") {
      outputs = append(outputs, expression.result)
    } else {
      wires = append(wires, expression.result)
    }
    fmt.Println(expression.String())
  }

  input1 := registersToNum(values, "x")
  input2 := registersToNum(values, "y")

  targetResult := input1 + input2
  targetStr := strconv.FormatInt(targetResult, 2)

  valuesCopy := map[string]bool{}
  for k, v := range values {
    valuesCopy[k] = v
  }

  gotResult := solveSystem(valuesCopy, expressions)

  fmt.Println(targetStr)
  fmt.Println(gotResult)

  return nil
}

func solveSystem(values map[string]bool, expressions []*Expression) int64 {
  solved := false
  atLeastOneResolved := true
  for !solved && atLeastOneResolved {
    solved = true
    atLeastOneResolved = false
    for _, expression := range expressions {
      if _, calculated := values[expression.result]; calculated {
        continue
      }

      value1, calculatedValue1 := values[expression.input1]
      value2, calculatedValue2 := values[expression.input2]

      if !calculatedValue1 || !calculatedValue2 {
        solved = false
        continue
      }

      resultVal := expression.evaluate(value1, value2)
      values[expression.result] = resultVal
      atLeastOneResolved = true
    }
  }
  if !atLeastOneResolved {
    fmt.Println("No solution found for system")
    return 0
  }
  return registersToNum(values, "z")
}

func registersToNum(values map[string]bool, prefix string) int64 {
  var outRegisters []*Register
  for registerName, registerValue := range values {
    if strings.HasPrefix(registerName, prefix) {
      register := &Register{
        name:  registerName,
        value: registerValue,
      }
      outRegisters = append(outRegisters, register)
    }
  }

  sort.Slice(outRegisters, func(i, j int) bool {
    return outRegisters[i].name < outRegisters[j].name
  })

  resNum := ""
  for _, register := range outRegisters {
    if register.value {
      resNum = "1" + resNum
    } else {
      resNum = "0" + resNum
    }
  }

  res, _ := strconv.ParseInt(resNum, 2, 64)
  return res
}

func findAdditionFlowMistake(wires []*Expression) {
  idx := 0
  spareOneRegister := ""

  for idx < 45 {
    idxStr := fmt.Sprintf("%d", idx)
    if idx < 10 {
      idxStr = "0" + idxStr
    }
    xRegister := fmt.Sprintf("x%s", idxStr)
    yRegister := fmt.Sprintf("y%s", idxStr)

    oneSaved := ""

    for _, w := range wires {
      if w.kind == "AND" && ((w.input1 == xRegister && w.input2 == yRegister) || (w.input2 == xRegister && w.input1 == yRegister)) {
        oneSaved = w.result
        break
      }
    }

    if oneSaved == "" {
      fmt.Printf("cannot find %s AND %s \n", xRegister, yRegister)
      return
    }

    additionResWithoutSpareOne := ""
    for _, w := range wires {
      if w.kind == "XOR" && ((w.input1 == xRegister && w.input2 == yRegister) || (w.input2 == xRegister && w.input1 == yRegister)) {
        additionResWithoutSpareOne = w.result
        break
      }
    }

    if additionResWithoutSpareOne == "" {
      fmt.Printf("cannot find %s XOR %s \n", xRegister, yRegister)
      return
    }

    if spareOneRegister != "" {
      res := ""
      for _, w := range wires {
        if w.kind == "XOR" && ((w.input1 == spareOneRegister && w.input2 == additionResWithoutSpareOne) || (w.input2 == spareOneRegister && w.input1 == additionResWithoutSpareOne)) {
          res = w.result
          break
        }
      }

      if res == "" {
        fmt.Printf("cannot find %s XOR %s (for position %s %s) \n", spareOneRegister, additionResWithoutSpareOne, xRegister, yRegister)
        return
      }

      if !strings.HasPrefix(res, "z") {
        fmt.Printf("finded %s XOR %s (for position %s %s), but res not z, it is %s \n", spareOneRegister, additionResWithoutSpareOne, xRegister, yRegister, res)
        return
      }

      additionalSpareOne := ""
      for _, w := range wires {
        if w.kind == "AND" && ((w.input1 == spareOneRegister && w.input2 == additionResWithoutSpareOne) || (w.input2 == spareOneRegister && w.input1 == additionResWithoutSpareOne)) {
          additionalSpareOne = w.result
          break
        }
      }

      if additionalSpareOne == "" {
        fmt.Printf("cannot find %s AND %s (for position %s %s) \n", spareOneRegister, additionResWithoutSpareOne, xRegister, yRegister)
        return
      }

      newSpareOneRegister := ""
      for _, w := range wires {
        if w.kind == "OR" && ((w.input1 == oneSaved && w.input2 == additionalSpareOne) || (w.input2 == oneSaved && w.input1 == additionalSpareOne)) {
          newSpareOneRegister = w.result
          break
        }
      }

      if newSpareOneRegister == "" {
        fmt.Printf("cannot find %s OR %s (for position %s %s) \n", oneSaved, additionalSpareOne, xRegister, yRegister)
        return
      }

      spareOneRegister = newSpareOneRegister

    } else {
      spareOneRegister = oneSaved
    }

    idx++
  }
}
