package day_5

import (
  "fmt"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

type Ordering map[int64]map[int64]struct{}
type Sequence []int64

func readInput() (Ordering, []Sequence, error) {
  ordering := make(Ordering)
  var sequences []Sequence

  err := common.ReadInput("input/day_5.txt", func(lines []string) error {
    i := 0
    for i < len(lines) {
      if lines[i] == "" {
        i++
        break
      }

      parts := strings.Split(lines[i], "|")
      first, _ := strconv.ParseInt(parts[0], 10, 64)
      last, _ := strconv.ParseInt(parts[1], 10, 64)

      if _, ok := ordering[first]; !ok {
        ordering[first] = make(map[int64]struct{})
      }

      ordering[first][last] = struct{}{}
      i++
    }

    for i < len(lines) {
      var seq Sequence
      parts := strings.Split(lines[i], ",")
      for _, part := range parts {
        num, _ := strconv.ParseInt(part, 10, 64)
        seq = append(seq, num)
      }
      sequences = append(sequences, seq)
      i++
    }

    return nil
  })

  if err != nil {
    return nil, nil, err
  }

  return ordering, sequences, nil
}

func isOrderCorrect(sequence Sequence, order Ordering) bool {
  met := map[int64]struct{}{}
  for _, num := range sequence {
    if numbersAfter, ok := order[num]; ok {
      for numberAfter := range numbersAfter {
        if _, alreadyMet := met[numberAfter]; alreadyMet {
          return false
        }
      }
    }

    met[num] = struct{}{}
  }

  return true
}

func fixOrder(sequence Sequence, order Ordering) {
  i := 0
  met := map[int64]int{}

outer:
  for i < len(sequence) {
    if numbersAfter, ok := order[sequence[i]]; ok {
      for numberAfter := range numbersAfter {
        if lastIdx, alreadyMet := met[numberAfter]; alreadyMet && lastIdx < i {
          // exchange to fix ordering
          sequence[lastIdx], sequence[i] = sequence[i], sequence[lastIdx]
          met[numberAfter] = i
          i = lastIdx
          continue outer
        }
      }
    }

    met[sequence[i]] = i
    i++
  }
}

func Solve() error {
  ordering, sequences, err := readInput()
  if err != nil {
    return err
  }

  var correctSequences []Sequence

  for _, seq := range sequences {
    if isOrderCorrect(seq, ordering) {
      correctSequences = append(correctSequences, seq)
    }
  }

  var sum int64
  for _, sequence := range correctSequences {
    mid := len(sequence) / 2
    sum += sequence[mid]
  }

  fmt.Println(sum)
  return nil

}

func SolvePart2() error {
  ordering, sequences, err := readInput()
  if err != nil {
    return err
  }

  var correctedSequences []Sequence

  for _, seq := range sequences {
    if !isOrderCorrect(seq, ordering) {
      fixOrder(seq, ordering)
      correctedSequences = append(correctedSequences, seq)
    }
  }

  var sum int64
  for _, sequence := range correctedSequences {
    mid := len(sequence) / 2
    sum += sequence[mid]
  }

  fmt.Println(sum)
  return nil
}
