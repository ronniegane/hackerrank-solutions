package gridcrypt

import (
	"fmt"
	"testing"
)

type testCase struct {
	name       string
	cleartext  string
	ciphertext string
}

var testCases = []testCase{
	{"0", "haveaniceday", "hae and via ecy"},
	{"1", "feedthedog", "fto ehg ee dd"},
	{"2", "chillout", "clu hlt io"},
}

func TestEncryption(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := encryption(tc.cleartext)
			if encoded != tc.ciphertext {
				t.Errorf("Ciphertext mismatch")
				t.Logf("expected: %q", tc.ciphertext)
				t.Logf("received: %q", encoded)
			}
		})
	}
}

func TestQuickEncryption(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := quickencryption(tc.cleartext)
			if encoded != tc.ciphertext {
				t.Errorf("Ciphertext mismatch")
				t.Logf("expected: %q", tc.ciphertext)
				t.Logf("received: %q", encoded)
			}
		})
	}
}

var benchCases = []string{
	"haveaniced",
	"haveanicedhaveaniced",
	"haveanicedhaveanicedhaveanicedhaveaniced",
	"haveanicedhaveanicedhaveanicedhaveanicedhaveanicedhaveanicedhaveanicedhaveaniced",
}
var result string

func BenchmarkEncryption(b *testing.B) {
	for _, bench := range benchCases {
		b.Run(fmt.Sprintf("%d", len(bench)), func(b *testing.B) {
			var res string
			for i := 0; i < b.N; i++ {
				res = encryption(bench)
			}
			result = res
		})
	}
}
func BenchmarkQuickEncryption(b *testing.B) {
	for _, bench := range benchCases {
		b.Run(fmt.Sprintf("%d", len(bench)), func(b *testing.B) {
			var res string
			for i := 0; i < b.N; i++ {
				res = quickencryption(bench)
			}
			result = res
		})
	}
}
