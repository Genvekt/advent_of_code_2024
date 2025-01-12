package day_19

import (
  "fmt"
  "regexp"
  "strings"

  "advent_of_code_2024/common"
)

const (
  dataFile = "input/day_19.txt"
)

func readInput() ([]string, []string, error) {
  var patterns []string
  var desiredStrings []string

  err := common.ReadInput(dataFile, func(lines []string) error {
    lastLineId := 0
    for lineID, line := range lines {
      if line == "" {
        lastLineId = lineID + 1
        break
      }

      parts := strings.Split(line, ", ")
      patterns = append(patterns, parts...)
    }

    for _, line := range lines[lastLineId:] {
      desiredStrings = append(desiredStrings, line)
    }

    return nil
  })

  return patterns, desiredStrings, err
}

var cache = map[string]int{}

func findArragement(patterns []string, regexPattern *regexp.Regexp, testString string) int {
  if len(testString) == 0 {
    return 1
  }

  if solution, solved := cache[testString]; solved {
    return solution
  }

  matches := regexPattern.FindStringSubmatch(testString)
  if len(matches) == 0 {
    return 0
  }

  var count int
  for _, pattern := range patterns {
    if strings.HasPrefix(testString, pattern) {
      newTestString := testString[len(pattern):]
      count += findArragement(patterns, regexPattern, newTestString)
    }
  }

  cache[testString] = count
  return count
}

func Solve() error {
  patterns, desiredStrings, err := readInput()
  if err != nil {
    return err
  }

  var count int

  regexPattern := regexp.MustCompile("^(" + strings.Join(patterns, "|") + ")+$")
  for testID, testString := range desiredStrings {
    fmt.Println(testID, "/", len(desiredStrings))
    count += findArragement(patterns, regexPattern, testString)
    fmt.Println(count)
  }

  fmt.Println(count)

  return nil
}
