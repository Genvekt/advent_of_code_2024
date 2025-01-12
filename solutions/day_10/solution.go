package day_10

import (
  "fmt"

  "advent_of_code_2024/common"
)

type Map [][]int

func findEnds(m Map, rowId int, colId int, currentHeight int) map[string]struct{} {
  if rowId < 0 || rowId >= len(m) || colId < 0 || colId >= len(m[rowId]) {
    return map[string]struct{}{}
  }

  if m[rowId][colId] != currentHeight {
    return map[string]struct{}{}
  }

  if currentHeight == 9 {
    return map[string]struct{}{fmt.Sprintf("%d %d", rowId, colId): {}}
  }

  toTop := findEnds(m, rowId-1, colId, currentHeight+1)
  toRight := findEnds(m, rowId, colId+1, currentHeight+1)
  toBot := findEnds(m, rowId+1, colId, currentHeight+1)
  toLeft := findEnds(m, rowId, colId-1, currentHeight+1)

  for k, v := range toRight {
    toTop[k] = v
  }
  for k, v := range toBot {
    toTop[k] = v
  }
  for k, v := range toLeft {
    toTop[k] = v
  }

  return toTop
}

func findPaths(m Map, rowId int, colId int, currentHeight int) int64 {
  if rowId < 0 || rowId >= len(m) || colId < 0 || colId >= len(m[rowId]) {
    return 0
  }

  if m[rowId][colId] != currentHeight {
    return 0
  }

  if currentHeight == 9 {
    return 1
  }

  toTop := findPaths(m, rowId-1, colId, currentHeight+1)
  toRight := findPaths(m, rowId, colId+1, currentHeight+1)
  toBot := findPaths(m, rowId+1, colId, currentHeight+1)
  toLeft := findPaths(m, rowId, colId-1, currentHeight+1)

  return toTop + toRight + toBot + toLeft
}

func readInput() (Map, error) {
  var m Map
  err := common.ReadInput("input/day_10.txt", func(lines []string) error {
    for _, line := range lines {
      var row []int
      cells := []rune(line)
      for _, cell := range cells {
        row = append(row, int(cell-'0'))
      }

      m = append(m, row)
    }
    return nil
  })
  if err != nil {
    return nil, err
  }

  return m, nil
}

func Solve() error {
  m, err := readInput()
  if err != nil {
    return err
  }

  var count int64

  for rowId := range m {
    for colId := range m[rowId] {
      count += int64(len(findEnds(m, rowId, colId, 0)))
    }
  }

  fmt.Println(count)
  return nil
}

func SolvePart2() error {
  m, err := readInput()
  if err != nil {
    return err
  }

  var count int64

  for rowId := range m {
    for colId := range m[rowId] {
      count += findPaths(m, rowId, colId, 0)
    }
  }

  fmt.Println(count)
  return nil
}
