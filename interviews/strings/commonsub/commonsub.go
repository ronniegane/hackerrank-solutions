package commonsub

// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

// different from substrings: don't have to be consecutive characters, just in retained order

// Largest string you can make by simple deletion from both string A and string B

// Solvable in polynomial time by dynamic programming

// In this example we are only required to store the length, not the actual subsequence itself
// Although knowing the actual subsequence might be good for testing

func longestSub(s1, s2 string) int32 {

	// Approach suggested is to break down into smaller, simpler subproblems
	// solutions to subproblems are reused, so are memoized

	var count int32
	// If the strings start with the same element
	// remove it, repeat until no common element found or we run out
	for len(s1) > 0 && len(s2) > 0 && s1[0] == s2[0] {
		s1 = s1[1:]
		s2 = s2[1:]
		count++
	}

	if len(s1) == 0 || len(s2) == 0 {
		return count
	}

	// then the result is the longer of these two options:
	count += max(longestSub(s1[1:], s2), longestSub(s1, s2[1:]))

	// This is a recursive approach with branching so may not be very efficient
	return count
}

func arraySub(s1, s2 string) int32 {
	// This approach, suggested in the wikipedia article, uses a 2D array to keep track of the smaller subsets

	m := len(s1)
	n := len(s2)

	var arr = make([][]int32, m+1)
	// initialise to zeroes
	for i := 0; i <= m; i++ {
		arr[i] = make([]int32, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = max(arr[i][j-1], arr[i-1][j])
			}
		}
	}
	return arr[m][n]
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
