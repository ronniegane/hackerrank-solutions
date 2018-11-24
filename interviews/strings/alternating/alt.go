package alt

// Count the minimum number of deletions in a string
// required to have no duplicate consecutive characters
func Count(s string) int32 {
	var current rune
	var count int32
	for _, ch := range s {
		if ch == current {
			count++
		} else {
			current = ch
		}
	}
	return count
}
