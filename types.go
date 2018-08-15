package main

import "fmt"

func main() {

  var wholeNumber int = 1
  var fractionalNumber float64 = 2.75
  var wholeNumber2 int = int(fractionalNumber)
  var fractionalNumber2 float64 = float64(wholeNumber)
  fmt.Println("wholeNumber2:", wholeNumber2)
  fmt.Println("fractionalNumber2:", fractionalNumber2)
  
}