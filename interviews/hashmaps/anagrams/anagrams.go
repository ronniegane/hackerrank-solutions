package anagrams

import (
	"fmt"
	"sort"
	"strings"
)

// Sherlock finds all pairs of substrings in the string that are anagrams of each other
// string is between 2 and 100 chars
func Sherlock(s string) int32 {
	// Naive approach:
	// Find all substrings of s
	// if this substring exists in the map already, increase the count
	count := 0
	for l := 1; l < len(s); l++ {
		// Start with small substrings
		matches := map[string]int{}
		for i := 0; i < len(s)+1-l; i++ {
			sub := s[i : i+l]
			// Sort the characters in this substring alphabetically
			// so it will match any anagrams
			splitSub := strings.Split(sub, "")
			sort.Strings(splitSub)
			sortedSub := strings.Join(splitSub, "")
			if n, ok := matches[sortedSub]; ok {
				// If 2 matches already exist, we have 2 new possible pairs
				// and so on
				count += n
				matches[sortedSub]++
			} else {
				matches[sortedSub] = 1
			}
		}
	}
	return int32(count)
}

// 31619 ns/op            5920 B/op        180 allocs/op
// outer loop n times
// inner loop n-1 times
// inner loop includes a sort which is O(n log n)
// So approach is probably n^2 log n
// Is there a better way to account for anagrams?

// Signature creates an alphabetic signature (hash) for a substring
// Rather than sorting it.
func Signature(s string) int32 {
	// Naive approach:
	// Find all substrings of s
	// if this substring exists in the map already, increase the count
	count := 0
	for l := 1; l < len(s); l++ {
		// Start with small substrings
		matches := map[string]int{}
		for i := 0; i < len(s)+1-l; i++ {
			sub := s[i : i+l]
			// Sort the characters in this substring alphabetically
			// so it will match any anagrams
			sig := signature(sub)
			if n, ok := matches[sig]; ok {
				// If 2 matches already exist, we have 2 new possible pairs
				// and so on
				count += n
				matches[sig]++
			} else {
				matches[sig] = 1
			}
		}
	}
	return int32(count)
}

// Signature version is even slower than the sort version

func signature(s string) string {
	alpha := make([]int, 26)
	for _, ch := range s {
		ord := int(ch) - int('a')
		alpha[ord]++
	}
	return fmt.Sprint(alpha)
}
