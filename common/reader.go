package common

import (
  "bufio"
  "os"
)

func ReadInput(fileName string, processingFunc func(lines []string) error) error {
  file, err := os.Open(fileName)
  if err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  if err := processingFunc(lines); err != nil {
    return err
  }

  return nil
}
