package luck

import "testing"

type testCase struct {
	name     string
	contests [][]int32
	k        int32
	want     int32
}

var testCases = []testCase{
	{"0", [][]int32{
		{5, 1},
		{2, 1},
		{1, 1},
		{8, 1},
		{10, 0},
		{5, 0},
	}, 3, 29},
	{"1", [][]int32{
		{13, 1},
		{10, 1},
		{9, 1},
		{8, 1},
		{13, 1},
		{12, 1},
		{18, 1},
		{13, 1},
	}, 5, 42},
}

func TestSortLuck(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortLuck(tc.k, tc.contests)
			if got != tc.want {
				t.Errorf("Got %d; expected %d", got, tc.want)
			}
		})
	}
}
