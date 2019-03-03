package leaderboard

import (
	"reflect"
	"testing"
)

type testCase struct {
	name   string
	scores []int32
	alice  []int32
	ans    []int32
}

var testCases = []testCase{
	{"1", []int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120}, []int32{6, 4, 2, 1}},
	{"2", []int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}, []int32{6, 5, 4, 2, 1}},
}

func TestClimbing(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := climbingLeaderboard(tc.scores, tc.alice)
			// test list equality
			if !reflect.DeepEqual(got, tc.ans) {
				t.Errorf("Scores were %v\n\t\t\t   expected %v", got, tc.ans)
			}
		})
	}
}
