package usingimportpackage

import "math"

// Question :

// An integer n is a power of three, if there exists an integer x such that n == 3x.
// Given an integer n, return true if it is a power of three. Otherwise, return false.
// Example 1:
// Input: n = 27
// Output: true
// Explanation: 27 = 33
// Example 2:
// Input: n = 0
// Output: false
// Explanation: There is no x where 3x = 0.
// Example 3:
// Input: n = -1
// Output: false
// Explanation: There is no x where 3x = (-1).
// Constraints:
// -231 <= n <= 231 - 1
// Follow up: Could you solve it without loops/recursion?

// Answer :
func PowerOfThree(n int) bool {
	x := float64(0)
	for int(math.Pow(3, x)) <= n {
		if int(math.Pow(3, x)) == n {
			return  true
		}
		x++
	}	
	return false
}