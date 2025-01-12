package day_15

import (
  "fmt"
  "os"
  "os/exec"

  "advent_of_code_2024/common"
)

const (
  dataFile     = "input/test.txt"
  robotSign    = '🤖'
  boxSign      = '🎁'
  leftBoxSign  = '👉'
  rightBoxSign = '👈'
  wallSign     = '🎄'
  spaceSign    = '🟦'
)

type Move rune

type Map [][]rune

func (m Map) Print() {
  for _, row := range m {
    fmt.Println(string(row))
  }
}

func (m Map) CountGPS() int64 {
  var count int64
  for rowID, row := range m {
    for colID, col := range row {
      if col == boxSign {
        count += int64(rowID*100 + colID)
      }
    }
  }

  return count
}

func (m Map) CountGPSPart2() int64 {
  var count int64
  for rowID, row := range m {
    for colID, col := range row {
      if col == leftBoxSign {
        count += int64(rowID*100 + colID)
      }
    }
  }

  return count
}

type Robot struct {
  row int
  col int
}

func moveRobot(m Map, robot *Robot, move Move) {
  switch move {
  case '^':
    freeSpaceCol := robot.col
    freeSpaceRow := robot.row - 1
    for m[freeSpaceRow][freeSpaceCol] != spaceSign {
      if m[freeSpaceRow][freeSpaceCol] == wallSign {
        // no move possible
        return
      }
      freeSpaceRow--
    }
    for freeSpaceRow < robot.row-1 {
      m[freeSpaceRow][freeSpaceCol] = boxSign
      freeSpaceRow++
    }
    m[freeSpaceRow][freeSpaceCol] = robotSign
    m[freeSpaceRow+1][freeSpaceCol] = spaceSign
    robot.row--
  case 'v':
    freeSpaceCol := robot.col
    freeSpaceRow := robot.row + 1
    for m[freeSpaceRow][freeSpaceCol] != spaceSign {
      if m[freeSpaceRow][freeSpaceCol] == wallSign {
        // no move possible
        return
      }
      freeSpaceRow++
    }
    for freeSpaceRow > robot.row+1 {
      m[freeSpaceRow][freeSpaceCol] = boxSign
      freeSpaceRow--
    }
    m[freeSpaceRow][freeSpaceCol] = robotSign
    m[freeSpaceRow-1][freeSpaceCol] = spaceSign
    robot.row++
  case '>':
    freeSpaceCol := robot.col + 1
    freeSpaceRow := robot.row
    for m[freeSpaceRow][freeSpaceCol] != spaceSign {
      if m[freeSpaceRow][freeSpaceCol] == wallSign {
        // no move possible
        return
      }
      freeSpaceCol++
    }
    for freeSpaceCol > robot.col+1 {
      m[freeSpaceRow][freeSpaceCol] = boxSign
      freeSpaceCol--
    }
    m[freeSpaceRow][freeSpaceCol] = robotSign
    m[freeSpaceRow][freeSpaceCol-1] = spaceSign
    robot.col++
  case '<':
    freeSpaceCol := robot.col - 1
    freeSpaceRow := robot.row
    for m[freeSpaceRow][freeSpaceCol] != spaceSign {
      if m[freeSpaceRow][freeSpaceCol] == wallSign {
        // no move possible
        return
      }
      freeSpaceCol--
    }
    for freeSpaceCol < robot.col-1 {
      m[freeSpaceRow][freeSpaceCol] = boxSign
      freeSpaceCol++
    }
    m[freeSpaceRow][freeSpaceCol] = robotSign
    m[freeSpaceRow][freeSpaceCol+1] = spaceSign
    robot.col--
  }
}

