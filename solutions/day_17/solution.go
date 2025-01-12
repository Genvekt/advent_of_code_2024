package day_17

import (
  "fmt"
  "log"
  "math"
  "regexp"
  "strconv"
  "strings"

  "advent_of_code_2024/common"
)

const (
  dataFile  = "input/day_17.txt"
  registerA = "A"
  registerB = "B"
  registerC = "C"
)

type Program struct {
  registers         map[string]int
  instructionCursor int
  input             []int
  output            []string
}

func (p *Program) IsFinished() bool {
  return p.instructionCursor >= len(p.input)
}

func (p *Program) GetComboOperand(operand int) int {
  switch operand {
  case 0, 1, 2, 3:
    return operand
  case 4:
    return p.registers[registerA]
  case 5:
    return p.registers[registerB]
  case 6:
    return p.registers[registerC]
  default:
    log.Fatal("7 is reserved and must not be used as combo operand!")
    return 0
  }
}

func (p *Program) RunInstruction() {
  instruction := p.input[p.instructionCursor]
  operand := p.input[p.instructionCursor+1]

  switch instruction {

  case 0:
    nominator := float64(p.registers[registerA])
    denominator := math.Pow(2, float64(p.GetComboOperand(operand)))

    result := int(nominator / denominator)
    p.registers[registerA] = result
    p.instructionCursor += 2

  case 1:
    p.registers[registerB] = p.registers[registerB] ^ operand
    p.instructionCursor += 2

  case 2:
    p.registers[registerB] = p.GetComboOperand(operand) % 8
    p.instructionCursor += 2

  case 3:
    if p.registers[registerA] == 0 {
      p.instructionCursor += 2
    } else {
      p.instructionCursor = operand
    }

  case 4:
    p.registers[registerB] = p.registers[registerB] ^ p.registers[registerC]
    p.instructionCursor += 2

  case 5:
    result := p.GetComboOperand(operand) % 8
    p.output = append(p.output, fmt.Sprintf("%d", result))
    p.instructionCursor += 2

  case 6:
    nominator := float64(p.registers[registerA])
    denominator := math.Pow(2, float64(p.GetComboOperand(operand)))

    result := int(nominator / denominator)
    p.registers[registerB] = result
    p.instructionCursor += 2

  case 7:
    nominator := float64(p.registers[registerA])
    denominator := math.Pow(2, float64(p.GetComboOperand(operand)))

    result := int(nominator / denominator)
    p.registers[registerC] = result
    p.instructionCursor += 2
  }
}

func readInput() (*Program, error) {
  program := &Program{
    instructionCursor: 0,
    registers: map[string]int{
      registerA: 0,
      registerB: 0,
      registerC: 0,
    },
  }

  err := common.ReadInput(dataFile, func(lines []string) error {
    registerRegexp := regexp.MustCompile("^Register (.+): (.+)$")
    var lastLineId int
    for lineId, line := range lines {
      if line == "" {
        lastLineId = lineId + 1
        break
      }

      parts := registerRegexp.FindAllStringSubmatch(line, -1)

      registerName := parts[0][1]
      if _, present := program.registers[registerName]; !present {
        continue
      }
      registerValue, _ := strconv.Atoi(parts[0][2])

      program.registers[registerName] = registerValue
    }

    dataParts := strings.Split(strings.ReplaceAll(lines[lastLineId], "Program: ", ""), ",")
    for _, dataPart := range dataParts {
      data, _ := strconv.Atoi(dataPart)
      program.input = append(program.input, data)
    }

    return nil
  })

  return program, err
}

func Solve() error {
  program, err := readInput()
  if err != nil {
    return err
  }
  fmt.Printf("%+v\n", *program)

  for !program.IsFinished() {
    program.RunInstruction()
  }

  fmt.Printf("%+v\n", *program)
  fmt.Printf("%s\n", strings.Join(program.output, ","))
  return nil
}

func SolvePart2() error {
  targetOut := []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 0, 0, 3, 5, 5, 3, 0}

  step := 15
  a := 0

  fmt.Println(solveStep(a, step, targetOut))

  return nil
}

func solveStep(a int, step int, targetOut []int) (int, bool) {
  if step == -1 {
    return a >> 3, true
  }

  for i := 0; i < 8; i++ {
    candidateA := a + i

    b := candidateA % 8
    b = b ^ 1
    c := candidateA >> b
    b = b ^ 5
    b = b ^ c
    out := b % 8

    if out == targetOut[step] {
      result, found := solveStep(candidateA<<3, step-1, targetOut)
      if found {
        return result, found
      }
    }
  }

  return 0, false
}
