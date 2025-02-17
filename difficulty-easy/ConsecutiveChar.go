package difficultyeasy

// Question :

// The power of the string is the maximum length of a non-empty substring that contains only one unique character.
// Given a string s, return the power of s.
// Example 1:
// Input: s = "leetcode"
// Output: 2
// Explanation: The substring "ee" is of length 2 with the character 'e' only.
// Example 2:
// Input: s = "abbcccddddeeeeedcba"
// Output: 5
// Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
// Constraints:
// 1 <= s.length <= 500
// s consists of only lowercase English letters.

// Answer :
func ConsecutiveChar(s string) int {
	d := s[0]
	n := 0
	t := 0
	for i := 0; i < len(s); i++ {
		if d == s[i] {
			t++
		} else {
			d = s[i]
			t = 1
		}
		if t > n {
			n = t
		}
	}
	return n
}