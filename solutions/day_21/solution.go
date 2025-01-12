package day_21

import (
  "fmt"
  "regexp"
  "strconv"

  "advent_of_code_2024/common"
)

const (
  dataFile  = "input/day_21.txt"
  robotsNum = 25
  L         = '<'
  R         = '>'
  U         = '^'
  D         = 'v'
  A         = 'A'
)

var StepCache = map[string]int{}

func Hash(from rune, to rune) string {
  return string(from) + string(to)
}

func HashStep(from rune, to rune, stepsLeft int) string {
  return fmt.Sprintf("%d %d %d", from, to, stepsLeft)
}

var NumpadBestKeys = map[string][][]rune{}
var KeysBestKeys = map[string][][]rune{}

var NumpadKeys = map[rune]map[string]int{
  A: {
    "row": 3,
    "col": 2,
  },
  '0': {
    "row": 3,
    "col": 1,
  },
  '1': {
    "row": 2,
    "col": 0,
  },
  '2': {
    "row": 2,
    "col": 1,
  },
  '3': {
    "row": 2,
    "col": 2,
  },
  '4': {
    "row": 1,
    "col": 0,
  },
  '5': {
    "row": 1,
    "col": 1,
  },
  '6': {
    "row": 1,
    "col": 2,
  },
  '7': {
    "row": 0,
    "col": 0,
  },
  '8': {
    "row": 0,
    "col": 1,
  },
  '9': {
    "row": 0,
    "col": 2,
  },
}

var ArrowKeys = map[rune]map[string]int{
  'A': {
    "row": 0,
    "col": 2,
  },
  U: {
    "row": 0,
    "col": 1,
  },
  L: {
    "row": 1,
    "col": 0,
  },
  D: {
    "row": 1,
    "col": 1,
  },
  R: {
    "row": 1,
    "col": 2,
  },
}

func NumpadPaths(start rune, end rune) [][]rune {
  doVerticalStart := true
  doHorizontalStart := true
  if (start == '7' || start == '4' || start == '1') && (end == '0' || end == A) {
    doVerticalStart = false
  } else if (end == '7' || end == '4' || end == '1') && (start == '0' || start == A) {
    doHorizontalStart = false
  }

  var paths [][]rune
  buttonStart := NumpadKeys[start]
  buttonEnd := NumpadKeys[end]

  verticalLen := buttonEnd["row"] - buttonStart["row"]
  verticalButton := D
  if verticalLen < 0 {
    verticalButton = U
    verticalLen *= -1
  }
  if verticalLen == 0 {
    doVerticalStart = false
  }

  horizontalLen := buttonEnd["col"] - buttonStart["col"]
  horizontalButton := R
  if horizontalLen < 0 {
    horizontalButton = L
    horizontalLen *= -1
  }
  if horizontalLen == 0 {
    doHorizontalStart = false
  }

  // < must always go before ^ and v (by analysis)
  // we choose do > before ^ for simplicity, in that pair order do not maters
  if doHorizontalStart && horizontalButton == L {
    doVerticalStart = false
  }

  // v must always go before > (by analysis)
  if doVerticalStart && horizontalButton == R && verticalButton == D {
    doHorizontalStart = false
  }

  if doVerticalStart {
    var verticalPath []rune
    for i := 0; i < verticalLen; i++ {
      verticalPath = append(verticalPath, verticalButton)
    }
    for i := 0; i < horizontalLen; i++ {
      verticalPath = append(verticalPath, horizontalButton)
    }
    verticalPath = append(verticalPath, A)
    paths = append(paths, verticalPath)
  }

  if doHorizontalStart {
    var horizontalPath []rune
    for i := 0; i < horizontalLen; i++ {
      horizontalPath = append(horizontalPath, horizontalButton)
    }
    for i := 0; i < verticalLen; i++ {
      horizontalPath = append(horizontalPath, verticalButton)
    }
    horizontalPath = append(horizontalPath, A)
    paths = append(paths, horizontalPath)
  }

  return paths
}

