package usingimportpackage

import "strconv"

// Question :

// Given an integer n, add a dot (".") as the thousands separator and return it in string format.
// Example 1:
// Input: n = 987
// Output: "987"
// Example 2:
// Input: n = 1234
// Output: "1.234"
// Constraints:
// 0 <= n <= 231 - 1

// Answer :
func ThousandSeparator(n int) (r string) {
	s, d, temp := strconv.Itoa(n), 0, ""
	for i := len(s) - 1; i >= 0; i--{
		temp += string(s[i])
		d++
		if d == 3 && i != 0 {
			temp += "."
			d = 0
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		r += string(temp[i])
	}
	return
}