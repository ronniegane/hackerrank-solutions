package castle

/*
You are given a square grid with some cells open (.)
and some blocked (X). Your playing piece can move along
any row or column until it reaches the edge of the grid
or a blocked cell. Given a grid, a start and an end position,
determine the number of moves it will take to get to
the end position.
*/

// Complete the minimumMoves function below.
func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	// Implementing a breadth-first-search with a queue
	var steps, currX, currY int32
	var queue = [][]int32{{startX, startY}}
	for len(queue) > 0 {
		// pop from the queue
		currX = queue[0][0]
		currY = queue[0][1]
		queue = queue[1:]
		// try a step

		// move left if possible to the furthest

		// move right
		// move up
		// move down
		_, _ = currX, currY
	}

	return steps
}
