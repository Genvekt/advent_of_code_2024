package day_18

import (
  "fmt"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

const (
  dataFile       = "input/day_18.txt"
  sequenceLength = 1024

  fieldDimention = 71

  cellFree = '.'
  cellWall = '#'
)

type Coord struct {
  row, col int
}

func (c Coord) Hash() string {
  return fmt.Sprintf("%d %d", c.row, c.col)
}

type State struct {
  position     *Coord
  visitedCells map[string]struct{}
}

func (s *State) Copy() *State {
  newState := &State{}
  newState.visitedCells = map[string]struct{}{}
  for k, v := range s.visitedCells {
    newState.visitedCells[k] = v
  }

  pos := &Coord{row: s.position.row, col: s.position.col}
  newState.position = pos

  return newState
}

type Field [][]rune

func (f Field) Print() {
  for _, row := range f {
    line := string(row)
    fmt.Println(line)
  }
}

func (f Field) BFS(startRow int, startCol int) map[string]struct{} {
  startPos := &Coord{row: startRow, col: startCol}
  startState := &State{position: startPos, visitedCells: map[string]struct{}{startPos.Hash(): {}}}
  states := map[string]*State{startState.position.Hash(): startState}

  processedStates := map[string]*State{}

  for len(states) > 0 {
    var shortestState *State
    for _, state := range states {
      if shortestState == nil || len(state.visitedCells) < len(shortestState.visitedCells) {
        shortestState = state
      }
    }

    delete(states, shortestState.position.Hash())
    processedStates[shortestState.position.Hash()] = shortestState

    if shortestState.position.row == fieldDimention-1 && shortestState.position.col == fieldDimention-1 {
      return shortestState.visitedCells
    }

    // go top
    if shortestState.position.row > 0 {
      topState := shortestState.Copy()
      topState.position.row--
      if f[topState.position.row][topState.position.col] != cellWall {
        if _, visited := topState.visitedCells[topState.position.Hash()]; !visited {
          if _, processed := processedStates[topState.position.Hash()]; !processed {
            topState.visitedCells[topState.position.Hash()] = struct{}{}
            states[topState.position.Hash()] = topState
          }
        }
      }
    }

    // go down
    if shortestState.position.row < fieldDimention-1 {
      topState := shortestState.Copy()
      topState.position.row++
      if f[topState.position.row][topState.position.col] != cellWall {
        if _, visited := topState.visitedCells[topState.position.Hash()]; !visited {
          if _, processed := processedStates[topState.position.Hash()]; !processed {
            topState.visitedCells[topState.position.Hash()] = struct{}{}
            states[topState.position.Hash()] = topState
          }
        }
      }
    }

    // go right
    if shortestState.position.col < fieldDimention-1 {
      topState := shortestState.Copy()
      topState.position.col++
      if f[topState.position.row][topState.position.col] != cellWall {
        if _, visited := topState.visitedCells[topState.position.Hash()]; !visited {
          if _, processed := processedStates[topState.position.Hash()]; !processed {
            topState.visitedCells[topState.position.Hash()] = struct{}{}
            states[topState.position.Hash()] = topState
          }
        }
      }
    }

    // go left
    if shortestState.position.col > 0 {
      topState := shortestState.Copy()
      topState.position.col--
      if f[topState.position.row][topState.position.col] != cellWall {
        if _, visited := topState.visitedCells[topState.position.Hash()]; !visited {
          if _, processed := processedStates[topState.position.Hash()]; !processed {
            topState.visitedCells[topState.position.Hash()] = struct{}{}
            states[topState.position.Hash()] = topState
          }
        }
      }
    }
  }

  return nil
}

func readInput() ([]*Coord, Field, error) {
  var field Field
  var coords []*Coord

  err := common.ReadInput(dataFile, func(lines []string) error {
    for _, line := range lines {
      parts := strings.Split(line, ",")
      if len(parts) != 2 {
        return fmt.Errorf("invalid line: %s", line)
      }

      col, _ := strconv.Atoi(parts[0])
      row, _ := strconv.Atoi(parts[1])

      coords = append(coords, &Coord{row: row, col: col})
    }

    return nil
  })

  for rowId := 0; rowId < fieldDimention; rowId++ {
    var row []rune
    for colId := 0; colId < fieldDimention; colId++ {
      row = append(row, cellFree)
    }
    field = append(field, row)
  }

  return coords, field, err
}

func Solve() error {
  coords, _, err := readInput()
  if err != nil {
    return err
  }

  for i := 3000; i < len(coords); i++ {
    icoords, ifield, err := readInput()
    if err != nil {
      return err
    }

    for _, coord := range icoords[:i] {
      ifield[coord.row][coord.col] = cellWall
    }

    pathLen := len(ifield.BFS(0, 0)) - 1
    if pathLen <= 0 {
      fmt.Println("No path! ", i)
      return nil
    } else {
      fmt.Printf("Path for %d: %d\n", i, pathLen)
    }

  }

  return nil
}
