package golf

import (
	"fmt"
)

func ShotsDescription(par int, shots int) string {
  shotsOverPar := shots - par
  // YOUR CODE HERE
  switch shotsOverPar {
  // If the shotsOverPar variable is 1, return "bogey".
  case 1:
  	fmt.Println("bogey")
  	return "bogey"
  // If it's 0, return "par".
  case 0:
  	fmt.Println("par")
    return "par"
  // For -1 return "birdie".
  case -1:
  	fmt.Println("birdie")
    return "birdie"
  // For -2 return "eagle".
  case -2:
    fmt.Println("eagle")
    return "eagle"
  // For all other cases, return "no description".
  default:
  	fmt.Println("no description")
    return "no description"
  }
}