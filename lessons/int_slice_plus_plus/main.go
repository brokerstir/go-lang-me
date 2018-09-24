package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	//index out of range error
	//mySlice[1] = 8
	fmt.Println(mySlice[0])
	//mySlice = mySlice[0] + 1
	//mySlice += 1
	mySlice[0]++ // idiomatic, concise way
	fmt.Println(mySlice[0])
}
