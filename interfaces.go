package main

import (
	"fmt"
	"github.com/treehouse-projects/go-intro/parts"
)

type Part interface {
  Specs() string
  Price() string
}

func Summary(part Part) string {
  return part.Specs() + "\n" + part.Price()
}

func main() {
  catalog := []Part{}
	catalog = append(catalog, parts.Monitor{"HDMI", "1080p", 249.99})
  catalog = append(catalog, parts.HardDrive{"SSD", "SATA", 149.99})
  catalog = append(catalog, parts.Keyboard{108, "Mechanical", 99.99})
	for _, part := range catalog {
    fmt.Println(Summary(part))
	  fmt.Println("------------------------") 
  }	
}
