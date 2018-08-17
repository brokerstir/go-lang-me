package sales

// YOUR CODE HERE
func sales () {
	CalculateTax(100, 0.8)
}

func CalculateTax(total, taxRate float64) (tax float64) {
	tax = total * taxRate
  return
}