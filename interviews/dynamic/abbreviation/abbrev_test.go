package abbreviation

import "testing"

type testCase struct {
	name string
	a    string
	b    string
	want bool
}

var testCases = []testCase{
	{"simple", "abcd", "ABCD", true},
	{"match", "ABCD", "ABCD", true},
	{"0", "daBcd", "ABC", true},
	{"1", "abcDE", "ABDE", true},
	{"2", "AbcDE", "AFDE", false},
	{"can replace", "AbcBcD", "ABCD", true},
	{"don't skip capitals", "AbCBcD", "ABCD", false},
	{"extra capitals", "abcdeF", "ABCDE", false},
}

func TestAbbreviation(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("%q %q", tc.a, tc.b)
			got := abbreviation(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("%q %q = %t; wanted %t", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
