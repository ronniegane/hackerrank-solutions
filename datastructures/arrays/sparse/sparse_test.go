package sparse

import "testing"

type testCase struct {
	name string
	s    []string
	q    []string
	want []int32
}

var testCases = []testCase{
	{"1", []string{
		"aba",
		"baba",
		"aba",
		"xzxb",
	}, []string{
		"aba",
		"xzxb",
		"ab",
	}, []int32{
		2,
		1,
		0,
	}},
	{"2", []string{
		"def",
		"de",
		"fgh",
	}, []string{
		"de",
		"lmn",
		"fgh",
	}, []int32{1, 0, 1}},
	{"3", []string{

		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
		"na",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
	}, []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
	}, []int32{1, 3, 4, 3, 2}},
}

func TestMatch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := match(tc.s, tc.q)
			t.Log(got)
			t.Log(tc.want)
			if len(got) != len(tc.want) {
				t.Fatalf("Got %d results, wanted %d", len(got), len(tc.want))
			}
			for i, v := range tc.want {
				if got[i] != v {
					t.Errorf("Difference at index %d: got %d; wanted %d", i, got[i], v)
				}
			}
		})
	}
}

func TestMatchMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := matchMap(tc.s, tc.q)
			t.Log(got)
			t.Log(tc.want)
			if len(got) != len(tc.want) {
				t.Fatalf("Got %d results, wanted %d", len(got), len(tc.want))
			}
			for i, v := range tc.want {
				if got[i] != v {
					t.Errorf("Difference at index %d: got %d; wanted %d", i, got[i], v)
				}
			}
		})
	}
}

var result []int32

func BenchmarkMatch(b *testing.B) {
	var res []int32
	for i := 0; i < b.N; i++ {
		res = match(testCases[2].s, testCases[2].q)
	}
	result = res
}

func BenchmarkMatchMap(b *testing.B) {
	var res []int32
	for i := 0; i < b.N; i++ {
		res = matchMap(testCases[2].s, testCases[2].q)
	}
	result = res
}