func ArrowsPath(start rune, end rune) [][]rune {
  doVerticalStart := true
  doHorizontalStart := true
  if (start == U || start == A) && end == L {
    doHorizontalStart = false
  } else if (end == U || end == A) && start == L {
    doVerticalStart = false
  }

  var paths [][]rune
  buttonStart := ArrowKeys[start]
  buttonEnd := ArrowKeys[end]

  verticalLen := buttonEnd["row"] - buttonStart["row"]
  verticalButton := D
  if verticalLen < 0 {
    verticalButton = U
    verticalLen *= -1
  }

  horizontalLen := buttonEnd["col"] - buttonStart["col"]
  horizontalButton := R
  if horizontalLen < 0 {
    horizontalButton = L
    horizontalLen *= -1
  }

  // < must always go before ^ and v (by analysis)
  // we choose do > before ^ for simplicity, in that pair order do not maters
  if doHorizontalStart && horizontalButton == L {
    doVerticalStart = false
  }

  // v must always go before > (by analysis)
  if doVerticalStart && verticalButton == D && horizontalButton == R {
    doHorizontalStart = false
  }

  if doVerticalStart {
    var verticalPath []rune
    for i := 0; i < verticalLen; i++ {
      verticalPath = append(verticalPath, verticalButton)
    }
    for i := 0; i < horizontalLen; i++ {
      verticalPath = append(verticalPath, horizontalButton)
    }
    verticalPath = append(verticalPath, A)
    paths = append(paths, verticalPath)
  }

  if doHorizontalStart {
    var horizontalPath []rune
    for i := 0; i < horizontalLen; i++ {
      horizontalPath = append(horizontalPath, horizontalButton)
    }
    for i := 0; i < verticalLen; i++ {
      horizontalPath = append(horizontalPath, verticalButton)
    }
    horizontalPath = append(horizontalPath, A)
    paths = append(paths, horizontalPath)
  }

  return paths
}

func readInput() ([][]rune, error) {
  var inLines [][]rune
  err := common.ReadInput(dataFile, func(lines []string) error {
    for _, line := range lines {
      // numpad sequences start from A
      lineRunes := []rune{A}
      lineRunes = append(lineRunes, []rune(line)...)
      inLines = append(inLines, lineRunes)
    }
    return nil
  })

  return inLines, err
}

func NumpadToArrowsSeqs(numpadSeq []rune) [][]rune {
  var results [][]rune

  if len(numpadSeq) == 1 {
    return results
  }
  sequences := NumpadBestKeys[Hash(numpadSeq[0], numpadSeq[1])]
  nextSequences := NumpadToArrowsSeqs(numpadSeq[1:])

  for _, sequence := range sequences {
    if len(nextSequences) == 0 {
      results = append(results, sequence)
      continue
    }
    for _, nextSequence := range nextSequences {
      var result []rune
      result = append(result, sequence...)
      result = append(result, nextSequence...)

      results = append(results, result)
    }
  }

  return results
}

func StartWithA(seq []rune) []rune {
  newSeq := []rune{A}
  newSeq = append(newSeq, seq...)
  return newSeq
}

func Solve() error {
  numRegex := regexp.MustCompile(`(\d+)`)
  lines, err := readInput()
  if err != nil {
    return err
  }

  for _, numFrom := range []rune{A, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
    for _, numTo := range []rune{A, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
      if numFrom == numTo {
        NumpadBestKeys[Hash(numFrom, numTo)] = [][]rune{{A}}
        continue
      }

      numpadPath := NumpadPaths(numFrom, numTo)
      NumpadBestKeys[Hash(numFrom, numTo)] = numpadPath
    }
  }

  for _, keyFrom := range []rune{A, U, L, R, D} {
    for _, keyTo := range []rune{A, U, L, R, D} {
      if keyFrom == keyTo {
        KeysBestKeys[Hash(keyFrom, keyTo)] = [][]rune{{A}}
        continue
      }

      keyPath := ArrowsPath(keyFrom, keyTo)
      KeysBestKeys[Hash(keyFrom, keyTo)] = keyPath
    }
  }

  complexitySum := 0

  for _, line := range lines {
    numpadPaths := NumpadToArrowsSeqs(line)
    bestSum := -1
    for _, numpadPath := range numpadPaths {
      sum := 0
      from := A
      to := A
      for i := 0; i < len(numpadPath); i++ {
        to = numpadPath[i]
        sum += GetLengthAfterSteps(from, to, robotsNum)
        from = numpadPath[i]
      }

      if bestSum < 0 || sum < bestSum {
        bestSum = sum
      }
    }

    parts := numRegex.FindAllStringSubmatch(string(line), -1)
    numInLine, _ := strconv.Atoi(parts[0][1])

    fmt.Println(string(line), bestSum, numInLine)
    complexitySum += bestSum * numInLine
    //374687016443140
    //374687016443140
  }

  fmt.Println(complexitySum)
  return nil
}

func GetLengthAfterSteps(from rune, to rune, stepsLeft int) int {
  if stepsLeft == 0 {
    return 1
  }

  h := HashStep(from, to, stepsLeft)
  if res, calculated := StepCache[h]; calculated {
    return res
  }

  sequences := KeysBestKeys[Hash(from, to)]
  bestSum := -1
  for _, sequence := range sequences {
    nextFrom := A
    nextTo := A

    sum := 0
    for i := 0; i < len(sequence); i++ {
      nextTo = sequence[i]
      sum += GetLengthAfterSteps(nextFrom, nextTo, stepsLeft-1)
      nextFrom = sequence[i]
    }

    if bestSum < 0 || sum < bestSum {
      bestSum = sum
    }
  }

  StepCache[h] = bestSum

  return bestSum
}
