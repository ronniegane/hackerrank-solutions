package abbreviation

// Can you make two strings match by performing the following operations
// to the first one:
// 1. Capitalize as many lower case letters as you want
// 2. Delete all lower case letters

/*
Dynamic Programming

https://www.geeksforgeeks.org/dynamic-programming/
https://en.wikipedia.org/wiki/Dynamic_programming



*/

// a = the string to modify
// b = the string to match
func abbreviationDyn(a string, b string) bool {
	// Is there a similar approach like we had with non-adj sums?
	// ie go through "a" and at each character decide whether to
	// capitalise or discard it
	// doesn't seem likely because there's not an easy way to "choose"
	// like how we had max(incl, excl) before
	return false
}

// Below doesn't work so good
func abbreviation(a string, b string) bool {
	// iterate through string b
	runes := []rune(a) // Convert a to an array of runes for index access
	var i int
	for _, v := range b {
		// look through available letters of a
		for i <= len(a) {
			// if upper(a) = b then we can use this character
			// here we assume all characters are ASCII and that
			// b is all uppercase

			// if we reach the end of a before the end of b,
			// it is not possible
			if i == len(a) {
				return false
			}

			// fmt.Printf("%c %d %c %d\n", v, v, runes[i], runes[i])
			// fmt.Printf("%c <= %c, %d <= %d, %t\n", runes[i], 'Z', runes[i], 'Z', runes[i] <= 'Z')
			// if a is uppercase and doesn't match b, we have a problem
			if runes[i] <= 'Z' && runes[i] != v {
				// fmt.Println("a is uppercase and doesn't match b")
				return false
			}

			if v == runes[i] || v == runes[i]-32 {
				i++
				break // found a match we can use
			}

			// Otherwise a is lowercase non-matching, skip it
			i++
		}
	}
	// check that there are no remaining uppercase letters in a

	// problem: my code assumes that you should convert the first lowercase letter
	// that you find if it matches b.
	// but what if you need to skip some matching lower case letters in favour
	// of some uppercase letters (since those can't be removed from "a")?

	// if we reach here then all letters of b are assigned
	return true
}
