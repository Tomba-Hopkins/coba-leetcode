package difficultyeasy

// Question :

// Given two strings s and t, return true if t is an
// anagram of s, and false otherwise.
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:
// Input: s = "rat", t = "car"
// Output: false
// Constraints:
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

// Answer :
func ValidAnagram(s string, t string) bool {
	m := map[string]int{}
	for _, k := range t {
		m[string(k)]++
	}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] > 0 {
			m[string(s[i])]--
		} else {
			return false
		}
	}
	return len(s) == len(t)
}