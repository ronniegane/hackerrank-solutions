package triangles

/*
We will call a right angled triangle perfect if

* it is a primitive right angled triangle
  * this means the greatest common divisor of all three sides is 1.
* its hypotenuse is a perfect square (= n^2 for some n)

We will call a right angled triangle super-perfect if

* it is a perfect right angled triangle and
* its area is a multiple of the perfect numbers 6 and 28.
  * a "perfect number" is equal to the sum of its divisors,
	eg 6 = 1 + 2 +3 and 28 = 1 + 2 + 4 + 7 + 14

Generating pythagorean triples:
https://en.wikipedia.org/wiki/Pythagorean_triple
*/
func main() {
	// Read input from STDIN. Print output to STDOUT
	// format will be: first line = number of queries q.
	// next q lines are values of n to process.
}

// triangleCount takes in an upper bound on the largest
// side of the triangle (n)
// and returns the count of triangles up to this size
// that are perfect but not super-perfect

func triangleCount(n int) int {
	return -1
}
