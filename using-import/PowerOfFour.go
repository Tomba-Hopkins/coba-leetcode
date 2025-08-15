package usingimportpackage

import "math"

// Question :

// Given an integer n, return true if it is a power of four. Otherwise, return false.
// An integer n is a power of four, if there exists an integer x such that n == 4x.
// Example 1:
// Input: n = 16
// Output: true
// Example 2:
// Input: n = 5
// Output: false
// Example 3:
// Input: n = 1
// Output: true
// Constraints:
// -231 <= n <= 231 - 1
// Follow up: Could you solve it without loops/recursion?

// Answer :
func PowerOfFour(n int) bool {
	x := 0
	for int(math.Pow(float64(4), float64(x))) <= n{
		if int(math.Pow(float64(4), float64(x))) == n {
			return  true
		}
		x++
	}
	return false
}