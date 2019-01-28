package manipulation

import "sort"

func arrayManipulation(n int32, queries [][]int32) int64 {
	var max int64
	// initialize the array
	arr := make([]int64, int(n))
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := q[2]
		for i := int(a) - 1; i <= int(b)-1; i++ {
			arr[i] += int64(k)
			if arr[i] > max {
				max = arr[i]
			}
		}
	}
	return max
}

func mapBased(n int32, queries [][]int32) int64 {
	var max int32
	// This uses a map to avoid large empty arrays
	sums := map[int32]int32{}
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := q[2]
		for i := a - 1; i <= b-1; i++ {
			if v, ok := sums[i]; ok {
				sums[i] = v + k
			} else {
				sums[i] = k
			}

			if sums[i] > max {
				max = sums[i]
			}
		}
	}
	return int64(max)
}

func diffBased(n int32, queries [][]int32) int64 {
	var max, curr int64
	// Here we record only the differences (+k and -k) at the ends of the range
	arr := make([]int64, int(n))
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := int64(q[2])
		arr[int(a)-1] += k
		if int(b) < len(arr) {
			arr[int(b)] -= k
		}
	}
	for _, v := range arr {
		curr += v
		if curr > max {
			max = curr
		}
	}
	return max
}

func diffMap(n int32, queries [][]int32) int64 {
	// using differences and a map-based approach
	var max, curr int64
	// Here we record only the differences (+k and -k) at the ends of the range
	sums := map[int32]int64{}
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := int64(q[2])
		if v, ok := sums[a]; ok {
			sums[a] = v + k
		} else {
			sums[a] = k
		}
		if b < n {
			if v, ok := sums[b+1]; ok {
				sums[b+1] = v - k
			} else {
				sums[b+1] = -k
			}
		}
	}
	// Put the keys of the map into an array so we can sort it
	var keys = make([]int32, len(sums))
	i := 0
	for k := range sums {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(a, b int) bool {
		if keys[a] < keys[b] {
			return true
		}
		return false
	})
	for _, k := range keys {
		curr += sums[k]
		if curr > max {
			max = curr
		}
	}

	return max
}
