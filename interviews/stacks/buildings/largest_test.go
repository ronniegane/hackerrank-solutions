package buildings

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

type testCase struct {
	name string
	arr  []int32
	ans  int64
}

var testCases = []testCase{
	{"0", []int32{3, 2, 3}, 6},
	{"5", []int32{1, 2, 3, 4, 5}, 9},
	{"25", []int32{10, 8, 6, 4, 2, 10, 8, 6, 4, 2, 10, 8, 6, 4, 2, 10, 8, 6, 4, 2, 10, 8, 6, 4, 2}, 50},
	{"50", []int32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 50},
	{"100", []int32{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 100},
	{"end minimum", []int32{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116}, 26152},
}

func TestLargestRectangle(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := largestRectangle(tc.arr)
			if got != tc.ans {
				t.Errorf("%v largest rect = %d; wanted %d", tc.arr, got, tc.ans)
			}
		})
	}
}

type longtest struct {
	filename string
	answer   int64
}

func TestLargeInputs(t *testing.T) {
	inputs := []longtest{
		{"input06.txt", 10558350},
		{"input08.txt", 12984467},
	}

	for _, tc := range inputs {
		// testing with a 10^5 long array
		file, err := os.Open(tc.filename)
		check(err)
		defer file.Close()
		t.Run(tc.filename, testLargeInput(file, tc.answer))
	}
}

func testLargeInput(file io.Reader, answer int64) func(t *testing.T) {

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Sizes split by spaces
	// first line is array length
	scanner.Scan()
	numBuildings, err := strconv.Atoi(scanner.Text())
	fmt.Println("number of buildings:", numBuildings)
	check(err)

	buildingArr := make([]int32, numBuildings)

	// Read the building sizes into the array

	var height int
	for i := range buildingArr {
		scanner.Scan()
		height, err = strconv.Atoi(scanner.Text())
		check(err)

		buildingArr[i] = int32(height)
	}

	fmt.Println("Array built")
	fmt.Printf("%d ... %d\n", buildingArr[0], buildingArr[numBuildings-1])
	// Test the array

	return func(t *testing.T) {
		got := largestRectangle(buildingArr)

		if got != answer {
			t.Errorf("Max area = %d; expected %d", got, answer)
		}
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var result int64

func BenchmarkLargestRectangle_5(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = largestRectangle(testCases[1].arr)
	}
	result = res
}

func BenchmarkLargestRectangle_25(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = largestRectangle(testCases[2].arr)
	}
	result = res
}

func BenchmarkLargestRectangle_50(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = largestRectangle(testCases[3].arr)
	}
	result = res
}

func BenchmarkLargestRectangle_100(b *testing.B) {
	var res int64
	for i := 0; i < b.N; i++ {
		res = largestRectangle(testCases[4].arr)
	}
	result = res
}
