package main

import "fmt"

func main() {
	myNumber := 2.6
	halve(myNumber)                              // Does nothing!
	fmt.Println("myNumber in 'main':", myNumber) // Prints 2.6!
}

func halve(number float64) {
	number = number / 2
	fmt.Println("number in 'halve':", number) // Prints 1.3, but change is only effective here!
}