func readInput() (Map, *Robot, []Move, error) {
  var m Map
  var moves []Move
  robot := &Robot{}
  err := common.ReadInput(dataFile, func(lines []string) error {
    lastRow := 0
    for rowID, line := range lines {
      if line == "" {
        lastRow = rowID + 1
        break
      }
      var cells []rune
      for cellID, cell := range []rune(line) {
        switch cell {
        case '@':
          robot.row = rowID
          robot.col = cellID
          cells = append(cells, robotSign)
        case '#':
          cells = append(cells, wallSign)
        case '.':
          cells = append(cells, spaceSign)
        case 'O':
          cells = append(cells, boxSign)
        }
      }
      m = append(m, cells)
    }

    for _, line := range lines[lastRow:] {
      movesPart := []Move(line)
      moves = append(moves, movesPart...)
    }

    return nil
  })

  if err != nil {
    return nil, nil, nil, err
  }

  return m, robot, moves, nil
}

func Solve() error {
  m, robot, moves, err := readInput()
  if err != nil {
    return err
  }

  for _, move := range moves {
    moveRobot(m, robot, move)
    //cmd := exec.Command("clear") //Linux example, its tested
    //cmd.Stdout = os.Stdout
    //cmd.Run()
    //fmt.Println("Move: ", string(move))
    //m.Print()
    //fmt.Println()
  }

  fmt.Println(m.CountGPS())
  return nil
}

func readInputPart2() (Map, *Robot, []Move, error) {
  var m Map
  var moves []Move
  robot := &Robot{}
  err := common.ReadInput(dataFile, func(lines []string) error {
    lastRow := 0
    for rowID, line := range lines {
      if line == "" {
        lastRow = rowID + 1
        break
      }
      var cells []rune
      for cellID, cell := range []rune(line) {
        switch cell {
        case '@':
          robot.row = rowID
          robot.col = cellID * 2
          cells = append(cells, robotSign)
          cells = append(cells, spaceSign)
        case '#':
          cells = append(cells, wallSign)
          cells = append(cells, wallSign)
        case '.':
          cells = append(cells, spaceSign)
          cells = append(cells, spaceSign)
        case 'O':
          cells = append(cells, leftBoxSign)
          cells = append(cells, rightBoxSign)
        }
      }
      m = append(m, cells)
    }

    for _, line := range lines[lastRow:] {
      movesPart := []Move(line)
      moves = append(moves, movesPart...)
    }

    return nil
  })

  if err != nil {
    return nil, nil, nil, err
  }

  return m, robot, moves, nil
}

func SolvePart2() error {
  m, robot, moves, err := readInputPart2()
  if err != nil {
    return err
  }

  cmd := exec.Command("clear") //Linux example, its tested
  cmd.Stdout = os.Stdout
  cmd.Run()
  fmt.Println("Move: ")
  m.Print()
  fmt.Println()

  for _, move := range moves {
    moveRobotPart2(m, robot, move)
    cmd := exec.Command("clear") //Linux example, its tested
    cmd.Stdout = os.Stdout
    cmd.Run()
    fmt.Println("Move: ", string(move))
    m.Print()
    fmt.Println()
  }

  fmt.Println(m.CountGPSPart2())
  return nil
}

func moveRobotPart2(m Map, robot *Robot, move Move) {
  switch move {
  case '^':
    moveRobotUp(m, robot)
  case 'v':
    moveRobotDown(m, robot)
  case '<':
    moveRobotLeft(m, robot)
  case '>':
    moveRobotRight(m, robot)
  }
}

func moveRobotUp(m Map, robot *Robot) {
  switch m[robot.row-1][robot.col] {
  case wallSign:
    return
  case leftBoxSign:
    if !canMoveBoxUp(m, robot.col, robot.col+1, robot.row-1) {
      return
    }
    moveBoxUp(m, robot.col, robot.col+1, robot.row-1)
  case rightBoxSign:
    if !canMoveBoxUp(m, robot.col-1, robot.col, robot.row-1) {
      return
    }
    moveBoxUp(m, robot.col-1, robot.col, robot.row-1)
  }

  m[robot.row-1][robot.col] = robotSign
  m[robot.row][robot.col] = spaceSign

  robot.row--
}

