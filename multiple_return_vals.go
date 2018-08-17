package main

import (
  "fmt"
  "math"
  "log"
)

func main() {
  squareRoot, err := squareRoot(-1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(squareRoot)  
}

func squareRoot(x float64) (float64, error) {
  if x < 0 {
    return 0, fmt.Errorf("can't take square root of negative number")  
  }
  return math.Sqrt(x), nil
}