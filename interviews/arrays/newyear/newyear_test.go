package newyear

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	name  string
	input []int32
	want  int
}

var testCases = []testCase{
	{"simple", []int32{2, 1, 5, 3, 4}, 3},
	{"impossible", []int32{2, 5, 1, 3, 4}, -1},
	{"longer impossible", []int32{5, 1, 2, 3, 7, 8, 6, 4}, -1},
	{"longer", []int32{1, 2, 5, 3, 7, 8, 6, 4}, 7},
	{"longer2", []int32{1, 2, 5, 3, 4, 7, 8, 6}, 4},
}

func TestMinBribes(t *testing.T) {
	// Compare test case longer
	// {1, 2, 3, 4, 5, 6, 7, 8}
	// {1, 2, 5, 3, 7, 8, 6, 4}
	// Naive approach guesses 6 swaps
	// But that would leave us with {4, 6} at the end
	// Need to account for other things out of order
	// what about counting disorder?
	// for each value V look through to find the index of (V-1)
	// add that value to the total

	// is this a reverse/"undoing" bubble sort?
	// https://www.geeksforgeeks.org/number-swaps-sort-adjacent-swapping-allowed/
	// https://www.geeksforgeeks.org/counting-inversions/
	// "number of inversion" count the number of elements to the right of N that are smaller than N
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minBribes(tc.input)
			if got != tc.want {
				t.Errorf("Minimum bribes for %v = %d, should be %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestNaive(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := minBribesNaive(tc.input)
			if got != tc.want {
				t.Errorf("Minimum bribes for %v = %d, should be %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestRevBribes(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := revBribes(tc.input)
			if got != tc.want {
				t.Errorf("Minimum bribes for %v = %d, should be %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestMinBribes_long(t *testing.T) {
	// Long cases to test efficiency
	// Original naive approach
	// does (n-1) + (n-2) + (n-3)... operations
	// seems like O(n^2)?

	// Discussion suggests that using merge sort to count inversions
	// Will be more time efficient - O(n log(n))

	// test input
	// first line: number of cases
	// each pair of lines after: number of elements, then space-separated elements

	file, err := os.Open("long_test_input")
	check(err)
	defer file.Close()

	answerFile, err := os.Open("long_test_output")
	check(err)
	defer answerFile.Close()

	reader := bufio.NewReaderSize(file, 1024*1024)

	str, _, err := reader.ReadLine()
	check(err)
	numCases, err := strconv.Atoi(string(str))
	check(err)

	ansReader := bufio.NewReader(answerFile)

	for i := 1; i <= numCases; i++ {
		_, _, _ = reader.ReadLine() // throw away number of elements line
		str, _, err := reader.ReadLine()
		check(err)

		// Parse the string of space-separated ints
		var q []int32
		for _, f := range strings.Fields(string(str)) {
			n, err := strconv.Atoi(f)
			check(err)
			q = append(q, int32(n))
		}
		// use the queue
		str, _, err = ansReader.ReadLine()
		check(err)
		want, err := strconv.Atoi(string(str))
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := revBribes(q)
			if got != want {
				t.Errorf("Calculated moves: %d; expected %d", got, want)
			}
		})
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var result int

func BenchmarkMinBribes_10(b *testing.B) {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4, 9, 10}
	var res int
	for i := 0; i < b.N; i++ {
		res = minBribes(q)
	}
	result = res
}

func BenchmarkMinBribes_20(b *testing.B) {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4, 9, 10,
		11, 12, 13, 14, 15, 16, 19, 18, 17, 20}
	var res int
	for i := 0; i < b.N; i++ {
		res = minBribes(q)
	}
	result = res
}

func BenchmarkMinBribes_40(b *testing.B) {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4, 9, 10,
		11, 12, 13, 14, 15, 16, 19, 18, 17, 20,
		21, 22, 23, 24, 25, 26, 29, 28, 27, 30,
		31, 32, 33, 34, 35, 36, 39, 38, 37, 40}
	var res int
	for i := 0; i < b.N; i++ {
		res = minBribes(q)
	}
	result = res
}
