package leaderboard

// First solution returns correct for most test cases, but
// fails on some of the large (20,000 length array) ones.
// Not timing out, just wrong.

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	// Build a map of scores to positions?
	// Or (without maps)
	// Make an ordered array of scores with duplicates removed
	// Then simply find position of alice in that array with binary search
	var nonDupeScores = []int32{}
	var currentScore int32
	for _, s := range scores {
		if s != currentScore {
			// New score
			nonDupeScores = append(nonDupeScores, s)
		}
		currentScore = s
	}

	var positions = make([]int32, len(alice))
	for i, a := range alice {
		// Find the position that alice belongs in the scores (log n for binary search)
		positions[i] = findPosition(a, nonDupeScores) + 1
	}

	// Position is the position of that score + 1
	return positions
}

func findPosition(val int32, slice []int32) int32 {
	// Returns the position that value should be inserted at for binary search
	// Expects array sorted descending

	upper := len(slice) - 1
	if val < slice[upper] {
		return int32(upper + 1)
	}
	if val > slice[0] {
		return 0
	}
	lower := 0
	var mid int
	for upper != lower {
		mid = (upper + lower) / 2
		if slice[mid] == val {
			return int32(mid)
		} else if slice[mid] < val {
			upper = mid
		} else {
			lower = mid + 1
		}
	}
	return int32(upper)
}
