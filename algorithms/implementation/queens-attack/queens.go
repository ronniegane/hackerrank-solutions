package attack

import "fmt"

/*
Return the number of squares that the queen can attack
(note this is the number of squares she can move to excluding
those with obstacles)

- n: an integer, the number of rows and columns in the board
- k: an integer, the number of obstacles on the board
- r_q: integer, the row number of the queen's position
- c_q: integer, the column number of the queen's position
- obstacles: a two dimensional (2 x k) array of integers where each element
	 is an array of integers, the row and column of an obstacle
	 (in that order)

n up to 10^5
k up to 10^5

	// first idea: make a map of obstacles?

	// second idea: build the actual n x n array of the chessboard
	// with obstacles as elements
	// 2x2 array probably more space, and very sparse (up to 10^10 squares)
	// map will probably be more efficient
*/

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	// make a map of obstacles to look up
	var obsMap = map[string]bool{}
	for _, o := range obstacles {
		// don't care if there are duplicates
		// store with key x_y
		// is a string key less or more efficient than an []int key?
		obsMap[fmt.Sprintf("%d_%d", o[0], o[1])] = true
	}

	var count int32
	// move out in each of the directions from the queen until
	// we hit an obstacle

	// North
	for row := r_q + 1; row <= n; row++ {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, c_q)]; ok {
			break
		}
		count++
	}
	// South
	for row := r_q - 1; row > 0; row-- {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, c_q)]; ok {
			break
		}
		count++
	}
	// East
	for col := c_q + 1; col <= n; col++ {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", r_q, col)]; ok {
			break
		}
		count++
	}
	// West
	for col := c_q - 1; col > 0; col-- {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", r_q, col)]; ok {
			break
		}
		count++
	}

	// North-East
	for row, col := r_q+1, c_q+1; col <= n && row <= n; row, col = row+1, col+1 {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, col)]; ok {
			break
		}
		count++
	}

	// South-East
	for row, col := r_q-1, c_q+1; col <= n && row > 0; row, col = row-1, col+1 {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, col)]; ok {
			break
		}
		count++
	}

	// South-West
	for row, col := r_q-1, c_q-1; col > 0 && row > 0; row, col = row-1, col-1 {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, col)]; ok {
			break
		}
		count++
	}

	// North-West
	for row, col := r_q+1, c_q-1; col > 0 && row <= n; row, col = row+1, col-1 {
		if _, ok := obsMap[fmt.Sprintf("%d_%d", row, col)]; ok {
			break
		}
		count++
	}

	return count
}
