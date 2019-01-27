package pairs

import (
	"sort"
)

// Pairs returns the number of pairs in an array that
// have a difference of k
func pairs(k int32, arr []int32) int32 {
	// sort array first (n lg n)
	// then for each of n values we can binary search
	// for another element that is k different
	var count int32

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i, v := range arr {
		if binarySearch(v+k, arr[i+1:]) {
			count++
		}
	}
	return count
}

/*
With array copying and conversion from int32 -> int
BenchmarkNaivePairs_10-4        10000000               137 ns/op
BenchmarkNaivePairs_20-4         3000000               452 ns/op
BenchmarkNaivePairs_40-4         1000000              1837 ns/op
BenchmarkNaivePairs_80-4          200000              6618 ns/op
BenchmarkPairs_10-4              2000000               579 ns/op
BenchmarkPairs_20-4              1000000              1361 ns/op
BenchmarkPairs_40-4               500000              2914 ns/op
BenchmarkPairs_80-4               200000              6131 ns/op

without copying array
BenchmarkNaivePairs_10-4        10000000               154 ns/op
BenchmarkNaivePairs_20-4         3000000               462 ns/op
BenchmarkNaivePairs_40-4         1000000              1855 ns/op
BenchmarkNaivePairs_80-4          200000              7448 ns/op
BenchmarkPairs_10-4              5000000               330 ns/op
BenchmarkPairs_20-4              2000000               798 ns/op
BenchmarkPairs_40-4              1000000              1481 ns/op
BenchmarkPairs_80-4               500000              3187 ns/op

without copying, array will be mutated and thus in order each time after the first
which may affect benchmarks
*/

func binarySearch(target int32, arr []int32) bool {
	if len(arr) == 0 {
		return false
	}
	idx := len(arr) / 2
	switch {
	case arr[idx] == target:
		return true
	case arr[idx] > target:
		return binarySearch(target, arr[:idx])
	default:
		return binarySearch(target, arr[idx+1:])
	}
}

// Building a map and iterating over it should be O(n)
// But maybe more setup time, and more storage space
func withSet(k int32, arr []int32) int32 {
	var count int32
	// Build a set / map with a + k for every item in arr
	var offsets = map[int32]bool{}

	for _, v := range arr {
		offsets[v+k] = true
	}

	// for every value b in arr, check if it is in the set created earlier
	for _, b := range arr {
		if _, ok := offsets[b]; ok {
			count++
		}
	}

	return count
}

// 0 < k < 10^9
// 2 < n < 10^5
// naivePairs checks every possible combination, O(n^2)
func naivePairs(k int32, arr []int32) int32 {
	var count int32
	for i, v := range arr {
		for _, u := range arr[i:] {
			if u-v == k || v-u == k {
				count++
			}
		}
	}
	return count
}
