package day_20

import (
  "fmt"

  "advent_of_code_2024/common"
)

const (
  dataFile    = "input/day_20.txt"
  cellWall    = '#'
  cellFree    = '.'
  cellStart   = 'S'
  cellEnd     = 'E'
  maxCheatLen = 20
)

type Distance struct {
  pos      *Cell
  distance int
}

type State struct {
  pos *Cell

  visitedCells map[string]struct{}
}

func (s *State) pathLen() int {
  return len(s.visitedCells)
}

func (s *State) Copy() *State {
  newState := &State{
    pos:          &Cell{row: s.pos.row, col: s.pos.col},
    visitedCells: map[string]struct{}{},
  }

  for k, v := range s.visitedCells {
    newState.visitedCells[k] = v
  }

  return newState
}

type Cheat struct {
  node1 *Cell
  node2 *Cell
  cost  int
}

func (c *Cheat) Hash() string {
  var first *Cell
  var second *Cell

  if c.node1.row == c.node2.row {
    if c.node1.col <= c.node2.col {
      first = c.node1
      second = c.node2
    } else {
      first = c.node2
      second = c.node1
    }
  } else if c.node1.row < c.node2.row {
    first = c.node1
    second = c.node2

  } else {
    first = c.node2
    second = c.node1
  }

  return fmt.Sprintf("%s %s", first.Hash(), second.Hash())

}

type Map [][]rune

func (m Map) Print() {
  for rowID, row := range m {
    fmt.Print(rowID, ":")
    if rowID < 10 {
      fmt.Print(" ")
    }
    fmt.Println(string(row))
  }
}

func (m Map) CalculateMinDists(start *Cell) map[string]int {
  minDists := map[string]int{}

  dists := map[string]*Distance{start.Hash(): &Distance{start, 0}}

  for len(dists) > 0 {
    var bestDist *Distance
    for _, dist := range dists {
      if bestDist == nil || dist.distance < bestDist.distance {
        bestDist = dist
      }
    }

    if bestDist == nil {
      continue
    }

    delete(dists, bestDist.pos.Hash())
    minDists[bestDist.pos.Hash()] = bestDist.distance

    for _, move := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
      nextDist := &Distance{pos: &Cell{row: bestDist.pos.row, col: bestDist.pos.col}, distance: bestDist.distance + 1}
      nextDist.pos.row += move[0]
      nextDist.pos.col += move[1]

      if m[nextDist.pos.row][nextDist.pos.col] == cellWall {
        continue
      }

      if _, registered := minDists[nextDist.pos.Hash()]; registered {
        continue
      }

      if existingDist, ok := dists[nextDist.pos.Hash()]; !ok || nextDist.distance < existingDist.distance {
        dists[nextDist.pos.Hash()] = nextDist
      }
    }
  }

  return minDists
}

func (m Map) FindCheats(start *Cell) map[string]*Cheat {
  cheats := map[string]*Cheat{}

  if m[start.row][start.col] == cellWall {
    return cheats
  }

  for rowStep := 0; rowStep <= maxCheatLen; rowStep++ {
    for colStep := 0; colStep <= maxCheatLen-rowStep; colStep++ {

      for _, direction := range [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}} {
        cell := &Cell{row: start.row + (rowStep * direction[0]), col: start.col + (colStep * direction[1])}
        if cell.row < 0 || cell.col < 0 || cell.row > len(m)-1 || cell.col > len(m)-1 {
          continue
        }

        if m[cell.row][cell.col] == cellWall {
          continue
        }

        cheat := &Cheat{node1: start, node2: cell, cost: rowStep + colStep}
        cheats[cheat.Hash()] = cheat
      }

    }
  }
  return cheats
}

type Cell struct {
  row int
  col int
}

func (c *Cell) Hash() string {
  return fmt.Sprintf("%d %d", c.row, c.col)
}

func readInput() (Map, *Cell, *Cell, error) {
  var m Map
  start := &Cell{}
  end := &Cell{}

  err := common.ReadInput(dataFile, func(lines []string) error {
    for rowID, line := range lines {
      row := []rune(line)
      for cellID, cell := range row {
        if cell == cellStart {
          start.row = rowID
          start.col = cellID
        } else if cell == cellEnd {
          end.row = rowID
          end.col = cellID
        }
      }

      m = append(m, row)
    }

    return nil
  })

  return m, start, end, err
}

func Solve() error {
  m, start, end, err := readInput()
  if err != nil {
    return err
  }

  distancesFromStart := m.CalculateMinDists(start)
  distancesFromEnd := m.CalculateMinDists(end)

  originalMinPath := distancesFromEnd[start.Hash()]

  allCheats := map[string]*Cheat{}

  for rowID := 1; rowID < len(m)-1; rowID++ {
    for colID := 1; colID < len(m[rowID])-1; colID++ {
      cheatStart := &Cell{row: rowID, col: colID}
      cheats := m.FindCheats(cheatStart)
      for _, cheat := range cheats {
        allCheats[cheat.Hash()] = cheat
      }
    }
  }

  count := map[int]int{}

  for _, cheat := range allCheats {
    path1 := distancesFromStart[cheat.node1.Hash()] + cheat.cost + distancesFromEnd[cheat.node2.Hash()]
    path2 := distancesFromStart[cheat.node2.Hash()] + cheat.cost + distancesFromEnd[cheat.node1.Hash()]

    bestPath := path1
    if path1 > path2 {
      bestPath = path2
    }

    if bestPath >= originalMinPath {
      continue
    }

    count[originalMinPath-bestPath] += 1
  }

  sum := 0
  for saving, freq := range count {
    if saving >= 100 {
      sum += freq
    }
  }

  fmt.Println(sum)
  return nil
}
