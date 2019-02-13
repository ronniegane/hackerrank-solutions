package buildings

// length of array is up to 10^5
// made an incorrect assumption at first that we were including
// all buildings from i to len(h): h[i:]
// when it is actually any group of consecutive buildings,
// ie i to k : h[i:k]
// this complicates the issue because there are O(n^2) such sets

// Memoing with triangle array of mins is still too slow
// works for the 10^4 case but not the 10^5 case
// it's set in the "stacks and queues" section of puzzles,
// so that implies we need a stack or queue to solve it efficiently.

func largestRectangle(h []int32) int64 {
	// how can we easily keep track of the minimum?
	// naive approach - recalculate it each time
	// will be something like O(n^2) since calculating
	// minimum is O(n)
	// and brute-force trying all possible sets of buildings
	// is also O(n^2)
	// so maybe worse than n^2

	// array of answers where i = index of start
	// and k = index of end
	// will be a triangular array since k >= i
	// equivalently, could have array of i and length
	/*
		eg for array [6 3 2 5 4 1]
		would look like (index, length)
		6 3 2 2 2 1
		3 2 2 2 1
		2 2 2 1
		5 4 1
		4 1
		1
		or (startindex, endindex)
		6 3 2 2 2 1
		- 3 2 2 2 1
		- - 2 2 2 1
		- - - 5 4 1
		- - - - 4 1
		- - - - - 1
	*/

	// build up 2D array of minimum heights
	// 0 index means length of 1
	// min(i, 0) = i
	// min(i, 1) = min(min(i, 0), i + 1)
	// min(i, k +1) = min(min(i, k), i+k+1)
	// in array notation:
	// arr[i][k+1] = min(arr[i][k], arr[i+k])
	var minHeightArray = make([][]int64, len(h))

	// initialise array
	for i, v := range h {
		minHeightArray[i] = make([]int64, len(h))
		minHeightArray[i][0] = int64(v)
	}
	// flesh out rest of array
	for i := range h { // here j is length of slice
		for j := 1; j < len(h)-i; j++ { // start index must be low enough to allow slice length
			minHeightArray[i][j] = minInt(minHeightArray[i][j-1], minHeightArray[i+j][0])
		}
	}

	var area, maxArea int64

	for i := range h {
		for j := 1; j < len(h)-i; j++ {
			area = int64(j+1) * minHeightArray[i][j]
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func minInt(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func min(arr []int32) int32 {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
