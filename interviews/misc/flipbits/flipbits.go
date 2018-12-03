package flipbits

// Working with 32 bits
func flip(n int64) int64 {
	return 4294967295 - n
}

// Flipping bits is the same as subtraction?
// makes sense
// 0010 + 1101 = 1111 = (1000 - 1) or 2^3-1 in base 10
// so subtract n from the max integer possible in 32 bits
// to find the flipped answer

func actualFlip(n int64) int64 {
	// Go doesn't have native support for binary integers yet
	// XOR with a 111111 mask
	return n ^ 4294967295
}
