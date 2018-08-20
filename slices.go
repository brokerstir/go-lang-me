package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    for i, v := range s {
      fmt.Println("element:", i, "value:", v)
    }
}