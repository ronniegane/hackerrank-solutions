package staircase

import (
	"fmt"
	"testing"
)

type testCase struct {
	in  int32
	out int32
}

var testCases = []testCase{
	{1, 1},
	{2, 2},
	{3, 4},
	{4, 7},
	{5, 13},
	{7, 44},
}

func TestNaiveStepPerms(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			got := stepPerms(tc.in)
			if tc.out != got {
				t.Errorf("perms(%d) = %d; expected %d", tc.in, got, tc.out)
			}
		})
	}
}

func TestMemo(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			got := stepMemo(tc.in)
			if tc.out != got {
				t.Errorf("memo(%d) = %d; expected %d", tc.in, got, tc.out)
			}
		})
	}
}

var result int32

func BenchmarkNaive_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepPerms(10)
	}
	result = res
}

func BenchmarkNaive_20(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepPerms(20)
	}
	result = res
}

func BenchmarkNaive_36(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepPerms(36)
	}
	result = res
}

func BenchmarkMemo_10(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepMemo(10)
	}
	result = res
}

func BenchmarkMemo_20(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepMemo(20)
	}
	result = res
}

func BenchmarkMemo_36(b *testing.B) {
	var res int32
	for i := 0; i < b.N; i++ {
		res = stepMemo(36)
	}
	result = res
}
