package luck

import (
	"sort"
)

func sortLuck(k int32, contests [][]int32) int32 {
	// Sort contests by first element (luck) descending
	// Lose the first k important contents encountered
	var lost, luck int32
	sort.Slice(contests, func(i, j int) bool {
		return contests[i][0] > contests[j][0]
	})
	for _, con := range contests {
		if con[1] == 1 {
			if lost < k {
				// Important high value contest, should lose
				lost++
				luck += con[0]
			} else {
				// Can't lose any more important contests; must win
				luck -= con[0]
			}
		} else {
			luck += con[0]
		}
	}
	return luck
}

func onePass(k int32, contests [][]int32) int32 {
	// Keep a list of the k smallest important contests encountered
	// Add to and remove from the list as we iterate
	// If we add to the list, subtract luck
	// If we remove from the list, add luck
	var luck int32
	var smallest = make([]int32, k)
	for _, c := range contests {
		if c[1] == 1 {
			// Important
			if len(smallest) < int(k) {
				// Just add to list
			} else if c[0] < smallest[0] {
				// Smaller than the largest lost important contest
				// Bump that off and add this one
				// Find the right place to insert this entry
				// To maintain list sorting
				insertIdx := int(k)
				for i, v := range smallest {
					if v < c[0] {
						insertIdx = i - 1
					}
				}
				_ = insertIdx
				// smallest = append(smallest[1:insertIdx], c[0], smallest[insertIdx:]...)
				// Is re-sorting the list the only way to get this
			}
		}
		// Not important, just lose the contest
		luck += c[0]
	}
	return luck
}
