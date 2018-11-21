package toys

import (
	"sort"
)

// Returns the maximum number of toys possible to buy
// given a slice of prices and a budget (k)
func max(prices []int32, k int32) int32 {

	// Naive approach: sort prices, and add starting from the cheapest

	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	for i := 0; i < len(prices); i++ {
		if k < prices[i] {
			return int32(i)
		}
		k -= prices[i]
	}

	return int32(len(prices))

}

// Is there a faster way to do it without having to sort the array first?
