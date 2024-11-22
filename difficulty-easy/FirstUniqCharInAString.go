package difficultyeasy

// Question :

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
// Example 1:
// Input: s = "leetcode"
// Output: 0
// Explanation:
// The character 'l' at index 0 is the first character that does not occur at any other index.
// Example 2:
// Input: s = "loveleetcode"
// Output: 2
// Example 3:
// Input: s = "aabb"
// Output: -1
// Constraints:
// 1 <= s.length <= 105
// s consists of only lowercase English letters.

// Answer :
func FirstUniqCharInAString(s string) int {
	m := map[string]int{}
	for _, r := range s {
		m[string(r)]++
	}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] == 1 {
			return i
		}
	}
	return -1
}