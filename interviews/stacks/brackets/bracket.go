package brackets

// Implementing stacks/queues in Go
// Way to do it seems to be slices

func balanced(s string) bool {
	match := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack = make([]rune, 0)
	for _, ch := range s {
		// If opening bracket, add to the stack
		if _, ok := match[ch]; ok {
			stack = append(stack, ch)
		} else {
			// If we hit a closing bracket and the stack is empty, mismatched
			if len(stack) == 0 {
				return false
			}
			// If closing bracket, check it matches the last open bracket
			last := stack[len(stack)-1]
			if match[last] != ch {
				// Unmatched bracket
				return false
			}
			// pop the stack
			stack = stack[:len(stack)-1]
		}
	}

	// finally check for any remaining open brackets
	return len(stack) == 0
}
