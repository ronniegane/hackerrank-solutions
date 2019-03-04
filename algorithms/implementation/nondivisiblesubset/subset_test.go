package subset

import "testing"

type testCase struct {
	name string
	k    int32
	set  []int32
	ans  int32
}

var testCases = []testCase{
	{"0", 4, []int32{19, 10, 12, 10, 24, 25, 22}, 3},
	{"1", 3, []int32{1, 7, 2, 4}, 3},
	{"2", 5, []int32{6, 7, 8, 9, 10, 11, 12}, 5},
}

func TestNonDivisibleSubset(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := nonDivisibleSubset(tc.k, tc.set)
			if got != tc.ans {
				t.Errorf("Found size %d set, expected %d", got, tc.ans)
			}
		})
	}
}
