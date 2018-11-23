package abs

import (
	"sort"
)

func minAbsDiff(arr []int32) int32 {
	// Find the minimum absolute difference between any two numbers
	// in the array

	// First approach: sort, then compare consecutive pairs
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})

	// Iterate comparing pairs
	min := arr[1] - arr[0] // With array sorted ascending, this will be >= 0
	for i := 0; i < len(arr)-1; i++ {
		next := arr[i+1] - arr[i]
		if next < min {
			min = next
		}
		if min == 0 {
			// If we find zero diff we can stop early
			break
		}
	}

	return min
}
