package minswaps

// Used to record visitation of array elements
type visit struct {
	val  int
	seen bool
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {

	var trackArray []visit
	for _, v := range arr {
		trackArray = append(trackArray, visit{int(v), false})
	}
	// traversing through array finding cycles
	var swapCount int
	for i, v := range trackArray {
		if v.seen {
			continue
		}

		// Mark this number as seen (note that we have to access the original array rather than v)
		trackArray[i].seen = true

		if i == v.val {
			// Number already in correct place
			continue
		}
		// Start tracking a cycle
		// Check the number sitting in v's "home"
		current := v
		for {
			next := trackArray[current.val-1]

			// Mark next as seen (have to update original array)
			trackArray[current.val-1].seen = true

			if next.val == v.val {
				// End of the cycle
				break
			}
			// Necessary swaps is (length of cycle) - 1
			swapCount++
			current = next
		}
	}

	return int32(swapCount)

}

// Original is 765 ns/op            1008 B/op          6 allocs/op

// Improved is 91.4 ns/op            32 B/op          1 allocs/op

// Variation just using indices without any special types
// Or array copying
func withoutTrackStruct(arr []int32) int32 {
	seen := make([]bool, len(arr))

	var swapCount int32

	for i, v := range arr {
		if seen[i] {
			continue
		}

		// Mark this number as seen
		seen[i] = true

		// Find cycles
		current := v
		for {
			nextIdx := int(current) - 1
			next := arr[nextIdx]
			seen[nextIdx] = true
			if next == v {
				// End of the cycle
				break
			}
			// Needed swaps is (length of cycle) - 1
			swapCount++
			current = next
		}
	}
	return swapCount
}
