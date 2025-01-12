package day_14

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strconv"

  "advent_of_code_2024/common"
)

const (
  //dataFile = "input/test.txt"
  //mapH     = 7
  //mapW     = 11

  dataFile = "input/day_14.txt"
  mapH     = 103
  mapW     = 101
)

type Velocity struct {
  rowsPerSecond int64
  colsPerSecond int64
}

type Robot struct {
  row      int64
  col      int64
  velocity *Velocity
}

func (r *Robot) Step(seconds int64, mapHeight int64, mapWidth int64) {
  r.row = ((r.row+r.velocity.rowsPerSecond*seconds)%mapHeight + mapHeight) % mapHeight
  r.col = ((r.col+r.velocity.colsPerSecond*seconds)%mapWidth + mapWidth) % mapWidth
}

func readInput() ([]*Robot, error) {
  var robots []*Robot

  err := common.ReadInput(dataFile, func(lines []string) error {
    robotRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
    for _, line := range lines {
      parts := robotRegex.FindAllStringSubmatch(line, -1)

      row, _ := strconv.ParseInt(parts[0][2], 10, 64)
      col, _ := strconv.ParseInt(parts[0][1], 10, 64)

      vRow, _ := strconv.ParseInt(parts[0][4], 10, 64)
      vCol, _ := strconv.ParseInt(parts[0][3], 10, 64)

      robot := &Robot{
        row: row,
        col: col,
        velocity: &Velocity{
          rowsPerSecond: vRow,
          colsPerSecond: vCol,
        },
      }
      robots = append(robots, robot)
    }

    return nil
  })

  return robots, err
}

func Solve() error {
  robots, err := readInput()
  if err != nil {
    return err
  }

  for _, robot := range robots {
    robot.Step(100, mapH, mapW)
  }

  // count in quartiles
  halfH := int64(mapH / 2)
  halfW := int64(mapW / 2)

  var q1, q2, q3, q4 int64
  for _, robot := range robots {
    if robot.row < halfH {
      if robot.col < halfW {
        q1 += 1
      } else if robot.col > halfW {
        q2 += 1
      }
    } else if robot.row > halfH {
      if robot.col < halfW {
        q3 += 1
      } else if robot.col > halfW {
        q4 += 1
      }
    }
  }

  fmt.Println(q1, q2, q3, q4, q1*q2*q3*q4)

  return nil
}

func PrintMap(robots []*Robot, mapHeight int64, mapWidth int64) {
  m := make([][]int, mapHeight)
  for i := range m {
    m[i] = make([]int, mapWidth)
  }

  for _, robot := range robots {
    m[robot.row][robot.col] += 1
  }

  for _, row := range m {
    for _, val := range row {
      if val == 0 {
        fmt.Print(".")
      } else {
        fmt.Printf("%d", val)
      }
    }
    fmt.Print("\n")
  }
}

func PrintToFile(robots []*Robot, mapHeight int64, mapWidth int64, filename string) {
  f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    fmt.Println(err)
  }
  defer f.Close()

  m := make([][]int, mapHeight)
  for i := range m {
    m[i] = make([]int, mapWidth)
  }

  for _, robot := range robots {
    m[robot.row][robot.col] += 1
  }

  for _, row := range m {
    var rowStr []rune
    for _, val := range row {
      if val == 0 {
        rowStr = append(rowStr, '.')
      } else {
        rowStr = append(rowStr, rune('0'+val))
      }
    }
    _, err := f.WriteString(string(rowStr) + "\n")
    if err != nil {
      fmt.Println(err)
    }
  }
}

func SolvePart2() error {
  robots, err := readInput()
  if err != nil {
    return err
  }

  secondsPassed := 0
  for {
    for _, robot := range robots {
      robot.Step(1, mapH, mapW)
    }
    secondsPassed++

    if IsPossibleTreeTest(robots) {
      //PrintToFile(robots, mapH, mapW, "robotMap.txt")
      PrintMap(robots, mapH, mapW)
      fmt.Println(secondsPassed)
      fmt.Println()
      cmd := exec.Command("clear") //Linux example, its tested
      cmd.Stdout = os.Stdout
      cmd.Run()
    }
  }

  return nil
}

func IsPossibleTreeTest(robots []*Robot) bool {
  mByRow := map[int64]int64{}

  for _, robot := range robots {
    mByRow[robot.row] += 1
    if mByRow[robot.row] >= 25 {
      return true
    }
  }

  return false
}
