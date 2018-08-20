package main

import "fmt"

func main() {
    months := [3]string{"Apr", "May", "Jun"}
    salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
    for i := 0; i < len(months); i++{
        fmt.Println(months[i], salesByMonth[i])  
    }
}
