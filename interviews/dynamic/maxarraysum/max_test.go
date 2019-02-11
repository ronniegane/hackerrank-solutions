package maxarraysum

import "testing"

type testCase struct {
	name   string
	arr    []int32
	maxSum int32
}

var testCases = []testCase{
	{"0", []int32{3, 7, 4, 6, 5}, 13},
	{"1", []int32{2, 1, 5, 8, 4}, 11},
	{"2", []int32{3, 5, -7, 8, 10}, 15},
}

func TestMaxSubsetSum(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Testing array:", tc.arr)
			got := maxSubsetSum(tc.arr)
			if got != tc.maxSum {
				t.Errorf("Got max sum = %d; wanted %d", got, tc.maxSum)
			}
		})
	}
}
