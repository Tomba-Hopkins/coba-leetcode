package usingimportpackage

import "strconv"

// Question :

// Reverse bits of a given 32 bits signed integer.
// Example 1:
// Input: n = 43261596
// Output: 964176192
// Explanation:
// Integer	Binary
// 43261596	00000010100101000001111010011100
// 964176192	00111001011110000010100101000000
// Example 2:
// Input: n = 2147483644
// Output: 1073741822
// Explanation:
// Integer	Binary
// 2147483644	01111111111111111111111111111100
// 1073741822	00111111111111111111111111111110
// Constraints:
// 0 <= n <= 231 - 2
// n is even.
// Follow up: If this function is called many times, how would you optimize it?

// Answer :
func ReverseBits(n int) int {
	s, x := strconv.FormatInt(int64(n), 2), ""
	for len(s) < 32{
		s = "0" + s
	}
	for i := len(s) - 1; i >= 0; i-- {
		x += string(s[i])
	}
	r, _ := strconv.ParseInt(x, 2, 64)
	return int(r)
}