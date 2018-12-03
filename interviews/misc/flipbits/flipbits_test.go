package flipbits

import "testing"

type testCase struct {
	name string
	in   int64
	out  int64
}

var testCases = []testCase{
	{"0a", 2147483647, 2147483648},
	{"0b", 1, 4294967294},
	{"0c", 0, 4294967295},
	{"1a", 4, 4294967291},
	{"1b", 123456, 4294843839},
}

func TestFlip(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := flip(tc.in)
			if got != tc.out {
				t.Errorf("flip(%d) = %d; expected %d", tc.in, got, tc.out)
			}
		})
	}
}

func TestActualFlip(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := actualFlip(tc.in)
			if got != tc.out {
				t.Errorf("actualFlip(%d) = %d; expected %d", tc.in, got, tc.out)
			}
		})
	}
}

var result int64

func BenchmarkFlip(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = flip(2147483647)
	}
	result = res
}

func BenchmarkActualFlip(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = actualFlip(1234748367)
	}
	result = res
}
