package day_11

import (
  "fmt"
  "math"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

func ChangeNum(num int64) []int64 {

  if num == 0 {
    return []int64{1}
  }

  strVal := fmt.Sprintf("%d", num)
  if len(strVal)%2 == 0 {
    half := int64(math.Pow(10, float64(len(strVal))/2))

    firstVal := num / half
    secondVal := num % half

    return []int64{firstVal, secondVal}
  }

  return []int64{num * 2024}
}

type Line map[int64]int64

func (l Line) Lenght() int64 {
  var res int64
  for _, v := range l {
    res += v
  }

  return res
}

func (l Line) FullChange() Line {
  newLine := make(Line, len(l))

  for num, count := range l {
    newNums := ChangeNum(num)
    for _, newNum := range newNums {
      newLine[newNum] = newLine[newNum] + count
    }
  }

  return newLine
}

func readInput() (Line, error) {
  line := Line{}
  err := common.ReadInput("input/day_11.txt", func(lines []string) error {
    for _, row := range lines {
      parts := strings.Split(row, " ")
      for _, part := range parts {
        val, _ := strconv.ParseInt(part, 10, 64)
        line[val] = line[val] + 1
      }
      return nil
    }

    return nil
  })
  if err != nil {
    return nil, err
  }

  return line, nil
}

func Solve() error {
  line, err := readInput()
  if err != nil {
    return err
  }

  for i := 0; i < 75; i++ {
    line = line.FullChange()
  }

  fmt.Println(line.Lenght())
  return nil
}
