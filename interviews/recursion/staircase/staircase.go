package staircase

/*
Davis has a number of staircases in his house and he
likes to climb each staircase 1,2, or 3 steps at a time.

How many ways are there to reach the top of a staircase of n
steps?

This is expected to be a recursive algorithm, so presumably
there is a pattern that depends on previous values of n

first few results
n   p
1	1
2	2
3	4
4	7
5	13
6	?
7	44

recursive thinking can go like this:
- first step is either 1,2,or 3
- then we can solve for (n-1), (n-2), (n-3) remaining steps
  and add the results


  The branching factor is 3 which makes this very inefficient,
  should try memoization of lower results.

n can be up to 36 which would imply ~3^36 recursive calls
*/

// also for some reason we are asked to return the number
// modulo (10^9 + 7)

var memo = map[int32]int32{
	0: 0,
	1: 1,
	2: 2,
	3: 4,
}

func stepPerms(n int32) int32 {
	// base cases
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		return stepPerms(n-1) + stepPerms(n-2) + stepPerms(n-3)
	}
}

// Memoization improves speed but stack size may still be an issue

func stepMemo(n int32) int32 {
	if v, ok := memo[n]; ok {
		return v
	}
	val := stepMemo(n-1) + stepMemo(n-2) + stepMemo(n-3)
	memo[n] = val
	return val
}
