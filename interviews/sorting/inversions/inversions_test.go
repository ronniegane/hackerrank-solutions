package inversions

import "testing"

type testCase struct {
	name string
	arr  []int32
	want int64
}

var testCases = []testCase{
	// {"00a", []int32{1, 1, 1, 2, 2}, 0},
	// {"00b", []int32{2, 1, 3, 1, 2}, 4},
	// {"14a", []int32{1, 5, 3, 7}, 1},
	{"14b", []int32{7, 5, 3, 1}, 6},
	{"15a", []int32{1, 3, 5, 7}, 0},
	{"15b", []int32{3, 2, 1}, 3},
}

var benchmarks = []testCase{
	{"10", []int32{6, 7, 5, 3, 1, 6, 7, 5, 3, 1}, -1},
	{"50", []int32{6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1}, -1},
	{"100", []int32{6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1}, -1},
	{"500", []int32{6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1}, -1},
	{"1000", []int32{6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1, 6, 7, 5, 3, 1}, -1},
}

func TestNaiveCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := naiveCount(tc.arr)
			if got != tc.want {
				t.Errorf("Wanted %d ; got %d", tc.want, got)
			}
		})
	}
}

func TestLoopMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := naiveCount(tc.arr)
			if got != tc.want {
				t.Errorf("Wanted %d ; got %d", tc.want, got)
			}
		})
	}
}

func TestMergeCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeCount(tc.arr)
			if got != tc.want {
				t.Errorf("Wanted %d ; got %d", tc.want, got)
			}
		})
	}
}

var result int64

func BenchmarkNaiveCount(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			var res int64
			for i := 0; i < b.N; i++ {
				res = naiveCount(bm.arr)
			}
			result = res
		})
	}
}

func BenchmarkLoopMap(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			var res int64
			for i := 0; i < b.N; i++ {
				res = loopMap(bm.arr)
			}
			result = res
		})
	}
}

func BenchmarkMergeCount(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			var res int64
			for i := 0; i < b.N; i++ {
				res = mergeCount(bm.arr)
			}
			result = res
		})
	}
}
