package usingimportpackage

// Question :

// Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.
// Substrings that occur multiple times are counted the number of times they occur.
// Example 1:
// Input: s = "00110011"
// Output: 6
// Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
// Notice that some of these substrings repeat and are counted the number of times they occur.
// Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
// Example 2:
// Input: s = "10101"
// Output: 4
// Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
// Constraints:
// 1 <= s.length <= 105
// s[i] is either '0' or '1'.

// Answer :

// gpt ans
func CountBinStr(s string) (r int) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		l, h := i, i+1
		for l >= 0 && h < len(s) && s[l] == s[i] && s[h] == s[i+1] {
			r++
			l--
			h++
		}
	}
	return
}

// TLE
// func CountBinStr(s string) (r int) {
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			a := s[i : j+1]
// 			if len(a)%2 != 0 {
// 				continue
// 			}
// 			x, y := a[:len(a)/2], a[len(a)/2:]
// 			xF, yF, d := x[0], y[0], true
// 			if strings.Contains(x, string(yF)) {
// 				d = false
// 			}
// 			if strings.Contains(y, string(xF)) {
// 				d = false
// 			}
// 			if d {
// 				r++
// 			}
// 		}
// 	}
// 	return
// }
