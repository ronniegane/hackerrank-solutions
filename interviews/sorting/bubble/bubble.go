package bubble

// Sort sorts and input array using bubble sort algorithm
// returns number of swaps required, first element in the array,
// and last element in the array
func Sort(a []int32) (int, int, int) {
	var swaps int
	min := a[0]
	max := a[0]
	for i, v := range a {
		if max < a[i] {
			max = a[i]
		}
		if min > a[i] {
			min = a[i]
		}
		// Count the number of elements to the right of v
		// that are lower than v
		for _, x := range a[i:] {
			if x < v {
				swaps++
			}
		}
	}
	return swaps, int(min), int(max)
}
