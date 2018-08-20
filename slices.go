package main

import "fmt"

func main() {
    a := [5]int{0, 1, 2, 3, 4}
    s1 := a[0:3]
    s2 := a[2:5]
    s1 = s1[0:4]
    s2 = s2[0:4]
    fmt.Println(a, s1, s2)
}