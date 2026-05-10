package weeklycontest491

// Question :

// You are given a string s that consists of lowercase English letters.
// Return the string obtained by removing all trailing vowels from s.
// The vowels consist of the characters 'a', 'e', 'i', 'o', and 'u'.
//
// Example 1:
// Input: s = "idea"
// Output: "id"
// Explanation:
// Removing "idea", we obtain the string "id".
// Example 2:
// Input: s = "day"
// Output: "day"
// Explanation:
// There are no trailing vowels in the string "day".
// Example 3:
// Input: s = "aeiou"
// Output: ""
// Explanation:
// Removing "aeiou", we obtain the string "".
//
// Constraints:
// 1 <= s.length <= 100
// s consists of only lowercase English letters.

// Answer :

func TrimTrailingVowels(s string) string {
	x, vows, m := 0, "aiueo", map[string]bool{}
	for _, v := range vows {
		m[string(v)] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		if !m[string(s[i])] {
			x = i
			break
		}
		if i == 0 {
			return ""
		}
	}
	return string(s[0 : x+1])
}
