package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // update lowest price so far
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice // update max profit
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // Output: 5
}
