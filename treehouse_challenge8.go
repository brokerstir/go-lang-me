package sales

func HalfPriceSale(prices map[string]float64) map[string]float64 {
  // YOUR CODE HERE
  for code, price := range prices {
  	prices[code] = price/2
  }
  return prices
}
