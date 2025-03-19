package difficultyeasy

import (
	"slices"
	"strconv"
)

// Question :

// Given an alphanumeric string s, return the second largest numerical digit that appears in s, or -1 if it does not exist.
// An alphanumeric string is a string consisting of lowercase English letters and digits.
// Example 1:
// Input: s = "dfa12321afd"
// Output: 2
// Explanation: The digits that appear in s are [1, 2, 3]. The second largest digit is 2.
// Example 2:
// Input: s = "abc1111"
// Output: -1
// Explanation: The digits that appear in s are [1]. There is no second largest digit.
// Constraints:
// 1 <= s.length <= 500
// s consists of only lowercase English letters and digits.

// Answer :
func SecondLargestDigitInAStr(s string) (r int) {
	r = -1
	m := map[string]bool{}
	x := "0123456789"
	for _, v := range x {
		m[string(v)] = true
	}
	t := []int{}
	for _, k := range s {
		if m[string(k)] {
			n, _ := strconv.Atoi(string(k))
			t = append(t, n)
		}
	}
	if len(t) > 0 {
		max := slices.Max(t)
		for i := 0; i < len(t); i++{
			if t[i] > r && t[i] != max {
				r = t[i]
			}
			
		}
	}
	return 
}