package day_2

import (
  "fmt"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

type Report []int64

func (r Report) IsSafe(idToIgnore int) bool {
  if len(r) < 3 {
    return true
  }

  isIncreasing := r[0] < r[1]
  if idToIgnore == 0 {
    isIncreasing = r[1] < r[2]
  } else if idToIgnore == 1 {
    isIncreasing = r[0] < r[2]
  }

  for i := 0; i < len(r)-1; i++ {
    left := i
    right := i + 1
    if left == idToIgnore {
      if idToIgnore == 0 {
        continue
      }
      left--
    }

    if right == idToIgnore {
      if right == len(r)-1 {
        continue
      }
      right++
    }

    if r[left] < r[right] != isIncreasing {
      return false
    }

    diff := r[right] - r[left]

    if isIncreasing && (diff < 1 || diff > 3) {
      return false
    } else if !isIncreasing && (diff < -3 || diff > -1) {
      return false
    }
  }

  return true
}

func readReports() ([]Report, error) {
  var reports []Report

  err := common.ReadInput("input/day_2.txt", func(lines []string) error {
    for _, line := range lines {
      var report Report
      parts := strings.Fields(line)
      for _, part := range parts {
        num, err := strconv.ParseInt(part, 10, 64)
        if err != nil {
          return err
        }
        report = append(report, num)
      }

      reports = append(reports, report)
    }
    return nil
  })
  if err != nil {
    return nil, err
  }

  return reports, nil
}

func Solve() error {
  reports, err := readReports()
  if err != nil {
    return err
  }

  counter := 0
  for _, report := range reports {

    if report.IsSafe(-1) {
      counter++
    } else {
      for i := 0; i < len(report); i++ {
        if report.IsSafe(i) {
          counter++
          break
        }
      }
    }
  }

  fmt.Println(counter)
  return nil
}
