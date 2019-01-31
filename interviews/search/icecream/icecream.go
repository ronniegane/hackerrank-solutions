package icecream

// cost of icecream flavours i...n stored in an unordered array
// need to get two unique indices that total to the cost

func whatFlavors(cost []int32, money int32) {
	//

}

func flavorIndices(cost []int32, money int32) [2]int {
	// Make a hashmap from an amount to an array index
	lookup := map[int32]int{}
	for i, v := range cost {
		lookup[v] = i // will overwrite but I don't think it affects correctness
	}
	for i, v := range cost {
		// Get the amount left over after buying this flavor
		remain := money - v
		// Check if there is any flavor that costs the remaining amount
		if idx, ok := lookup[remain]; ok && idx != i {
			// adjust results to 1-based array
			return [2]int{i + 1, idx + 1}
		}
	}
	return [2]int{0, 0}
}
