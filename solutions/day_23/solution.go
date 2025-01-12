package day_23

import (
  "fmt"
  "sort"
  "strings"

  "advent_of_code_2024/common"
)

const (
  dataFile = "input/day_23.txt"
)

func readInput() (map[string]map[string]struct{}, error) {
  m := map[string]map[string]struct{}{}

  err := common.ReadInput(dataFile, func(lines []string) error {
    for _, line := range lines {
      parts := strings.Split(line, "-")
      comp1 := parts[0]
      comp2 := parts[1]

      if _, ok := m[comp1]; !ok {
        m[comp1] = map[string]struct{}{}
      }
      if _, ok := m[comp2]; !ok {
        m[comp2] = map[string]struct{}{}
      }

      m[comp1][comp2] = struct{}{}
      m[comp2][comp1] = struct{}{}
    }

    return nil
  })

  return m, err
}

func Hash(comps map[string]struct{}) string {
  names := []string{}
  for comp := range comps {
    names = append(names, comp)
  }
  sort.Strings(names)
  return strings.Join(names, ",")
}

func Solve() error {
  m, err := readInput()
  if err != nil {
    return err
  }

  parties := map[string]struct{}{}
  for comp, connections := range m {
    if !strings.HasPrefix(comp, "t") {
      continue
    }

    candidates := []string{}
    for c := range connections {
      candidates = append(candidates, c)
    }

    for i := 0; i < len(candidates)-1; i++ {
      for j := i + 1; j < len(candidates); j++ {
        iConns := m[candidates[i]]
        if _, exist := iConns[candidates[j]]; exist {
          h := Hash(map[string]struct{}{comp: {}, candidates[i]: {}, candidates[j]: {}})
          parties[h] = struct{}{}
        }
      }
    }

  }

  fmt.Println(len(parties))
  return nil
}

func SolvePart2() error {
  m, err := readInput()
  if err != nil {
    return err
  }

  parties := map[string]int{}
  for comp, connections := range m {
    candidates := []string{}
    for c := range connections {
      candidates = append(candidates, c)
    }

    compParties := []map[string]struct{}{}

    for i := 0; i < len(candidates); i++ {
    partyLoop:
      for _, partyMembers := range compParties {
        for partyComp := range partyMembers {
          if _, connected := m[partyComp][candidates[i]]; !connected {
            continue partyLoop
          }
        }

        partyCopy := map[string]struct{}{}
        for k, v := range partyMembers {
          partyCopy[k] = v
        }
        partyCopy[candidates[i]] = struct{}{}
        compParties = append(compParties, partyCopy)
      }
      compParties = append(compParties, map[string]struct{}{candidates[i]: {}, comp: {}})
    }

    for _, party := range compParties {
      parties[Hash(party)] = len(party)
    }
  }

  maxLen := 0
  pass := ""
  for partyPass, partyLen := range parties {
    if partyLen > maxLen {
      maxLen = partyLen
      pass = partyPass
    }
  }

  fmt.Println(pass)
  return nil
}
