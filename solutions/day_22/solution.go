package day_22

import (
  "fmt"
  "strconv"

  "advent_of_code_2024/common"
)

const (
  dataFile = "input/day_22.txt"
)

func readInput() ([]int, error) {
  var nums []int
  err := common.ReadInput(dataFile, func(lines []string) error {
    for _, l := range lines {
      num, err := strconv.Atoi(l)
      if err != nil {
        return err
      }
      nums = append(nums, num)
    }

    return nil
  })

  return nums, err
}

func Mix(secret int, num int) int {
  return secret ^ num
}

func Prune(secret int) int {
  return secret % 16777216
}

func GenNewSecret(secret int) int {
  temp := secret * 64
  secret = Mix(secret, temp)
  secret = Prune(secret)

  temp = secret / 32
  secret = Mix(secret, temp)
  secret = Prune(secret)

  temp = secret * 2048
  secret = Mix(secret, temp)
  secret = Prune(secret)

  return secret
}

func Solve() error {
  nums, err := readInput()
  if err != nil {
    return err
  }
  windowsSum := map[string]int{}

  for _, initialSecret := range nums {
    secret := initialSecret
    prices := []int{secret % 10}
    var changes []int
    prevPrice := secret % 10
    for i := 0; i < 2000; i++ {
      secret = GenNewSecret(secret)
      price := secret % 10
      changes = append(changes, price-prevPrice)
      prices = append(prices, price)
      prevPrice = price
    }

    windowPrices := strToWindows(changesToString(changes), prices)
    for window, windowPrice := range windowPrices {
      windowsSum[window] += windowPrice
    }
  }

  bestSum := 0
  for _, sum := range windowsSum {
    if sum > bestSum {
      bestSum = sum
    }
  }

  fmt.Println(bestSum)

  return nil
}

func changesToString(changes []int) string {
  var res string
  for _, change := range changes {
    if change >= 0 {
      res += "+" + strconv.Itoa(change)
    } else {
      res += strconv.Itoa(change)
    }
  }

  return res
}

func strToWindows(str string, prices []int) map[string]int {

  windows := map[string]int{}
  for i := 0; i <= len(str)-7; i += 2 {
    window := str[i : i+8]
    if _, ok := windows[window]; !ok {
      price := prices[(i+8)/2]
      windows[window] = price
    }
  }

  return windows
}
