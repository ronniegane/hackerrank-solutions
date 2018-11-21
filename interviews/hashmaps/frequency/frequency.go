package frequency

// Queries processes the list of queries provided
func Queries(q [][]int32) []int32 {
	var res []int32
	var counts = map[int32]int32{}
	for _, v := range q {
		switch v[0] {
		case 1:
			// Insert statement
			if _, ok := counts[v[1]]; ok {
				counts[v[1]]++
			} else {
				counts[v[1]] = 1
			}
		case 2:
			// Delete statement
			// Reduce count (won't go below zero)
			if count, ok := counts[v[1]]; ok && count > 0 {
				counts[v[1]]--
			}
		case 3:
			// Frequency lookup statement
			var found int32
			for _, val := range counts {
				if val == v[1] {
					found = 1
					break
				}
			}
			res = append(res, found)
		}
	}

	return res
}

// Naive approach gets terminated due to timeout
// Potential slow points: iterating over all map entries
// Original: 105490 ns/op with long queue (230 entries)

// Frequency map is about 5x faster

// Plans for change: store frequency map as well
func FreqMap(q [][]int32) []int32 {
	var res []int32
	var counts = map[int32]int32{}
	var frequencies = map[int32]map[int32]bool{}
	for _, v := range q {
		entry := v[1]
		switch v[0] {
		case 1:
			// Insert statement
			if count, ok := counts[entry]; ok {
				// Move to the next highest frequency map entry
				delete(frequencies[count], entry)
				counts[entry]++
				if frequencies[count+1] == nil {
					// Create the map entry
					frequencies[count+1] = make(map[int32]bool)
				}
				frequencies[count+1][entry] = true
			} else {
				counts[entry] = 1
				if frequencies[1] == nil {
					// Create the map entry
					frequencies[1] = make(map[int32]bool)
				}
				frequencies[1][entry] = true
			}
		case 2:
			// Delete statement
			// Reduce count (won't go below zero)
			if count, ok := counts[v[1]]; ok && count > 0 {
				delete(frequencies[count], entry)
				counts[v[1]]--
				if count > 1 {
					frequencies[count-1][entry] = true
				}
			}
		case 3:
			// Frequency lookup statement
			if len(frequencies[entry]) > 0 {
				res = append(res, 1)
			} else {
				res = append(res, 0)
			}
		}
	}

	return res
}
