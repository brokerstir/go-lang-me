package main

import "fmt"

func main() {
	var aValue float64 = 1.23
	var aPointer *float64 = &aValue
	fmt.Println("aPointer:", aPointer)
	fmt.Println("*aPointer:", *aPointer)
}