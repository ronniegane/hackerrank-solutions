package toys

import "testing"

type testCase struct {
	name   string
	prices []int32
	budget int32
	want   int32
}

var testCases = []testCase{
	{"0", []int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4},
	{"1", []int32{1, 2, 3, 4}, 7, 3},
	{"2", []int32{3, 7, 2, 9, 4}, 15, 3},
}

func TestMax(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := max(tc.prices, tc.budget)
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}
