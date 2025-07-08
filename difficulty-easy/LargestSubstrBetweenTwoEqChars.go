package difficultyeasy

// Question :

// Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
// Input: s = "aa"
// Output: 0
// Explanation: The optimal substring here is an empty substring between the two 'a's.
// Example 2:
// Input: s = "abca"
// Output: 2
// Explanation: The optimal substring here is "bc".
// Example 3:
// Input: s = "cbzxy"
// Output: -1
// Explanation: There are no characters that appear twice in s.
// Constraints:
// 1 <= s.length <= 300
// s contains only lowercase English letters.

// Answer :
func LargestSubstrBetweenTwoEqChars(s string) int {
	f, e, d := 0, 0, false
	k := ""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				d = true
				t := s[i+1 : j]
				if len(t) > len(k) {
					k = t
					f = i
					e = j
				}

				break
			}
		}

	}
	if !d {
		return -1
	} else if e-f == 1 {
		return 0
	}
	return len(k)
}