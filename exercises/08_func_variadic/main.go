/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/

package main

import "fmt"

func max(numbers ...int) int {
	fmt.Printf("%T \n", numbers)

	var biggest int
	for _, i := range numbers {
		if i > biggest {
			biggest = i
		}
	}
	return biggest
}

func main() {
	greatest := max(3, 0, 7, 9, 8, 98, 76, 73, 4)
	fmt.Println(greatest)
}
