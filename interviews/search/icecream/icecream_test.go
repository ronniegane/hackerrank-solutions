package icecream

import "testing"

type testCase struct {
	name     string
	cost     []int32
	money    int32
	solution [2]int
}

var testCases = []testCase{
	{"0", []int32{1, 4, 5, 3, 2}, 4, [2]int{1, 4}},
	{"1", []int32{2, 2, 4, 3}, 4, [2]int{1, 2}},
}

func TestFlavorIndices(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := flavorIndices(tc.cost, tc.money)
			if got != tc.solution {
				t.Errorf("Got answer %v, wanted %v", got, tc.solution)
			}
		})
	}
}
