package castle

import "testing"

type testCase struct {
	name   string
	grid   []string
	startX int32
	startY int32
	goalX  int32
	goalY  int32
	ans    int32
}

var testCases = []testCase{
	{"0", []string{"...", ".X.", "..."}, 0, 0, 1, 2, 2},
	{"1", []string{".X.", ".X.", "..."}, 0, 0, 0, 2, 3},
}

func TestMinimumMoves(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minimumMoves(tc.grid, tc.startX, tc.startY, tc.goalX, tc.goalY)
			if got != tc.ans {
				t.Errorf("Got %d, expected %d", got, tc.ans)
			}
		})
	}
}
