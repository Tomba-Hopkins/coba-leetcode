package usingimportpackage

import (
	"strconv"
	"strings"
)

// Question :

// You are given a positive integer n.
// A binary string x is valid if all substrings of x of length 2 contain at least one "1".
// Return all valid strings with length n, in any order.
// Example 1:
// Input: n = 3
// Output: ["010","011","101","110","111"]
// Explanation:
// The valid strings of length 3 are: "010", "011", "101", "110", and "111".
// Example 2:
// Input: n = 1
// Output: ["0","1"]
// Explanation:
// The valid strings of length 1 are: "0" and "1".
// Constraints:
// 1 <= n <= 18

// Answer :
func GenBinStrWithoutAdjZeros(n int) (r []string) {
	if n == 1 {
		return []string{"0", "1"}
	}
	x := 0
	for 0 < 1 {
		s := strconv.FormatInt(int64(x), 2)
		if len(s) > n{
			break
		}
		if !strings.Contains(s, "00") && strings.Contains(s, "1")  {
				if len(s) == n || len(s) == n -1 {
					if len(s) == n-1 {
					r = append(r, "0" + s)
				} else {
					r = append(r, s)
				}
				}
			}
		x++
	}
	return
}