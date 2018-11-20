package newyear

// minimumBribes prints out the number of position swaps needed
// to reach the final configuration
func minimumBribes(q []int32) {

}

// IminiBribes is similar but returns the number, or -1 if impossible
func minBribes(q []int32) int {
	var moves int
	for i, v := range q {
		// A number cannot be more than 2 indices ahead of home
		if i < int(v)-3 { // Subtract 3 as array index starts at 0
			return -1
		}
		// count the number of elements to the right of N that are smaller than N
		for _, x := range q[i:] {
			if x < v {
				moves++
			}
		}

	}
	return moves
}

func revBribes(q []int32) int {
	var moves int
	// Start from the end of the queue, count how many elements
	// are in front of this one (between it and its home)
	// with value > N
	for i := len(q) - 1; i >= 0; i-- {
		// A number cannot be more than 2 indices ahead of home
		if i < int(q[i])-3 {
			return -1
		}

		for j := max(int(q[i])-2, 0); j < i; j++ {
			if q[i] < q[j] {
				moves++
			}
		}
	}
	return moves
}

// minBribes is O(n^2) probably, outer loop is O(n) and inner loop
// is n, n-1, n-2,....
// Limits are t <= 10 (number of test cases)
// and n <= 10^5 (length of queue)

func minBribesNaive(q []int32) int {
	// just compare number to its index
	var moves int
	for i, v := range q {
		if i < int(v)-3 {
			return -1
		}
		if i < int(v)-1 {
			// element has moved from its place
			moves += (int(v) - 1) - i
		}
	}
	return moves
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
