package day_7

import (
  "fmt"
  "regexp"
  "strconv"

  "advent_of_code_2024/common"
)

type Operator rune

const (
  sumOp    Operator = '+'
  mulOp    Operator = '*'
  concatOp Operator = '|'
)

type Equation struct {
  Result    int64
  Values    []int64
  Operators []Operator
}

func (e *Equation) IsCorrect() bool {
  res := e.Values[0]

  for opId := range e.Operators {
    switch e.Operators[opId] {
    case sumOp:
      res += e.Values[opId+1]
    case mulOp:
      res *= e.Values[opId+1]
    case concatOp:
      concated := fmt.Sprintf("%d%d", res, e.Values[opId+1])
      res, _ = strconv.ParseInt(concated, 10, 64)
    }
  }

  return res == e.Result
}

func readInput() ([]*Equation, error) {
  var equations []*Equation
  err := common.ReadInput("input/day_7.txt", func(lines []string) error {
    equationRegex := regexp.MustCompile(`(\d+)\D`)
    for _, line := range lines {
      data := equationRegex.FindAllStringSubmatch(line+" ", -1)
      eq := &Equation{}
      eq.Result, _ = strconv.ParseInt(data[0][1], 10, 64)
      for _, elem := range data[1:] {
        val, _ := strconv.ParseInt(elem[1], 10, 64)
        eq.Values = append(eq.Values, val)
      }
      equations = append(equations, eq)
    }
    return nil
  })

  if err != nil {
    return nil, err
  }

  return equations, nil
}

func Solve() error {
  equations, err := readInput()
  if err != nil {
    return err
  }

  possibleOperators := []Operator{sumOp, mulOp, concatOp}

  var sum int64

  for _, equation := range equations {
    if findCombo(equation, possibleOperators, 0) {
      sum += equation.Result
    }
  }

  fmt.Println(int64(sum))
  return nil
}

func findCombo(eq *Equation, operations []Operator, opId int) bool {
  if opId == len(eq.Values)-1 {
    if eq.IsCorrect() {
      return true
    }
    return false
  }

  for _, operator := range operations {
    if opId == len(eq.Operators) {
      eq.Operators = append(eq.Operators, operator)
    } else {
      eq.Operators[opId] = operator
    }

    if findCombo(eq, operations, opId+1) {
      return true
    }
  }

  return false
}
