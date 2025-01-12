package day_16

import (
  "fmt"

  "advent_of_code_2024/common"
)

type Direction int

const (
  UpDirection Direction = iota
  RightDirection
  DownDirection
  LeftDirection
)

const (
  dataFile = "input/day_16.txt"
)

type Map [][]rune

type Cell struct {
  row int
  col int
}

func (cell Cell) Hash() string {
  return fmt.Sprintf("%d %d", cell.row, cell.col)
}

type Deer struct {
  position     Cell
  direction    Direction
  score        int64
  visitedCells map[string]struct{}
}

func (d *Deer) Hash() string {
  return fmt.Sprintf("%d %d %v", d.position.row, d.position.col, d.direction)
}

func (d *Deer) RotateClockwise() *Deer {
  rotatedDeer := *d

  switch rotatedDeer.direction {
  case UpDirection:
    rotatedDeer.direction = RightDirection
  case RightDirection:
    rotatedDeer.direction = DownDirection
  case DownDirection:
    rotatedDeer.direction = LeftDirection
  case LeftDirection:
    rotatedDeer.direction = UpDirection
  }

  rotatedDeer.score += 1000

  return &rotatedDeer
}

func (d *Deer) RotateCounterClockwise() *Deer {
  rotatedDeer := *d

  switch rotatedDeer.direction {
  case UpDirection:
    rotatedDeer.direction = LeftDirection
  case LeftDirection:
    rotatedDeer.direction = DownDirection
  case DownDirection:
    rotatedDeer.direction = RightDirection
  case RightDirection:
    rotatedDeer.direction = UpDirection
  }

  rotatedDeer.score += 1000

  return &rotatedDeer
}

func (d *Deer) CanGoForward(m Map) bool {
  switch d.direction {
  case UpDirection:
    if m[d.position.row-1][d.position.col] == '#' {
      return false
    }
  case LeftDirection:
    if m[d.position.row][d.position.col-1] == '#' {
      return false
    }
  case DownDirection:
    if m[d.position.row+1][d.position.col] == '#' {
      return false
    }
  case RightDirection:
    if m[d.position.row][d.position.col+1] == '#' {
      return false
    }
  }
  return true
}

func (d *Deer) GoForward() *Deer {
  movedDeer := *d
  movedDeer.visitedCells = map[string]struct{}{}
  for k, v := range d.visitedCells {
    movedDeer.visitedCells[k] = v
  }

  switch d.direction {
  case UpDirection:
    movedDeer.position.row--
  case LeftDirection:
    movedDeer.position.col--
  case DownDirection:
    movedDeer.position.row++
  case RightDirection:
    movedDeer.position.col++
  }

  movedDeer.score += 1
  movedDeer.visitedCells[movedDeer.position.Hash()] = struct{}{}

  return &movedDeer
}

func (d *Deer) UniteVisitedCellsWith(anotherDeer *Deer) {
  for k, v := range anotherDeer.visitedCells {
    d.visitedCells[k] = v
  }
}

func readInput() (Map, *Deer, *Cell, error) {
  var m Map
  deer := &Deer{direction: RightDirection, visitedCells: map[string]struct{}{}}
  endCell := &Cell{}

  err := common.ReadInput(dataFile, func(lines []string) error {
    for rowID, line := range lines {

      cells := []rune(line)

      for colID, cell := range cells {
        switch cell {
        case 'S':
          deer.position = Cell{row: rowID, col: colID}
          deer.visitedCells[deer.position.Hash()] = struct{}{}
        case 'E':
          endCell.row = rowID
          endCell.col = colID
        }
      }

      m = append(m, cells)
    }

    return nil
  })

  if err != nil {
    return nil, nil, nil, err
  }

  return m, deer, endCell, err
}

func Solve() error {
  m, deer, endCell, err := readInput()
  if err != nil {
    return err
  }

  deers := map[string]*Deer{}
  deers[deer.Hash()] = deer

  var deerTheWinner *Deer
  processedDeers := map[string]struct{}{}

  for {

    // find the cheapestDeer
    var cheapestDeer *Deer
    for _, d := range deers {
      if cheapestDeer == nil || cheapestDeer.score > d.score {
        cheapestDeer = d
      }
    }

    delete(deers, cheapestDeer.Hash())

    if cheapestDeer.position.row == endCell.row && cheapestDeer.position.col == endCell.col {
      deerTheWinner = cheapestDeer
      break
    }

    deer1 := cheapestDeer.RotateClockwise()
    if _, notProcessed := processedDeers[deer1.Hash()]; !notProcessed {
      if deer1Old, exists := deers[deer1.Hash()]; !exists || deer1Old.score >= deer1.score {
        if deer1Old != nil && deer1Old.score == deer1.score {
          deer1.UniteVisitedCellsWith(deer1Old)
        }
        deers[deer1.Hash()] = deer1
      }
    }

    deer2 := cheapestDeer.RotateCounterClockwise()
    if _, notProcessed := processedDeers[deer2.Hash()]; !notProcessed {
      if deer2Old, exists := deers[deer2.Hash()]; !exists || deer2Old.score >= deer2.score {
        if deer2Old != nil && deer2Old.score == deer2.score {
          deer2.UniteVisitedCellsWith(deer2Old)
        }
        deers[deer2.Hash()] = deer2
      }
    }

    if cheapestDeer.CanGoForward(m) {
      deer3 := cheapestDeer.GoForward()
      if _, notProcessed := processedDeers[deer3.Hash()]; !notProcessed {
        if deer3Old, exists := deers[deer3.Hash()]; !exists || deer3Old.score >= deer3.score {
          if deer3Old != nil && deer3Old.score == deer3.score {
            deer3.UniteVisitedCellsWith(deer3Old)
          }
          deers[deer3.Hash()] = deer3
        }
      }
    }

    processedDeers[cheapestDeer.Hash()] = struct{}{}
  }

  fmt.Println(len(deerTheWinner.visitedCells))
  fmt.Println(deerTheWinner.score)

  return nil
}
