package difficultyeasy

// Question :

// Given a string s, return the maximum length of a substring such that it contains at most two occurrences of each character.
// Example 1:
// Input: s = "bcbbbcba"
// Output: 4
// Explanation:
// The following substring has a length of 4 and contains at most two occurrences of each character: "bcbbbcba".
// Example 2:
// Input: s = "aaaa"
// Output: 2
// Explanation:
// The following substring has a length of 2 and contains at most two occurrences of each character: "aaaa".
// Constraints:
// 2 <= s.length <= 100
// s consists only of lowercase English letters.

// Answer :
func MaxLenSubstrWithTwoOccurrences(s string) (r int) {
	for i := len(s); i > 1; i-- {
		for j := 0; j < len(s); j++ {
			x, m, d := "", map[string]int{}, false
			if i+j <= len(s) {
				x = s[j : i+j]
			} else {
				break
			}
			for _, v := range x {
				m[string(v)]++
			}
			for _, v := range x {
				if m[string(v)] > 2 {
					d = true
				}
			}
			if !d {
				return len(x)
			}
		}
	}
	return
}