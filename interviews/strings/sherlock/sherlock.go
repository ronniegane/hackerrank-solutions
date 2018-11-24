package sherlock

func isValid(s string) string {
	// valid if all characters of the string appear
	// the same number of times

	// Also valid if just *one* appearance of a character can
	// be removed to satisfy the above requirement

	// First approach: build a frequency map
	freq := map[rune]int{}
	for _, ch := range s {
		if _, ok := freq[ch]; ok {
			freq[ch]++
		} else {
			freq[ch] = 1
		}
	}

	// Make a second map: frequency of letter frequencies
	counts := map[int]int{}
	for _, v := range freq {
		if _, ok := counts[v]; ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}

	// String is only valid if the second map
	// - has one entry, OR
	// - has two entries, AND one of them
	//		- has value 1 (one occurrence) AND key one more than the other key, OR
	//		- has value 1 (one occurrence) AND key == 1 (one appearance, can delete)

	if len(counts) > 2 {
		return "NO"
	} else if len(counts) == 1 {
		return "YES"
	} else {
		for key, val := range counts {
			if val == 1 {
				// Only one instance of this key/count
				if _, ok := counts[key-1]; ok {
					return "YES"
				} else if key == 1 {
					return "YES"
				} else if f, ok := counts[key+1]; ok && f == 1 {
					return "YES"
				}
				return "NO"
			}
		}
	}
	return "NO"
}
