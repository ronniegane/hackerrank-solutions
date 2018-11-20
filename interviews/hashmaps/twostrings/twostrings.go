package twostrings

// Common finds a common string between the two input strings
func Common(s1, s2 string) bool {

	// make sets of the unique characters in the first string
	uniques := make(map[rune]bool)
	for _, ch := range s1 {
		uniques[ch] = true
	}
	// Iterate over the characters of the second string and check if
	// they exist in the set
	for _, ch := range s2 {
		if uniques[ch] {
			return true
		}
	}
	// Iterates over s1 and s2 once each, so is O(n)
	return false
}
