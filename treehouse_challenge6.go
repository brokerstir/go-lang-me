package utility

import "fmt"

func First(slice []int) (int, error) {
  // YOUR CODE HERE
  if len(slice) == 0 {
  	return 0, fmt.Errorf("Slice is empty")
  } else {
  	return slice[0], nil
  }  
}