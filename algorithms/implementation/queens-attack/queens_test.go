package attack

import "testing"

type testCase struct {
	name      string
	n         int32
	k         int32
	Qrow      int32
	Qcol      int32
	obstacles [][]int32
	answer    int32
}

var testCases = []testCase{
	{"sample 0", 4, 0, 4, 4, [][]int32{}, 9},
	{"sample 1", 5, 3, 4, 3, [][]int32{{5, 5}, {4, 2}, {2, 3}}, 10},
	{"sample 2", 1, 0, 1, 1, [][]int32{}, 0},
}

func TestQueensAttack(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := queensAttack(tc.n, tc.k, tc.Qrow, tc.Qcol, tc.obstacles)
			if got != tc.answer {
				t.Errorf("Found %d attackable squares; expected %d", got, tc.answer)
			}
		})
	}
}
