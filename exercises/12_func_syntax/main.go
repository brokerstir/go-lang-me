package main

import (
	"fmt"
)

func main() {
	foo()
	greet("World")
	ans := fortyTwo(42)
	fmt.Println(ans)
	x, y := lampin("Col' Lampin'", 1996)
	fmt.Println(x)
	fmt.Println(y)

}

// Function has following syntax
// func (r receiver) identifier(parameters) (return) {...}
// Know diffence between parameters and arguments

func foo() {
	fmt.Println("Hello from Bar")
}

// Everything in Go is passed by value
func greet(s string) {
	fmt.Println("Hello, ", s)
}

// Return a value
func fortyTwo(x int) string {
	return fmt.Sprint("The answer is ", x)
}

func lampin(a string, b int) (string, bool) {
	p := fmt.Sprint("Yo, we ", a, " like ", b, " Boy!")
	q := true
	return p, q
}
