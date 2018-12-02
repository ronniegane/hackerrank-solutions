package sparse

// https://www.hackerrank.com/challenges/sparse-arrays/problem

func match(strings []string, queries []string) []int32 {
	ret := make([]int32, len(queries))
	// Naive approach: iterate over queries and strings
	// O(q * s)
	var count int32
	for i, q := range queries {
		count = 0
		for _, s := range strings {
			if q == s {
				count++
			}
		}
		ret[i] = count
	}
	return ret
}

func matchMap(strings []string, queries []string) []int32 {
	ret := make([]int32, len(queries))
	// Make a map by
	lookup := map[string]int32{}
	for _, s := range strings {
		lookup[s]++
	}

	for i, q := range queries {
		ret[i] = lookup[q]
	}

	return ret
}
