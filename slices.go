package main

import "fmt"

func main() {
    a := [5]int{0, 1, 2, 3, 4}
    s1 := a[0:3]
    s2 := a[2:5]
    fmt.Println(a, s1, s2)
}