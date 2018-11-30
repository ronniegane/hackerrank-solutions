package palindrome

import "testing"

type testCase struct {
	name string
	n    int32
	s    string
	want int64
}

var testCases = []testCase{
	{"0", 5, "asasd", 7},
	{"1", 7, "abcbaba", 10},
	{"2", 4, "aaaa", 10},
}

func TestSubstrCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := substrCount(tc.n, tc.s)
			if got != tc.want {
				t.Errorf("substrCount(%d, %q) = %d; expected %d", tc.n, tc.s, got, tc.want)
			}
		})
	}
}

var result int64

// Worst case is every letter is the same
func BenchmarkSubstrCount10(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrCount(10, "aaaaaaaaaa")
	}
	result = res
}

func BenchmarkSubstrCount20(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrCount(20, "aaaaaaaaaaaaaaaaaaaa")
	}
	result = res
}

func BenchmarkSubstrCount50(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrCount(50, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
	result = res
}

func BenchmarkSubstrCount100(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrCount(100, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
	result = res
}

// Behaviour of the naive approach seems to be
// double the size -> 5-7x the solve time
// since n can be as large as 1e6 this is an issue.

func TestSubstrChain(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := substrChain(tc.n, tc.s)
			if got != tc.want {
				t.Errorf("substrChain(%d, %q) = %d; expected %d", tc.n, tc.s, got, tc.want)
			}
		})
	}
}

// Worst case for new approach?
// All characters the same is the easiest case for the new approach
// Is all characters different harder??
func BenchmarkSubstrChain10(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrChain(10, "aaabaaabaa")
	}
	result = res
}
func BenchmarkSubstrChain10_allDifferent(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrChain(10, "abcdefghij")
	}
	result = res
}

func BenchmarkSubstrChain20(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrChain(20, "aaaaabaaaaaaabaaaaaa")
	}
	result = res
}

func BenchmarkSubstrChain50(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrChain(50, "aaaaaabaaaaaabbaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaa")
	}
	result = res
}

func BenchmarkSubstrChain100(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = substrChain(100, "aaaaaaaabbbaaaaaaaabaaaaaabaaaaaaaaaaaaaaaacaaaaaaaaaaaaaaabaaaaaaaaacaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
	result = res
}
