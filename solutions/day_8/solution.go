package day_8

import (
  "fmt"

  "advent_of_code_2024/common"
)

type Antena struct {
  kind rune
  row  int
  col  int
}

func readInput() ([]*Antena, int, int, error) {
  var antennas []*Antena
  var height, width int

  err := common.ReadInput("input/day_8.txt", func(lines []string) error {
    height = len(lines)
    width = len(lines[0])

    for rowId, line := range lines {
      cells := []rune(line)
      for colId, cell := range cells {
        if cell != '.' {
          antennas = append(antennas, &Antena{
            kind: cell,
            row:  rowId,
            col:  colId,
          })
        }
      }
    }

    return nil
  })
  if err != nil {
    return nil, 0, 0, err
  }

  return antennas, height, width, nil
}

func Solve() error {
  antennas, height, width, err := readInput()
  if err != nil {
    return err
  }

  m := map[rune][]*Antena{}
  antennaCoords := map[string]struct{}{}
  for _, a := range antennas {
    m[a.kind] = append(m[a.kind], a)
    antennaCoords[fmt.Sprintf("%d %d", a.row, a.col)] = struct{}{}
  }

  for _, oneTypeAntennas := range m {
    if len(oneTypeAntennas) == 1 {
      continue
    }

    for i := 0; i < len(oneTypeAntennas)-1; i++ {
      for j := i + 1; j < len(oneTypeAntennas); j++ {
        vectorRows := oneTypeAntennas[i].row - oneTypeAntennas[j].row
        vectorCol := oneTypeAntennas[i].col - oneTypeAntennas[j].col

        firstAntiRow := oneTypeAntennas[i].row + vectorRows
        firstAntiCol := oneTypeAntennas[i].col + vectorCol

        for firstAntiRow >= 0 && firstAntiRow < height && firstAntiCol >= 0 && firstAntiCol < width {
          firstHash := fmt.Sprintf("%d %d", firstAntiRow, firstAntiCol)
          antennaCoords[firstHash] = struct{}{}
          firstAntiRow = firstAntiRow + vectorRows
          firstAntiCol = firstAntiCol + vectorCol
        }

        secondAntiRow := oneTypeAntennas[j].row - vectorRows
        secondAntiCol := oneTypeAntennas[j].col - vectorCol

        for secondAntiRow >= 0 && secondAntiRow < height && secondAntiCol >= 0 && secondAntiCol < width {
          secondHash := fmt.Sprintf("%d %d", secondAntiRow, secondAntiCol)
          antennaCoords[secondHash] = struct{}{}
          secondAntiRow = secondAntiRow - vectorRows
          secondAntiCol = secondAntiCol - vectorCol
        }
      }
    }
  }

  fmt.Println(len(antennaCoords))
  return nil

}