func moveRobotDown(m Map, robot *Robot) {
  switch m[robot.row+1][robot.col] {
  case wallSign:
    return
  case leftBoxSign:
    if !canMoveBoxDown(m, robot.col, robot.col+1, robot.row+1) {
      return
    }
    moveBoxDown(m, robot.col, robot.col+1, robot.row+1)
  case rightBoxSign:
    if !canMoveBoxDown(m, robot.col-1, robot.col, robot.row+1) {
      return
    }
    moveBoxDown(m, robot.col-1, robot.col, robot.row+1)
  }

  m[robot.row+1][robot.col] = robotSign
  m[robot.row][robot.col] = spaceSign

  robot.row++
}

func moveRobotRight(m Map, robot *Robot) {
  switch m[robot.row][robot.col+1] {
  case wallSign:
    return
  case leftBoxSign:
    if !canMoveBoxRight(m, robot.col+1, robot.col+2, robot.row) {
      return
    }
    moveBoxRight(m, robot.col+1, robot.col+2, robot.row)
  }

  m[robot.row][robot.col+1] = robotSign
  m[robot.row][robot.col] = spaceSign

  robot.col++
}

func moveRobotLeft(m Map, robot *Robot) {
  switch m[robot.row][robot.col-1] {
  case wallSign:
    return
  case rightBoxSign:
    if !canMoveBoxLeft(m, robot.col-2, robot.col-1, robot.row) {
      return
    }
    moveBoxLeft(m, robot.col-2, robot.col-1, robot.row)
  }

  m[robot.row][robot.col-1] = robotSign
  m[robot.row][robot.col] = spaceSign

  robot.col--
}

func moveBox(m Map, boxLeftCol int, boxRightCol int, boxRow int, move Move) {
  switch move {
  case '^':
    moveBoxUp(m, boxLeftCol, boxRightCol, boxRow)
  case '>':
    moveBoxRight(m, boxLeftCol, boxRightCol, boxRow)
  case 'v':
    moveBoxDown(m, boxLeftCol, boxRightCol, boxRow)
  case '<':
    moveBoxLeft(m, boxLeftCol, boxRightCol, boxRow)
  }
}

func moveBoxUp(m Map, boxLeftCol int, boxRightCol int, boxRow int) {
  switch m[boxRow-1][boxLeftCol] {
  case wallSign:
    return
  case rightBoxSign:
    moveBoxUp(m, boxLeftCol-1, boxRightCol-1, boxRow-1)
  case leftBoxSign:
    moveBoxUp(m, boxLeftCol, boxRightCol, boxRow-1)
  }

  switch m[boxRow-1][boxRightCol] {
  case wallSign:

    return
  case leftBoxSign:
    moveBoxUp(m, boxLeftCol+1, boxRightCol+1, boxRow-1)
  }

  m[boxRow-1][boxLeftCol] = leftBoxSign
  m[boxRow-1][boxRightCol] = rightBoxSign

  m[boxRow][boxLeftCol] = spaceSign
  m[boxRow][boxRightCol] = spaceSign
}

func moveBoxDown(m Map, boxLeftCol int, boxRightCol int, boxRow int) {
  switch m[boxRow+1][boxLeftCol] {
  case wallSign:
    return
  case rightBoxSign:
    moveBoxDown(m, boxLeftCol-1, boxRightCol-1, boxRow+1)
  case leftBoxSign:
    moveBoxDown(m, boxLeftCol, boxRightCol, boxRow+1)
  }

  switch m[boxRow+1][boxRightCol] {
  case wallSign:
    return
  case leftBoxSign:
    moveBoxDown(m, boxLeftCol+1, boxRightCol+1, boxRow+1)
  }

  m[boxRow+1][boxLeftCol] = leftBoxSign
  m[boxRow+1][boxRightCol] = rightBoxSign

  m[boxRow][boxLeftCol] = spaceSign
  m[boxRow][boxRightCol] = spaceSign
}

