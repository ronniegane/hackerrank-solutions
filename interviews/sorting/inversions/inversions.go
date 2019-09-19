package inversions

// Count the number of swaps that merge sort will have to make
// to order an array

func naiveCount(arr []int32) int64 {
	// nested loop algo - will be very inefficient O(n^2)
	var count int64
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	return count
}

func loopMap(arr []int32) int64 {
	// still kind of a nested loop, but looping over a map we have built up of the count of each value seen so far
	// so will be O(N*K) where there are K distinct numbers in the array
	var count int64
	seen := map[int32]int64{}

	for _, x := range arr {
		// add to the map if not already seen
		if _, ok := seen[x]; !ok {
			seen[x] = 1
		} else {
			seen[x]++
		}
		// conut the number of numbers already seen that are larger than x
		for key, val := range seen {
			if key > x {
				count += val
			}
		}
	}

	return count
}

// can we do something similar to the triplets problem?
// have a left map and a right map as we go through?

// is it faster just to implement merge sort?
// https://www.geeksforgeeks.org/counting-inversions/ suggests it is
// sort arrays using merge sort.
// at each merge step, add to total if left number > right number
// if a[i] > b[j], then all the numbers a[i], a[i+1], a[i+2]... will be > b[j].
// so the number of inversions will increase by len(a)-i

func mergeCount(arr []int32) int64 {
	if len(arr) < 2 {
		return 0
	}
	sorted, total := mergeHelper(arr[:len(arr)/2], arr[len(arr)/2:])
	_ = sorted
	return total
}

func mergeHelper(left, right []int32) ([]int32, int64) {
	var count, leftCount, rightCount int64
	// recurse if needed
	if len(left) > 1 {
		left, leftCount = mergeHelper(left[:len(left)/2], left[len(left)/2:])
	}
	if len(right) > 1 {
		right, rightCount = mergeHelper(right[:len(right)/2], right[len(right)/2:])
	}
	count += leftCount
	count += rightCount
	// merge step
	out := []int32{}
	var i, j int
	for i < len(left) || j < len(right) {
		if i == len(left) {
			// left array empty
			out = append(out, right[j:]...)
			break
		}
		if j == len(right) {
			// right array empty
			// every element remaining in left array is > right array
			// total number of inversions is (len - i) + (len - i + 1) + ... 1
			// = 1/2 (len - i)^2
			count += int64((len(left) - i) * (len(left) - i) / 2)
			out = append(out, left[i:]...)
			break
		}
		if left[i] > right[j] {
			// inversion found
			count += int64(len(left) - i)
			out = append(out, right[j])
			j++
		} else {
			out = append(out, left[i])
			i++
		}
	}

	return out, count

}
