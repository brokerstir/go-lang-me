/* Create a program that prints to the terminal asking for a user to enter a
small number and a larger number. Print the remainder of the larger number
divided by the smaller number. */

package main

import "fmt"

func main() {
	var small int
	var big int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a big number: ")
	fmt.Scan(&big)
	result := big % small
	fmt.Println(result)

}