func moveBoxLeft(m Map, boxLeftCol int, boxRightCol int, boxRow int) {
  switch m[boxRow][boxLeftCol-1] {
  case wallSign:
    return
  case rightBoxSign:
    moveBoxLeft(m, boxLeftCol-2, boxRightCol-2, boxRow)
  }

  m[boxRow][boxLeftCol-1] = leftBoxSign
  m[boxRow][boxRightCol-1] = rightBoxSign

  m[boxRow][boxRightCol] = spaceSign
}

func moveBoxRight(m Map, boxLeftCol int, boxRightCol int, boxRow int) {
  switch m[boxRow][boxRightCol+1] {
  case wallSign:
    return
  case leftBoxSign:
    moveBoxRight(m, boxLeftCol+2, boxRightCol+2, boxRow)
  }

  m[boxRow][boxLeftCol+1] = leftBoxSign
  m[boxRow][boxRightCol+1] = rightBoxSign

  m[boxRow][boxLeftCol] = spaceSign
}

func canMoveBox(m Map, boxLeftCol int, boxRightCol int, boxRow int, move Move) bool {
  switch move {
  case '^':
    return canMoveBoxUp(m, boxLeftCol, boxRightCol, boxRow)
  case '>':
    return canMoveBoxRight(m, boxLeftCol, boxRightCol, boxRow)
  case 'v':
    return canMoveBoxDown(m, boxLeftCol, boxRightCol, boxRow)
  case '<':
    return canMoveBoxLeft(m, boxLeftCol, boxRightCol, boxRow)
  }

  return false
}

func canMoveBoxUp(m Map, boxLeftCol int, boxRightCol int, boxRow int) bool {
  canMove := true

  switch m[boxRow-1][boxLeftCol] {
  case wallSign:
    canMove = false
  case rightBoxSign:
    canMove = canMoveBoxUp(m, boxLeftCol-1, boxRightCol-1, boxRow-1)
  case leftBoxSign:
    canMove = canMoveBoxUp(m, boxLeftCol, boxRightCol, boxRow-1)
  default:
    canMove = true
  }

  if !canMove {
    return false
  }

  switch m[boxRow-1][boxRightCol] {
  case wallSign:
    canMove = false
  case leftBoxSign:
    canMove = canMoveBoxUp(m, boxLeftCol+1, boxRightCol+1, boxRow-1)
  default:
    canMove = true
  }

  return canMove
}

func canMoveBoxDown(m Map, boxLeftCol int, boxRightCol int, boxRow int) bool {
  canMove := true

  switch m[boxRow+1][boxLeftCol] {
  case wallSign:
    canMove = false
  case rightBoxSign:
    canMove = canMoveBoxDown(m, boxLeftCol-1, boxRightCol-1, boxRow+1)
  case leftBoxSign:
    canMove = canMoveBoxDown(m, boxLeftCol, boxRightCol, boxRow+1)
  default:
    canMove = true
  }

  if !canMove {
    return false
  }

  switch m[boxRow+1][boxRightCol] {
  case wallSign:
    canMove = false
  case leftBoxSign:
    canMove = canMoveBoxDown(m, boxLeftCol+1, boxRightCol+1, boxRow+1)
  default:
    canMove = true
  }

  return canMove
}

func canMoveBoxLeft(m Map, boxLeftCol int, boxRightCol int, boxRow int) bool {
  canMove := true

  switch m[boxRow][boxLeftCol-1] {
  case wallSign:
    canMove = false
  case rightBoxSign:
    canMove = canMoveBoxLeft(m, boxLeftCol-2, boxRightCol-2, boxRow)
  default:
    canMove = true
  }
  return canMove
}

func canMoveBoxRight(m Map, boxLeftCol int, boxRightCol int, boxRow int) bool {
  canMove := true

  switch m[boxRow][boxRightCol+1] {
  case wallSign:
    canMove = false
  case leftBoxSign:
    canMove = canMoveBoxRight(m, boxLeftCol+2, boxRightCol+2, boxRow)
  default:
    canMove = true
  }
  return canMove
}
