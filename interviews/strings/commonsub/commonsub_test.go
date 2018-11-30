package commonsub

import "testing"

type testCase struct {
	name string
	s1   string
	s2   string
	want int32
}

var testCases = []testCase{
	{"0", "HARRY", "SALLY", 2},
	{"1", "AA", "BB", 0},
	{"2", "SHINCHAN", "NOHARAA", 3},
	{"3", "ABCDEF", "FBDAMN", 2},
	{"4", "BANANA", "ATANA", 4},
	{"5", "BANANA", "BANANA", 6},
}

func TestLongestSub(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestSub(tc.s1, tc.s2)
			if got != tc.want {
				t.Errorf("Longest common subsequence of %q and %q is %d chars long; expected %d", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

var result int32

func BenchmarkLongestSub_5(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = longestSub("SHINC", "NOHAR")
	}
	result = res
}

func BenchmarkLongestSub_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = longestSub("SHINCABCDE", "NOHARBCDEF")
	}
	result = res
}

/*
recursive approach is (unsurprisingly) exponential
BenchmarkLongestSub_5-4          1000000              1715 ns/op
BenchmarkLongestSub_10-4            5000            296048 ns/op
*/
// Even with 10-length strings the recursive approach is 29604 ns/op
// Needs to work with up to 5000-char strings

// ==== 2D array approach ====

func TestArraySub(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := arraySub(tc.s1, tc.s2)
			if got != tc.want {
				t.Errorf("Longest common subsequence of %q and %q is %d chars long; expected %d", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func BenchmarkArraySub_5(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = arraySub("SHINC", "NOHAR")
	}
	result = res
}

func BenchmarkArraySub_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = arraySub("SHINCABCDE", "NOHARBCDEF")
	}
	result = res
}
