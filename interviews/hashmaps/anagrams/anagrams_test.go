package anagrams

import "testing"

type testCase struct {
	s string
	n int32
}

func TestSherlock(t *testing.T) {
	testCases := []testCase{
		{"abba", 4},
		{"abcd", 0},
		{"ifailuhkqq", 3},
		{"kkkk", 10},
		{"cdcd", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := Sherlock(tc.s)
			if got != tc.n {
				t.Errorf("Sherlock(%q) = %d; wanted %d", tc.s, got, tc.n)
			}
		})
	}
}

func TestSignature(t *testing.T) {
	testCases := []testCase{
		{"abba", 4},
		{"abcd", 0},
		{"ifailuhkqq", 3},
		{"kkkk", 10},
		{"cdcd", 5},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := Signature(tc.s)
			if got != tc.n {
				t.Errorf("Signature(%q) = %d; wanted %d", tc.s, got, tc.n)
			}
		})
	}
}

var result int32

func BenchmarkSherlock(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = Sherlock("ifailuhkqq")
	}
	result = res
}

func BenchmarkSignature(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = Signature("ifailuhkqq")
	}
	result = res
}
