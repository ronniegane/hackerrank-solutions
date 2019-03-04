package subset

/*
Given a set of distinct integers, print the size of a maximal
subset of S where the sum of any numbers in is not evenly
divisible by k.
For example, the array S = [19,10,12,10,24,24,22] and k = 4.
One of the arrays that can be created is [10,12,25].
Another is [19,22,24].
After testing all permutations, the maximum length
solution array has 3 elements.

constraints:
n (length of set) <= 10^5
1 <= k <= 100
1 <= S[i] <= 10^9

far too many elements to generate all combinations, surely?
very many possible sets
Just looking at sums of pairs:
there are 10^5 * (10^5 - 1) / 2 combinations (10^5 choose 2)
which is ~ 5 * 10^9

Could we do this like a dynamic programming problem?
iterate along set, keep track of largest valid set containing i
and largest valid set not containing i ...

This approach may not be valid because maxset(0, i+1) is not necessarily
an extension of maxSet(0,i).

Explanation from discussion:
For any number K, the sum of 2 values (A & B) is evenly divisible by K
if the sum of the remainders of A/K + B/K is K.
(There is also a special case where both A & B are evenly
	divisible, giving a sum of 0.)

So we can go through and find out what the remainder (x % k) is
for each element, and store how many numbers have that remainder.
Then we know that for all potential remainder pairs (r1,r2) where
r1 + r2 = k, we choose the largest one.
Then we can also add at most _one_ value with zero remainder.
*/

// Complete the nonDivisibleSubset function below.
func nonDivisibleSubset(k int32, S []int32) int32 {
	var count int32
	// make "map" of remainders
	var remainders = make([]int32, k)

	// store counts of remainders
	for _, v := range S {
		rem := v % k
		remainders[rem]++
	}

	// iterate through remainders and choose the larger
	// of potential pairs
	// need to consider cases when k is odd or even
	// if k is even, then we can only choose at most one
	// element with remainder k/2

	for i := 1; i < int(k)-i; i++ {
		r1 := remainders[i]
		r2 := remainders[int(k)-i]
		// fmt.Printf("i=%d, r1=%d, r2=%d\n", i, r1, r2)
		if r1 > r2 {
			count += r1
		} else {
			count += r2
		}
	}

	// add one if k is even and we have an element
	// with remainder k/2
	if k%2 == 0 && remainders[k/2] != 0 {
		count++
	}

	// add one if we have a remainder 0
	if remainders[0] != 0 {
		count++
	}
	return count
}
