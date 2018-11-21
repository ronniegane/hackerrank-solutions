package fraud

import (
	"sort"
)

// BruteMedian recalculates the median at each step
func BruteMedian(expenditure []int32, d int32) int32 {
	// Naive approach: recalculate the median each time
	var count int32
	for i := int(d); i < len(expenditure); i++ {
		// Need to copy the slice to avoid sorting the underlying array
		history := append([]int32(nil), expenditure[i-int(d):i]...)

		// Calculate median
		sort.Slice(history, func(a, b int) bool {
			return history[a] > history[b]
		})

		var threshold int32
		// if d is even have to average
		if d%2 == 0 {
			threshold = (history[d/2-1] + history[d/2])
		} else {
			threshold = history[d/2] * 2
		}

		// Compare expenditure to notification level
		if expenditure[i] >= threshold {
			count++
		}
	}
	return count
}

// Brute force median times out on longer tests (200k entries, 40k median window)
// counting sort?
// expenditure is between 0 and 200 so that helps

// CountSort uses an array to store the counts of each integer seen so far
// This avoids the need to sort the median window
func CountSort(expenditure []int32, d int32) int32 {
	var count int32
	var countsArray [201]int32
	// Populate the counts array with the first slice
	for i := 0; i < int(d); i++ {
		countsArray[int(expenditure[i])]++
	}
	// Loop through the rest of the array
	for i := int(d); i < len(expenditure); i++ {
		// find the threshold
		var threshold, seen, first int32
		// For even number of samples, need to find average
		if d%2 == 0 {
			for j, v := range countsArray {
				seen += v
				if seen > d/2 {
					if first > 0 {
						threshold = int32(j) + first
					} else {
						threshold = int32(j) * 2
					}
					break
				} else if seen == d/2 {
					first = int32(j)
				}
			}
		} else {
			// for even number of samples, just need to find the (d/2 + 1)th sample
			for j, v := range countsArray {
				seen += v
				if seen > d/2 {
					threshold = int32(j) * 2
					break
				}
			}
		}
		if expenditure[i] >= threshold {
			count++
		}

		// Remove trailing number from median window
		countsArray[int(expenditure[i-int(d)])]--

		// Add this number to median window
		countsArray[int(expenditure[i])]++
	}
	return count
}
