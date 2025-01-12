package day_6

import (
  "fmt"

  "advent_of_code_2024/common"
)

type Direction int

type Move struct {
  Row int
  Col int
}

const (
  facingUp    Direction = 0
  facingDown  Direction = 1
  facingLeft  Direction = 2
  facingRight Direction = 3
)

type Guard struct {
  Row       int
  Col       int
  Direction Direction
}

func (g *Guard) MakeStep(m *Map) {
  switch g.Direction {
  case facingUp:
    if g.Row == 0 {
      g.Row = -1
      return
    }

    if (*m)[g.Row-1][g.Col] == '#' {
      g.Direction = facingRight
    } else {
      g.Row--
    }
  case facingRight:
    if g.Col == len((*m)[0])-1 {
      g.Col++
      return
    }

    if (*m)[g.Row][g.Col+1] == '#' {
      g.Direction = facingDown
    } else {
      g.Col++
    }
  case facingDown:
    if g.Row == len((*m))-1 {
      g.Row++
      return
    }

    if (*m)[g.Row+1][g.Col] == '#' {
      g.Direction = facingLeft
    } else {
      g.Row++
    }
  case facingLeft:
    if g.Col == 0 {
      g.Col--
      return
    }

    if (*m)[g.Row][g.Col-1] == '#' {
      g.Direction = facingUp
    } else {
      g.Col--
    }
  }
}

func (g *Guard) Hash() string {
  return fmt.Sprintf("%d %d", g.Row, g.Col)
}

func (g *Guard) HashWithDirection() string {
  return fmt.Sprintf("%d %d %d", g.Row, g.Col, g.Direction)
}

func (g *Guard) ObstacleCoords(m *Map) (int, int) {
  switch g.Direction {
  case facingUp:
    if g.Row == 0 {
      return -1, -1
    }
    return g.Row - 1, g.Col
  case facingDown:
    if g.Row == len(*m)-1 {
      return -1, -1
    }
    return g.Row + 1, g.Col
  case facingLeft:
    if g.Col == 0 {
      return -1, -1
    }
    return g.Row, g.Col - 1
  default:
    if g.Col == len((*m)[0])-1 {
      return -1, -1
    }
    return g.Row, g.Col + 1
  }
}

func (g *Guard) IsOut(m *Map) bool {
  if g.Row < 0 || g.Col < 0 || g.Row >= len(*m) || g.Col >= len((*m)[0]) {
    return true
  }
  return false
}

type Map [][]rune

func readInput() (*Guard, *Map, error) {
  guard := &Guard{Direction: facingUp}
  var m Map

  err := common.ReadInput("input/day_6.txt", func(lines []string) error {
    for rowId, line := range lines {
      row := []rune(line)
      m = append(m, row)
      for colId, cell := range row {
        if cell == '^' {
          guard.Row = rowId
          guard.Col = colId
        }
      }
    }
    return nil
  })

  if err != nil {
    return nil, nil, err
  }

  return guard, &m, nil
}

func Solve() error {
  guard, m, err := readInput()
  if err != nil {
    return err
  }

  moves := map[string]struct{}{}

  moves[guard.Hash()] = struct{}{}

  for !guard.IsOut(m) {
    guard.MakeStep(m)
    moves[guard.Hash()] = struct{}{}
  }

  fmt.Println(len(moves))
  return nil
}

func SolvePart2() error {
  guard, m, err := readInput()
  if err != nil {
    return err
  }

  initialRow, initialCol := guard.Row, guard.Col

  moves := map[string]*Move{}

  moves[guard.Hash()] = &Move{
    Row: guard.Row,
    Col: guard.Col,
  }

  for !guard.IsOut(m) {
    moves[guard.Hash()] = &Move{
      Row: guard.Row,
      Col: guard.Col,
    }
    guard.MakeStep(m)
  }

  count := 0

  for _, move := range moves {
    if move.Row == initialRow && move.Col == initialCol {
      continue
    }

    (*m)[move.Row][move.Col] = '#'
    directions := map[string]struct{}{}
    guard.Row = initialRow
    guard.Col = initialCol
    guard.Direction = facingUp

    for !guard.IsOut(m) {
      guard.MakeStep(m)
      if _, ok := directions[guard.HashWithDirection()]; ok {
        count++
        break
      }
      directions[guard.HashWithDirection()] = struct{}{}
    }
    (*m)[move.Row][move.Col] = '.'

  }

  fmt.Println(count)
  return nil

}
