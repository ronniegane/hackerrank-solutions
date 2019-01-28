package manipulation

import "testing"

type testCase struct {
	name    string
	n       int32
	queries [][]int32
	max     int64
}

var testCases = []testCase{
	{"0", 10, [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}, 10},
	{"1", 5, [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}, 200},
	{"bigarray", 100000, [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}, 10},
}

func TestManipulation(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := arrayManipulation(tc.n, tc.queries)
			if got != tc.max {
				t.Errorf("Max = %d, expected %d", got, tc.max)
			}
		})
	}
}

func TestMapBased(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := mapBased(tc.n, tc.queries)
			if got != tc.max {
				t.Errorf("Max = %d, expected %d", got, tc.max)
			}
		})
	}
}

func TestDiffBased(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := diffBased(tc.n, tc.queries)
			if got != tc.max {
				t.Errorf("Max = %d, expected %d", got, tc.max)
			}
		})
	}
}

func TestDiffMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := diffMap(tc.n, tc.queries)
			if got != tc.max {
				t.Errorf("Max = %d, expected %d", got, tc.max)
			}
		})
	}
}

var result int64

func BenchmarkManipulation(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = arrayManipulation(testCases[2].n, testCases[2].queries)
	}
	result = res
}

func BenchmarkMapBased(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = mapBased(testCases[2].n, testCases[2].queries)
	}
	result = res
}

func BenchmarkDiffBased(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = diffBased(testCases[2].n, testCases[2].queries)
	}
	result = res
}

func BenchmarkDiffMap(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = diffMap(testCases[2].n, testCases[2].queries)
	}
	result = res
}
