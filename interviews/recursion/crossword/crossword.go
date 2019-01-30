package crossword

/*
https://www.hackerrank.com/challenges/crossword-puzzle/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking

crossword is an array of 10 strings of length 10 forming a grid
if the character in crossword row is "-" that is a blank space

words is a semicolon-delimited list of words
*/

func crosswordPuzzle(crossword []string, words string) []string {
	/*
		High-level idea:
		find the intersections of words in the grid,
		to form constraints that eg word A char 3 = word B char 7
		then check the provided word list for candidates

		Other option: backtracking search
		- find the position of words in the grid (start,stop coords)
		- iterate through potential positions, try word A (if it fits),
		  then next position try word B, check for conflicts
		- if you find conflicts, then back up to the last position and try
		  a different option


	*/
	var output []string
	// Parse crossword grid

	// Parse words from semicolon-delimited string

	// solve

	return output

}
