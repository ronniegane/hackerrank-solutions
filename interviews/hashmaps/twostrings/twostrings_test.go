package twostrings

import "testing"

type testCase struct {
	name string
	a    string
	b    string
	want bool
}

func TestCommon(t *testing.T) {
	testCases := []testCase{
		{"1", "hello", "world", true},
		{"2", "hi", "world", false},
		{"3", "wouldyoulikefries", "abcabcabcabcabcabc", false},
		{"4", "hackerrankcommunity", "cdecdecdecde", true},
		{"5", "jackandjill", "wentupthehill", true},
		{"5", "writetoyourparents", "fghmqzldbc", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Common(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Common(%q, %q) = %t; wanted %t", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
