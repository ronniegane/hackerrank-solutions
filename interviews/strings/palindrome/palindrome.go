package palindrome

/*
A string is said to be a special palindromic string if
either of two conditions is met:

    All of the characters are the same, e.g. aaa.
	All characters except the middle one are the same, e.g. aadaa.

*/

// n is the length of s
// all characters are lowercase alphabet
func substrCount(n int32, s string) int64 {
	var count int64
	// single char substrings are all palindromic
	count += int64(n)

	// Check all possible substrings of length 2...n
	for l := 2; l <= int(n); l++ {
		for start := 0; start+l <= int(n); start++ {
			// Check if this string is palindromic
			// Not a normal palindrome, all characters have to be the SAME
			char := s[start] // All characters must match this
			// Ignore the middle character if length is odd
			count++
			for i, j := start, start+l-1; i < j; i, j = i+1, j-1 {
				if s[i] != char || s[j] != char {
					// not palindromic
					count--
					break
				}
			}
		}
	}
	return count
}

type letterChain struct {
	ch rune
	l  int64
}

// Original approach too inefficient
func substrChain(n int32, s string) int64 {
	var count int64

	// Record letters and the chain length of these
	chains := []letterChain{}

	// Counts of pure one-letter series can be calculated in one pass
	var chain int64
	for i, v := range s {
		chain++
		// When we reach a letter boundary, add the number of possible
		// palindromes made by this string to the count

		if i == int(n)-1 || rune(s[i+1]) != v {

			chains = append(chains, letterChain{v, chain})

			chain = 0
		}
	}

	// Iterate over the list of chains
	for i, v := range chains {
		// For chain length x, there are 1 + 2 + 3 +... + x
		// Arithmetic sequence, sum = x * (1 + x) / 2
		count += v.l * (v.l + 1) / 2

		// special case handling of palindromes with middle letter different
		// Only possible if chain = 1
		// and neighbouring chains use the same character
		if v.l == 1 &&
			i > 0 && i < len(chains)-1 &&
			chains[i-1].ch == chains[i+1].ch {
			// here we can make as many palindromes as the length
			// of the smalles neighbour chain
			count += min(chains[i-1].l, chains[i+1].l)
		}
	}

	return count
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
