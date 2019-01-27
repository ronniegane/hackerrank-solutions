package pairs

import "testing"

type testCase struct {
	name string
	arr  []int32
	k    int32
	ans  int32
}

var testCases = []testCase{
	{"0", []int32{1, 2, 3, 4}, 1, 3},
	{"1", []int32{1, 5, 3, 4, 2}, 2, 3},
}

func TestNaivePairs(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := naivePairs(tc.k, tc.arr)
			if got != tc.ans {
				t.Errorf("Found %d pairs, expected %d", got, tc.ans)
			}
		})
	}
}

func TestPairs(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := pairs(tc.k, tc.arr)
			if got != tc.ans {
				t.Errorf("Found %d pairs, expected %d", got, tc.ans)
			}
		})
	}
}

func TestWithSet(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := withSet(tc.k, tc.arr)
			if got != tc.ans {
				t.Errorf("Found %d pairs, expected %d", got, tc.ans)
			}
		})
	}
}

var result int32

var arr10 = []int32{1, 5, 3, 4, 2, 9, 7, 10, 8, 6}
var arr20 = []int32{1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6}
var arr40 = []int32{1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6}
var arr80 = []int32{1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6, 1, 5, 3, 4, 2, 9, 7, 10, 8, 6}

func BenchmarkNaivePairs_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = naivePairs(2, arr10)
	}
	result = res
}

func BenchmarkNaivePairs_20(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = naivePairs(2, arr20)
	}
	result = res
}

func BenchmarkNaivePairs_40(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = naivePairs(2, arr40)
	}
	result = res
}

func BenchmarkNaivePairs_80(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = naivePairs(2, arr80)
	}
	result = res
}

func BenchmarkPairs_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = pairs(2, arr10)
	}
	result = res
}

func BenchmarkPairs_20(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = pairs(2, arr20)
	}
	result = res
}

func BenchmarkPairs_40(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = pairs(2, arr40)
	}
	result = res
}
func BenchmarkPairs_80(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = pairs(2, arr80)
	}
	result = res
}

func BenchmarkSets_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = withSet(2, arr10)
	}
	result = res
}

func BenchmarkSets_20(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = withSet(2, arr20)
	}
	result = res
}

func BenchmarkSets_40(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = withSet(2, arr40)
	}
	result = res
}
func BenchmarkSets_80(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = withSet(2, arr80)
	}
	result = res
}
