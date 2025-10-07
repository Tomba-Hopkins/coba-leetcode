package usingimportpackage

import "strconv"

// Question :

// You are given a positive integer n.
// Let even denote the number of even indices in the binary representation of n with value 1.
// Let odd denote the number of odd indices in the binary representation of n with value 1.
// Note that bits are indexed from right to left in the binary representation of a number.
// Return the array [even, odd].
// Example 1:
// Input: n = 50
// Output: [1,2]
// Explanation:
// The binary representation of 50 is 110010.
// It contains 1 on indices 1, 4, and 5.
// Example 2:
// Input: n = 2
// Output: [0,1]
// Explanation:
// The binary representation of 2 is 10.
// It contains 1 only on index 1.
// Constraints:
// 1 <= n <= 1000

// Answer :
func NumOfEvenAndOddBits(n int) []int {
	s, newS := strconv.FormatInt(int64(n), 2), ""
	for i := len(s) - 1; i >= 0; i--{
		newS += string(s[i])
	}
	e, o := 0, 0
	for i := 0; i < len(newS); i++{
		if newS[i] == '1'{
			if i % 2 == 0 {
				e++
			} else {
				o++
			}
		}
	}
	return []int{e, o}
}