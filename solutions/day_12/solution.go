package day_12

import (
  "fmt"
  "sort"

  "advent_of_code_2024/common"
)

type Field [][]rune
type Fence struct {
  Row int
  Col int
}

func Hash(rowId int, colId int) string {
  return fmt.Sprintf("%d %d", rowId, colId)
}

func readInput() (Field, error) {
  var field Field
  err := common.ReadInput("input/day_12.txt", func(lines []string) error {
    for _, line := range lines {
      cells := []rune(line)
      field = append(field, cells)
    }
    return nil
  })
  if err != nil {
    return nil, err
  }

  return field, nil
}

func (f Field) FindRegion(
  rowId int,
  colId int,
  regionType rune,
  regionCells map[string]bool,
  regionFences map[string]map[string]*Fence,
) (map[string]bool, map[string]map[string]*Fence, bool) {
  if rowId < 0 || rowId >= len(f) || colId < 0 || colId >= len(f[rowId]) {
    return regionCells, regionFences, true
  }

  if f[rowId][colId] != regionType {
    return regionCells, regionFences, true
  }

  cellHash := Hash(rowId, colId)
  if _, included := regionCells[cellHash]; included {
    return regionCells, regionFences, false
  }

  regionCells[cellHash] = true

  moves := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
  for _, move := range moves {
    cells, fences, fenceTriggered := f.FindRegion(rowId+move[0], colId+move[1], regionType, regionCells, regionFences)
    if fenceTriggered {
      // we found one fence
      if move[0] == -1 {
        regionFences["up"][cellHash] = &Fence{Row: rowId, Col: colId}
      } else if move[0] == 1 {
        regionFences["down"][cellHash] = &Fence{Row: rowId, Col: colId}
      } else if move[1] == -1 {
        regionFences["left"][cellHash] = &Fence{Row: rowId, Col: colId}
      } else {
        regionFences["right"][cellHash] = &Fence{Row: rowId, Col: colId}
      }
    }
    for k, v := range cells {
      regionCells[k] = v
    }
    for ft, fs := range fences {
      for fh, fe := range fs {
        regionFences[ft][fh] = fe
      }
    }
  }

  return regionCells, regionFences, false
}

func Solve() error {
  field, err := readInput()
  if err != nil {
    return err
  }

  visitedCells := map[string]bool{}

  var cost int64

  for rowId := range field {
    for colId := range field[rowId] {
      if _, visited := visitedCells[Hash(rowId, colId)]; visited {
        // cell already included in some region
        continue
      }

      regionCells, fences, _ := field.FindRegion(rowId, colId, field[rowId][colId], map[string]bool{}, map[string]map[string]*Fence{"up": {}, "down": {}, "left": {}, "right": {}})
      //fmt.Println(string([]rune{field[rowId][colId]}))
      //for t, fs := range fences {
      //  fmt.Printf("%s: ", t)
      //  for _, f := range fs {
      //    fmt.Printf("%v ", f)
      //  }
      //  fmt.Println()
      //}
      lines := 0
      for fenceType, fenceParts := range fences {
        switch fenceType {
        case "up", "down":
          lines += CountHorizontalLines(fenceParts)
        default:
          lines += CountVerticalLines(fenceParts)
        }
      }

      //fmt.Println("Lines: ", lines)

      cost += int64(lines * len(regionCells))

      for key := range regionCells {
        visitedCells[key] = true
      }
    }
  }

  fmt.Println(cost)
  return nil
}

func CountHorizontalLines(fenceParts map[string]*Fence) int {
  var parts []*Fence
  for _, fe := range fenceParts {
    parts = append(parts, fe)
  }

  sort.Slice(parts, func(i, j int) bool {
    if parts[i].Row == parts[j].Row {
      return parts[i].Col < parts[j].Col
    }
    return parts[i].Row < parts[j].Row
  })

  lines := 0

  for i := 1; i < len(parts); i++ {
    if parts[i-1].Row != parts[i].Row {
      lines++
      continue
    }

    if parts[i].Col-parts[i-1].Col != 1 {
      lines++
      continue
    }
  }

  return lines + 1
}

func CountVerticalLines(fenceParts map[string]*Fence) int {
  var parts []*Fence
  for _, fe := range fenceParts {
    parts = append(parts, fe)
  }

  sort.Slice(parts, func(i, j int) bool {
    if parts[i].Col == parts[j].Col {
      return parts[i].Row < parts[j].Row
    }
    return parts[i].Col < parts[j].Col
  })

  lines := 0

  for i := 1; i < len(parts); i++ {
    if parts[i-1].Col != parts[i].Col {
      lines++
      continue
    }

    if parts[i].Row-parts[i-1].Row != 1 {
      lines++
      continue
    }
  }

  return lines + 1
}
