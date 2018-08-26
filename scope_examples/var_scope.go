package main

import (
	"fmt"
)

// Package variable
var a = "a"

func main() {
	var b = "b"
	{
		var c = "c"
		{
			var d = "d"
			fmt.Println(a, b, c, d)
		}
    // ERROR: "d" undefined
		fmt.Println(a, b, c, d)
	}
  // ERROR: "c", "d" undefined
	fmt.Println(a, b, c, d)
}
