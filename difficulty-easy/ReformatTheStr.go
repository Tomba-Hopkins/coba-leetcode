package difficultyeasy

// Question :

// You are given an alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).
// You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.
// Return the reformatted string or return an empty string if it is impossible to reformat the string.
// Example 1:
// Input: s = "a0b1c2"
// Output: "0a1b2c"
// Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
// Example 2:
// Input: s = "leetcode"
// Output: ""
// Explanation: "leetcode" has only characters so we cannot separate them by digits.
// Example 3:
// Input: s = "1229857369"
// Output: ""
// Explanation: "1229857369" has only digits so we cannot separate them by characters.
// Constraints:
// 1 <= s.length <= 500
// s consists of only lowercase English letters and/or digits.

// Answer :
func ReformatTheStr(s string) (r string) {
	if len(s) == 1 {
		return s
	}
	n, t := "", ""
	for _, x := range s {
		if x >= '0' && x <= '9' {
			n += string(x)
		} else if x >= 'a' && x <= 'z' {
			t += string(x)
		}
	}
	if len(n) < 1 || len(t) < 1 {
		return ""
	}
	if max(len(n), len(t))-min(len(n), len(t)) > 1 {
		return ""
	}
	x1, x2 := 0, 0
	for len(r) < len(s) {
		if len(t) > len(n) {
			if x2 < len(t) {
				r += string(t[x2])
				x2++
			}
			if x1 < len(n) {
				r += string(n[x1])
				x1++
			}
		} else {
			if x1 < len(n) {
				r += string(n[x1])
				x1++
			}
			if x2 < len(t) {
				r += string(t[x2])
				x2++
			}
		}
	}
	return
}