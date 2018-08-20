package main

import "fmt"

func main() {
  ages := map[string]float64{"Alice": 12, "Carol": 10, "Bob": 9}
  // Like the array/slice syntax, but you can use any value of the type you specified for the keys.
  for _, age := range ages {
  	fmt.Println(age)
  }
}

