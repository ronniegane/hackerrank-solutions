package minswaps

import "testing"

type testCase struct {
	name  string
	arr   []int32
	swaps int32
}

var testCases = []testCase{
	{"0", []int32{4, 3, 1, 2}, 3},
	{"1", []int32{2, 3, 4, 1, 5}, 3},
	{"2", []int32{1, 3, 5, 2, 4, 6, 7}, 3},
}

func TestMinimumSwaps(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minimumSwaps(tc.arr)
			if got != tc.swaps {
				t.Errorf("Got %d swaps; expected %d", got, tc.swaps)
			}
		})
	}
}

func TestWithoutTrackStruct(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := withoutTrackStruct(tc.arr)
			if got != tc.swaps {
				t.Errorf("Got %d swaps; expected %d", got, tc.swaps)
			}
		})
	}
}

var result int32

var bigArray = []int32{1, 3, 5, 2, 4, 6, 7, 9, 10, 8, 16, 15, 14, 12, 13, 11, 19, 17, 18, 20}

func BenchmarkMinimumSwaps(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = minimumSwaps(bigArray)
	}
	result = res
}

func BenchmarkWithoutTrackStruct(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = withoutTrackStruct(bigArray)
	}
	result = res
}
