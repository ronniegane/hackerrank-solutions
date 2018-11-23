package abs

import "testing"

type testCase struct {
	name string
	arr  []int32
	min  int32
}

var testCases = []testCase{
	{"0", []int32{3, -7, 0}, 3},
	{"1", []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}, 1},
	{"2", []int32{1, -3, 71, 68, 17}, 3},
}

func TestMinAbsDiff(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minAbsDiff(tc.arr)
			if got != tc.min {
				t.Errorf("Minimum diff found = %d; expected %d", got, tc.min)
			}
		})
	}
}
