package fraud

import "testing"

type testCase struct {
	name     string
	expend   []int32
	lookback int32
	want     int32
}

var testCases = []testCase{
	{"0", []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5, 2},
	{"1", []int32{1, 2, 3, 4, 4}, 4, 0},
	{"2", []int32{10, 20, 30, 40, 50}, 3, 1},
}

func TestBruteMedian(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BruteMedian(tc.expend, tc.lookback)
			if got != tc.want {
				t.Errorf("%v led to %d notifications; expected %d", tc.expend, got, tc.want)
			}
		})
	}
}

func TestCountSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountSort(tc.expend, tc.lookback)
			if got != tc.want {
				t.Errorf("%v led to %d notifications; expected %d", tc.expend, got, tc.want)
			}
		})
	}
}
