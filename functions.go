package main

import "fmt"

func main() {
  square(3)
  add(2, 4)
  subtract(10, 3)
}

func square(number int) {
  fmt.Println(number * number)
}

func add(a float64, b float64) {
	fmt.Println(a + b)
}

func subtract(a, b float64) {
	fmt.Println(a - b)
}