package main

import "fmt"

func main() {
    months := [3]string{"Apr", "May", "Jun"}
    salesByMonth := [3]float64{1710.26, 2245.97, 3032.40}
    for i, month := range months{
        fmt.Println(month, salesByMonth[i])  
    }
}