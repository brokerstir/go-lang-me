package main

import (
  "fmt"
  "os"
)

func main() {
  fileInfo, _ := os.Stat("existent.txt")
  fmt.Println(fileInfo.Size())
  fileInfo, _ = os.Stat("nonexisten.txt")
  fmt.Println(fileInfo.Size())
}