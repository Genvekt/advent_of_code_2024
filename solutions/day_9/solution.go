package day_9

import (
  "fmt"
  "strconv"

  "advent_of_code_2024/common"
)

type MemoryCell struct {
  FileID int
}

func (m *MemoryCell) IsFileCell() bool {
  return m.FileID >= 0
}

func readInput() ([]*MemoryCell, error) {
  var memory []*MemoryCell
  err := common.ReadInput("input/day_9.txt", func(lines []string) error {
    fileId := 0
    for _, line := range lines {
      runes := []rune(line)
      for cellId, cell := range runes {
        if cellId%2 == 0 {
          fileLength, _ := strconv.Atoi(string(cell))
          for i := 0; i < fileLength; i++ {
            memory = append(memory, &MemoryCell{FileID: fileId})
          }
          fileId++
        } else {
          freeSpaceLen, _ := strconv.Atoi(string(cell))
          for i := 0; i < freeSpaceLen; i++ {
            memory = append(memory, &MemoryCell{FileID: -1})
          }
        }
      }
    }

    return nil
  })

  if err != nil {
    return nil, err
  }

  return memory, err
}

func Solve() error {
  memory, err := readInput()
  if err != nil {
    return err
  }

  // find last file cell
  fileCellId := len(memory) - 1
  for fileCellId >= 0 && !memory[fileCellId].IsFileCell() {
    fileCellId--
  }

  // find first free cel
  freeCellId := 0
  for freeCellId < len(memory) && memory[freeCellId].IsFileCell() {
    freeCellId++
  }

  // move files at the beginning
  for freeCellId < fileCellId {
    memory[freeCellId], memory[fileCellId] = memory[fileCellId], memory[freeCellId]
    freeCellId++
    fileCellId--
    for fileCellId >= 0 && !memory[fileCellId].IsFileCell() {
      fileCellId--
    }
    for freeCellId < len(memory) && memory[freeCellId].IsFileCell() {
      freeCellId++
    }
  }

  var checksum int64
  for cellId, cell := range memory {
    if !cell.IsFileCell() {
      break
    }

    checksum += int64(cellId) * int64(cell.FileID)
  }

  fmt.Printf("Checksum: %d\n", checksum)
  return nil
}

func SolvePart2() error {
  memory, err := readInput()
  if err != nil {
    return err
  }

  fileStart, _, fileLen, fileFound := findFileBoundaries(memory, len(memory)-1)

  for fileFound {
    spaceStart, spaceEnd, spaceLen, spaceFound := findFreeSpace(memory, 0)
    for spaceFound && spaceEnd < fileStart {
      if spaceLen >= fileLen {
        for i := 0; i < fileLen; i++ {
          memory[spaceStart+i], memory[fileStart+i] = memory[fileStart+i], memory[spaceStart+i]
        }
        break
      } else {
        spaceStart, spaceEnd, spaceLen, spaceFound = findFreeSpace(memory, spaceEnd+1)
      }
    }
    fileStart, _, fileLen, fileFound = findFileBoundaries(memory, fileStart-1)
  }

  var checksum int64
  for cellId, cell := range memory {
    if !cell.IsFileCell() {
      continue
    }

    checksum += int64(cellId) * int64(cell.FileID)
  }

  //for cellId, cell := range memory {
  //  fmt.Println(cellId, *cell)
  //}

  fmt.Printf("Checksum: %d\n", checksum)
  return nil
}

func findFileBoundaries(memory []*MemoryCell, maxCellId int) (int, int, int, bool) {
  fileEnd := maxCellId
  for fileEnd >= 0 && !memory[fileEnd].IsFileCell() {
    fileEnd--
  }

  if fileEnd == -1 {
    return 0, 0, 0, false
  }

  fileStart := fileEnd
  fileId := memory[fileStart].FileID
  for fileStart > 0 && memory[fileStart-1].FileID == fileId {
    fileStart--
  }

  fileLen := fileEnd - fileStart + 1

  return fileStart, fileEnd, fileLen, true
}

func findFreeSpace(memory []*MemoryCell, minCellId int) (int, int, int, bool) {
  spaceStart := minCellId
  for spaceStart < len(memory) && memory[spaceStart].IsFileCell() {
    spaceStart++
  }

  if spaceStart == len(memory) {
    return 0, 0, 0, false
  }

  spaceEnd := spaceStart
  for spaceEnd < len(memory)-1 && !memory[spaceEnd+1].IsFileCell() {
    spaceEnd++
  }

  spaceLen := spaceEnd - spaceStart + 1

  return spaceStart, spaceEnd, spaceLen, true
}
