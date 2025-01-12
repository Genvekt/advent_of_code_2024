package day_25

import (
  "fmt"

  "advent_of_code_2024/common"
)

const (
  dataFile = "input/day_25.txt"
)

type Key struct {
  heights []int
}

func (k *Key) Hash() string {
  return fmt.Sprintf("%v", k.heights)
}

type Lock struct {
  space []int
}

func (l *Lock) Hash() string {
  return fmt.Sprintf("%v", l.space)
}

func (l *Lock) DoesKeyFit(key *Key) bool {
  for spaceID := range l.space {
    if key.heights[spaceID] > l.space[spaceID] {
      return false
    }
  }
  return true
}

func readInput() (map[string]*Key, map[string]*Lock, error) {
  keys := map[string]*Key{}
  locks := map[string]*Lock{}

  err := common.ReadInput(dataFile, func(lines []string) error {
    lineId := 0

    for lineId < len(lines) {
      //Lock
      if lines[lineId] == "#####" {
        lock := &Lock{
          space: make([]int, 5),
        }
        for i := 1; i <= 5; i++ {
          cells := []rune(lines[lineId+i])
          for cellID, cell := range cells {
            if cell == '.' {
              lock.space[cellID] += 1
            }
          }
        }
        locks[lock.Hash()] = lock

        // Key
      } else {
        key := &Key{
          heights: make([]int, 5),
        }
        for i := 1; i <= 5; i++ {
          cells := []rune(lines[lineId+i])
          for cellID, cell := range cells {
            if cell == '#' {
              key.heights[cellID] += 1
            }
          }
        }
        keys[key.Hash()] = key
      }

      lineId += 8
    }

    return nil
  })

  return keys, locks, err
}

func Solve() error {
  keys, locks, err := readInput()
  if err != nil {
    return err
  }

  uniquePairs := map[string]bool{}

  for _, lock := range locks {
    for _, key := range keys {
      if lock.DoesKeyFit(key) {
        uniquePairs[key.Hash()+" "+lock.Hash()] = true
      }
    }
  }

  fmt.Println(len(uniquePairs))

  return nil
}
