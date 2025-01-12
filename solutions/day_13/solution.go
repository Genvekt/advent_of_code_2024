package day_13

import (
  "fmt"
  "regexp"
  "strconv"

  "advent_of_code_2024/common"
)

type Problem struct {
  Xa int64
  Ya int64
  Xb int64
  Yb int64
  Xp int64
  Yp int64
}

func (p *Problem) Solve() (int64, int64, bool) {
  b := ((p.Yp * p.Xa) - (p.Ya * p.Xp)) / ((p.Xa * p.Yb) - (p.Xb * p.Ya))
  a := (p.Xp - p.Xb*b) / p.Xa

  if p.Xa*a+p.Xb*b == p.Xp && p.Ya*a+p.Yb*b == p.Yp {
    return a, b, true
  }

  return 0, 0, false
}

func readInput() ([]*Problem, error) {
  var problems []*Problem
  err := common.ReadInput("input/day_13.txt", func(lines []string) error {

    aButtonRegex := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
    bButtonRegex := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
    priceRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
    lineId := 0
    for lineId < len(lines)-2 {
      partsA := aButtonRegex.FindAllStringSubmatch(lines[lineId], -1)
      partsB := bButtonRegex.FindAllStringSubmatch(lines[lineId+1], -1)
      partsPrice := priceRegex.FindAllStringSubmatch(lines[lineId+2], -1)

      xa, _ := strconv.ParseInt(partsA[0][1], 10, 64)
      ya, _ := strconv.ParseInt(partsA[0][2], 10, 64)
      xb, _ := strconv.ParseInt(partsB[0][1], 10, 64)
      yb, _ := strconv.ParseInt(partsB[0][2], 10, 64)
      xp, _ := strconv.ParseInt(partsPrice[0][1], 10, 64)
      yp, _ := strconv.ParseInt(partsPrice[0][2], 10, 64)

      problems = append(problems, &Problem{Xa: xa, Ya: ya, Xb: xb, Yb: yb, Xp: xp, Yp: yp})

      lineId += 4
    }

    return nil
  })

  return problems, err
}

func Solve() error {
  problems, err := readInput()
  if err != nil {
    return err
  }

  var cost int64

  for _, problem := range problems {
    problem.Xp += 10000000000000
    problem.Yp += 10000000000000
    a, b, solvable := problem.Solve()
    if solvable {
      cost += a*3 + b
    }
  }

  fmt.Println(cost)
  return nil
}
