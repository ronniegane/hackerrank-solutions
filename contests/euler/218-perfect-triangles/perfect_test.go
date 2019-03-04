package triangles

import "testing"

type testCase struct {
	name string
	n    int
	ans  int
}

var testCases = []testCase{
	{"0", 25, 0},
}

func TestTriangleCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := triangleCount(tc.n)
			if got != tc.ans {
				t.Errorf("triangleCount(%d) = %d; expected %d", tc.n, got, tc.ans)
			}
		})
	}
}
