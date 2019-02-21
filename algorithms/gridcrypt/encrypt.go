package gridcrypt

// https://www.hackerrank.com/challenges/encryption/

import (
	"math"
)

func encryption(s string) string {
	l := len(s)
	rootL := math.Sqrt(float64(l))
	row := int(math.Floor(rootL)) // number of rows (column length)
	col := int(math.Ceil(rootL))  // number of columns (row length)
	// if this isn't large enough, increase the number of rows by 1
	if row*col < l {
		row++
	}

	grid := make([][]rune, row)
	rowidx := -1
	colidx := 0
	// Create the grid
	for i, v := range s {
		colidx = i % col
		if colidx == 0 {
			// Make a new row
			rowidx++
			grid[rowidx] = make([]rune, col)
		}
		grid[rowidx][colidx] = v
	}

	var ret []rune
	// Read the grid
	for j := 0; j < col; j++ {
		if j != 0 {
			// Move to the next column and add a space
			ret = append(ret, ' ')
		}
		for i := 0; i < row; i++ {
			// ignore null characters
			if grid[i][j] != 0 {
				ret = append(ret, grid[i][j])
			}
		}

	}

	return string(ret)
}

// Alternatively do it without creating a grid
// using modular math
func quickencryption(s string) string {
	l := len(s)
	rootL := math.Sqrt(float64(l))
	row := int(math.Floor(rootL)) // number of rows (column length)
	col := int(math.Ceil(rootL))  // number of columns (row length)
	// if this isn't large enough, increase the number of rows by 1
	if row*col < l {
		row++
	}
	// Assuming ASCII character set
	var ret []byte
	for j := 0; j < col; j++ {
		if j != 0 {
			ret = append(ret, ' ')
		}
		for i := 0; i < row; i++ {
			// index of character we want
			idx := i*col + j
			if idx < l {
				ret = append(ret, s[idx])
			}
		}
	}
	return string(ret)
}
