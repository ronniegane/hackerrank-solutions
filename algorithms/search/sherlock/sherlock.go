package sherlock

// Complete the balancedSums function below.
func balancedSums(arr []int32) string {
	var leftSum, rightSum int32

	// Get the starting right sum
	for _, v := range arr {
		rightSum += v
	}

	// Iterate through the array
	// At each step move the current value out of the right sum,
	// and move it into the left sum in the next step
	for _, v := range arr {
		rightSum -= v
		if rightSum == leftSum {
			return "YES"
		}
		leftSum += v
	}
	return "NO"
}
