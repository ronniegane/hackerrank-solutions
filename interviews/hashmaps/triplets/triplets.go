package triplets

// Find triplet sets of integers within an array
// that satisfy a geometric progression
// with the factor r
// AND have increasing indices in the array

// GeoTrips only works if the array is sorted
func GeoTrips(arr []int64, r int64) int64 {
	var count int64
	// Make a map of all unique integers and how often they appear
	lookup := map[int64]int64{}
	for _, v := range arr {
		if _, ok := lookup[v]; ok {
			lookup[v]++
		} else {
			lookup[v] = 1
		}
	}

	// Loop through the map looking for triplets
	for key, a := range lookup {
		// Special case: if r=1 we look for duplicates of this number
		if r == 1 && a >= 3 {
			// Number of increasing order combinations
			// = number of combinations / ways to order a 3-element set
			count += a * (a - 1) * (a - 2) / 6
		} else {
			// check for existence of the next two members in the triplet
			if b, ok1 := lookup[key*r]; ok1 {
				if c, ok2 := lookup[key*r*r]; ok2 {
					// Add the number of possible combinations
					count += a * b * c
				}
			}
		}
	}
	return count
}

func naiveCount(arr []int64, r int64) int64 {
	var count int64
	// naive looping approach
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] == arr[i]*r {
				for k := j + 1; k < len(arr); k++ {
					if arr[k] == arr[j]*r {
						count++
					}
				}
			}
		}
	}
	return count
}

func treeCount(arr []int64, r int64) int64 {
	/* Build up a tree starting from the last member of the array
	   each node points to other nodes that are:
	   - After it in the array, AND
	   - Value r times higher
	*/
	type node struct {
		index    int
		value    int64
		children []node
	}

	count := 0

	valueLookup := map[int64][]node{}

	for i := len(arr) - 1; i >= 0; i-- {
		// find values that work for this
		thisVal := arr[i]
		childNodes := valueLookup[arr[i]*r]
		newNode := node{i, thisVal, childNodes}

		if childNodes != nil {
			for _, child := range childNodes {
				// If these children have children of their own, we have a chain
				count += len(child.children)
			}
		}

		// Add this new node to the lookup map
		currentLookup := valueLookup[thisVal]
		if currentLookup == nil {
			valueLookup[thisVal] = []node{newNode}
		} else {
			valueLookup[thisVal] = append(valueLookup[thisVal], newNode)
		}

	}

	return int64(count)
}

func pathCount(arr []int64, r int64) int64 {
	var total int64
	/* Two maps:
	- value -> count of value seen so far
	- value -> total paths from this value x to x*r
	iterate from last element to first
	*/
	counts := map[int64]int64{}
	paths := map[int64]int64{}
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		// add to maps if it does not exist
		if _, ok := counts[val]; !ok {
			counts[val] = 0
		}
		if _, ok := paths[val]; !ok {
			paths[val] = 0
		}

		// If we find x*r exists, how many 3-number links can we make
		// x -> x*r -> x*r*r
		// starting from our current value?
		// the answer is the number of x*r -> x*r*r links seen so far
		if nextPaths, ok := paths[val*r]; ok {
			total += nextPaths
		}

		// check for the next number in the chain
		// the possible number of x -> x*r links increases
		// by the number of r*x values seen so far
		if nextCount, ok := counts[val*r]; ok {
			paths[val] += nextCount
		}

		counts[val]++
	}

	return total
}

func splitMap(arr []int64, r int64) int64 {
	var total int64

	// Keep two maps: left and right which show how many times numbers appear
	// to the left or right of the current index in the array.
	left := map[int64]int64{}
	right := map[int64]int64{}
	for _, v := range arr {
		if _, ok := right[v]; ok {
			right[v]++
		} else {
			right[v] = 1
		}
		left[v] = 0
	}

	// as we iterate through the array, we have our central number x.
	// this can make a chain of (x/r) -> x -> (x*r).
	// the number of possible chains involving this index
	// is count(x/r in left) * count(x*r in right)
	for _, x := range arr {
		right[x]--
		// only need to consider case when x is evenly divisible by r.
		if x%r == 0 {
			total += left[x/r] * right[x*r]
		}
		left[x]++
	}

	return total
}
