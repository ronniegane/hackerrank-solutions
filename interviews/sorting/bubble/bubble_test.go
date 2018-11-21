package bubble

import "testing"

type testCase struct {
	name  string
	arr   []int32
	swaps int
	first int
	last  int
}

var testCases = []testCase{
	{"0", []int32{1, 2, 3}, 0, 1, 3},
	{"1", []int32{3, 2, 1}, 3, 1, 3},
	{"2", []int32{4, 2, 3, 1}, 5, 1, 4},
}

func TestSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, f, l := Sort(tc.arr)
			if s != tc.swaps || f != tc.first || l != tc.last {
				t.Errorf("Array %v sorted in %d[%d] swaps, first = %d[%d], last=%d[%d]",
					tc.arr, s, tc.swaps, f, tc.first, l, tc.last)
			}
		})
	}
}
