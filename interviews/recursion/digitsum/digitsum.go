package digitsum

import (
	"strconv"
	"strings"
)

/*

We define super digit of an integer

using the following rules:

Given an integer x, we need to find the super digit of the integer.

    If x has only digit, then its super digit is x
.
Otherwise, the super digit of x
is equal to the super digit of the sum of the digits of x.

eg

	super_digit(9875)   9+8+7+5 = 29
	super_digit(29) 	2 + 9 = 11
	super_digit(11)		1 + 1 = 2
	super_digit(2)		= 2
*/

// First attempt works with n as a number,
// performing modulo division on it
// would it be better to work with an array of digits
// to avoid the modulo division?

func superDigit(n string, k int32) int32 {
	// turn n into an int
	// n could be much larger than int32
	// in fact, constraints say n is up to 10^100000
	// so no way can we turn it into any numeric value.
	// will have to sum it from the parts of the string.
	// will this time out though?
	// actually, probably not - it's not 10^100000 loop cycles,
	// it's only 100000 - the number of digits
	var sum int64
	var curr int
	strSlice := strings.Split(n, "")
	for _, s := range strSlice {
		curr, _ = strconv.Atoi(s)
		sum += int64(curr)
	}

	// Do the first level of digit summing
	// note that sum(n*k) = sum(n) * k
	// Is there potential for overflow here?
	// n capped at 10^5 digits, so sum is max 10^5
	// k is capped at 10^5, so sum * k could be ~10^10
	// int32 max is 2.15 * 10^10 so we're _probably_ okay
	// ...nope, case 7 is overflowing (sum = 44917 and k = 100000)
	// maybe work with int64 and then convert afterwards
	// annoying because only the very first mulitplication should need this
	return int32(super(sum * int64(k)))
}

func super(n int64) int64 {
	if n < 10 {
		return n
	}
	// Sum digits of n and recurse
	return super(sumDigits(n))
}

func sumDigits(n int64) int64 {
	var sum int64
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

// Apparently the sum of the digits of any number N to base 10
// is equivalent to it's remainder mod 9.
// But for that to work we would have to do the first sum by string anyway
// as 10^100000 is too large to fit into a numerical type
func justModNine(n string, k int32) int32 {
	// would be
	// rem := n * k % 9
	// if rem == 0 {
	// 	return 9
	// }
	// return rem
	return 0
}

// slice-of-ints version
// func superSlice(n string, k int32) int32 {
// 	// turn n into an array of ints
// 	strSlice := strings.Split(n, "")
// 	numSlice := make([]int, len(strSlice))
// 	for i, s := range strSlice {
// 		numSlice[i], _ = strconv.Atoi(s)
// 	}

// 	// Do the first level of digit summing
// 	// note that sum(n*k) = sum(n) * k
// 	return int32(sumSlice(numSlice)* int(k))
// }

// func sliceRecurse(digits []int) int {
// if len(digits) == 1 {
// 	// base case
// 	return digits[0]
// }
// return sliceRecurse()
// }

// Difficulty here: summing a []int -> int but then
// splitting that int back into digits ([]int)

// func sumSlice(digits []int) []int {
// 	var sum int
// for _, v := range digits {
// 	sum += v
// }
// // need to turn this back into a slice of digits
// return sum
// }
