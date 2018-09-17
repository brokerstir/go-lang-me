package main

import "fmt"

func main() {

	foo(1, 2)
	aSlice := []int{1, 3, 5}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
