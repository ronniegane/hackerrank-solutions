package sherlock

import "testing"

type testCase struct {
	name  string
	input string
	want  string
}

var testCases = []testCase{
	{"0", "aabbcd", "NO"},
	{"1", "aabbccddeefghi", "NO"},
	{"2", "abcdefghhgfedecba", "YES"},
	{"3", "aabbcc", "YES"},
	{"4", "aabbccd", "YES"},
	{"5", "abbccdd", "YES"},
	{"6", "abbcacdda", "YES"},
	{"7", "aabbb", "YES"},
}

func TestIsValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValid(tc.input)
			if got != tc.want {
				t.Errorf("isValid(%q) = %s; expected %s", tc.input, got, tc.want)
			}
		})
	}
}
