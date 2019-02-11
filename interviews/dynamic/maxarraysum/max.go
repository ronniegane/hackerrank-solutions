package maxarraysum

import (
	"fmt"
)

/*
Given an array of integers, find the subset of
_non-adjacent_ elements that gives the maximum sum.

length of array (n) <= 10^5
*/

func maxSubsetSum(arr []int32) int32 {
	// Iterate through array, keeping two sums:
	// - max sum including previous element
	// - max sum excluding previous element
	maxIn := arr[0]
	var maxEx int32

	for i := 1; i < len(arr); i++ {
		// At each step:
		// The maximum including the current element becomes
		// (sum excluding previous element) + current element value
		// The maximum excluding the current element becomes
		// MAX(sum excluding previous element, sum including previous)

		// Doing assignment simultaneously to avoid need for a temp variable
		maxIn, maxEx = maxEx+arr[i], maxInt(maxIn, maxEx)
	}
	return maxInt(maxIn, maxEx)
}

// helper function to find maximum of two int32's
func maxInt(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Below algo doesn't work
func maxSubsetSumGreedy(arr []int32) int32 {
	var count int32
	// Greedy algorithm:
	// start with the first postive number
	var i int
	for arr[i] <= 0 {
		i++
	}
	// consider this number or its neighbour, choose the highest
	for i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			fmt.Printf("adding index %d (= %d) to count\n", i, arr[i])
			count += arr[i]
			fmt.Println("Count is", count)
			i += 2
		} else {
			fmt.Printf("adding index %d (= %d) to count\n", i+1, arr[i+1])
			count += arr[i+1]
			fmt.Println("Count is", count)
			i += 3
		}
	}
	return count
}

func maxSubsetSumBrute(arr []int32) int32 {
	// horrible brute-force approach
	// calculate all possible subsets of numbers
	// find max
	// is brute-force even that easy to write here?
	return 0
}
