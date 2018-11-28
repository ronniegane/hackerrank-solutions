package brackets

import "testing"

type testCase struct {
	name  string
	input string
	want  bool
}

var testCases = []testCase{
	{"0", "{[()]}", true},
	{"1", "{[(])}", false},
	{"2", "{{[[(())]]}}", true},
	{"3", "}{}{}", false},
	{"4", "[](){}", true},
}

func TestBalanced(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := balanced(tc.input)
			if got != tc.want {
				t.Errorf("%q returned %t; expected %t", tc.input, got, tc.want)
			}
		})
	}
}
