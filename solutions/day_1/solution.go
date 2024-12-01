package day_1

import (
  "fmt"
  "sort"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

type List []int64

func readInput() (List, List, error) {
  var list1 List
  var list2 List

  process := func(lines []string) error {
    for _, line := range lines {
      parts := strings.Split(line, "   ")
      if len(parts) != 2 {
        return fmt.Errorf("invalid line: %s", line)
      }

      num1, err := strconv.ParseInt(parts[0], 10, 64)
      if err != nil {
        return err
      }

      num2, err := strconv.ParseInt(parts[1], 10, 64)
      if err != nil {
        return err
      }

      list1 = append(list1, num1)
      list2 = append(list2, num2)
    }

    return nil
  }

  err := common.ReadInput("input/day_1.txt", process)
  //err := common.ReadInput("input/test.txt", process)
  if err != nil {
    return nil, nil, err
  }

  return list1, list2, nil
}

func Solve() error {
  list1, list2, err := readInput()
  if err != nil {
    return err
  }

  sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
  sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

  res := listSimilarity(list1, list2)
  fmt.Println(res)
  return nil
}

func listDistance(list1, list2 List) int64 {
  var distance int64
  for i := 0; i < len(list1); i++ {
    dis := list1[i] - list2[i]
    if dis < 0 {
      dis = -dis
    }
    distance += dis
  }
  return distance
}

func listSimilarity(list1, list2 List) int64 {
  counter := map[int64]int64{}
  for _, num := range list2 {
    counter[num] = counter[num] + 1
  }

  var similarity int64
  for _, num := range list1 {
    if freq, found := counter[num]; found {
      similarity += freq * num
    }
  }

  return similarity
}
