package day_4

import (
  "errors"
  "fmt"

  "advent_of_code_2024/common"
)

type Field [][]rune

func readInput() (Field, error) {
  var field Field
  err := common.ReadInput("input/test.txt", func(lines []string) error {
    for _, line := range lines {
      field = append(field, []rune(line))
    }
    return nil
  })
  if err != nil {
    return nil, err
  }

  return field, nil
}

func Solve(word string) error {
  field, err := readInput()
  if err != nil {
    return err
  }

  searchSequence := []rune(word)

  searchFunctions := []func(word []rune, field Field, rowId int, colId int) bool{
    searchHorizontal,
    searchHorizontalReversed,
    searchVertical,
    searchVerticalReversed,
    searchDiagonalUp,
    searchDiagonalUpReversed,
    searchDiagonalDown,
    searchDiagonalDownReversed,
  }

  var count int64

  for rowId := range field {
    for colId := range field[rowId] {
      for _, searchFunc := range searchFunctions {
        if searchFunc(searchSequence, field, rowId, colId) {
          count++
        }
      }
    }
  }

  fmt.Println(count)
  return nil
}

func searchHorizontal(word []rune, field Field, rowId int, colId int) bool {
  if colId > len(field[0])-len(word) {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId][colId+i] != word[i] {
      return false
    }
  }

  return true
}

func searchHorizontalReversed(word []rune, field Field, rowId int, colId int) bool {
  if colId < len(word)-1 {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId][colId-i] != word[i] {
      return false
    }
  }
  return true
}

func searchVertical(word []rune, field Field, rowId int, colId int) bool {
  if rowId > len(field)-len(word) {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId+i][colId] != word[i] {
      return false
    }
  }

  return true
}

func searchVerticalReversed(word []rune, field Field, rowId int, colId int) bool {
  if rowId < len(word)-1 {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId-i][colId] != word[i] {
      return false
    }
  }

  return true
}

func searchDiagonalUp(word []rune, field Field, rowId int, colId int) bool {
  if rowId < len(word)-1 || colId > len(field[0])-len(word) {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId-i][colId+i] != word[i] {
      return false
    }
  }

  return true
}

func searchDiagonalUpReversed(word []rune, field Field, rowId int, colId int) bool {
  if rowId < len(word)-1 || colId < len(word)-1 {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId-i][colId-i] != word[i] {
      return false
    }
  }

  return true
}

func searchDiagonalDown(word []rune, field Field, rowId int, colId int) bool {
  if rowId > len(field)-len(word) || colId > len(field[0])-len(word) {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId+i][colId+i] != word[i] {
      return false
    }
  }

  return true
}

func searchDiagonalDownReversed(word []rune, field Field, rowId int, colId int) bool {
  if rowId > len(field)-len(word) || colId < len(word)-1 {
    return false
  }

  for i := 0; i < len(word); i++ {
    if field[rowId+i][colId-i] != word[i] {
      return false
    }
  }

  return true
}

func SolvePart2(word string) error {
  if len(word)%2 != 1 {
    return errors.New("invalid word length: must be odd")
  }

  field, err := readInput()
  if err != nil {
    return err
  }

  var count int64

  searchSequence := []rune(word)
  sequenceHalfLen := len(word) / 2

  for rowId := range field {
    for colId := range field[rowId] {
      if rowId < sequenceHalfLen ||
        colId < sequenceHalfLen ||
        rowId > len(field)-1-sequenceHalfLen ||
        colId > len(field[0])-1-sequenceHalfLen {
        continue
      }

      if (searchDiagonalUp(searchSequence, field, rowId+sequenceHalfLen, colId-sequenceHalfLen) || searchDiagonalDownReversed(searchSequence, field, rowId-sequenceHalfLen, colId+sequenceHalfLen)) &&
        (searchDiagonalDown(searchSequence, field, rowId-sequenceHalfLen, colId-sequenceHalfLen) || searchDiagonalUpReversed(searchSequence, field, rowId+sequenceHalfLen, colId+sequenceHalfLen)) {
        count++
      }
    }
  }

  fmt.Println(count)
  return nil
}
