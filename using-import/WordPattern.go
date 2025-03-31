package usingimportpackage

import "strings"

// Question :

// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:
// Each letter in pattern maps to exactly one unique word in s.
// Each unique word in s maps to exactly one letter in pattern.
// No two letters map to the same word, and no two words map to the same letter.
// Example 1:
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Explanation:
// The bijection can be established as:
// 'a' maps to "dog".
// 'b' maps to "cat".
// Example 2:
// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false
// Example 3:
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false
// Constraints:
// 1 <= pattern.length <= 300
// pattern contains only lower-case English letters.
// 1 <= s.length <= 3000
// s contains only lowercase English letters and spaces ' '.
// s does not contain any leading or trailing spaces.
// All the words in s are separated by a single space.

// Answer :
func WordPattern(pattern string, s string) bool {
	m := map[string]int{}
	for _, p := range pattern {
		if m[string(p)] == 0 {
			m[string(p)] = len(m) + 1
		}
	}
	a := strings.Split(s, " ")
	m2 := map[string]int{}
	for _, p := range a {
		if m2[p] == 0 {
			m2[p] = len(m2) + 1
		}
	}
	if len(a) != len(pattern) {
		return false
	}
	for i := 0; i < len(a); i++{
		if m[string(pattern[i])] != m2[a[i]] {
			return false
		}	
	}
	return true
}

// func WordPattern(pattern string, s string) bool {
//     x := strings.Split(s, " ")
//     m := map[string]string{}
//     for i := 0; i < len(pattern); i++{
//         m[string(pattern[i])] = x[i]
//         m[x[i]] = string(pattern[i])
//     }
//     for i := 0; i < len(pattern); i++{
//         if m[string(pattern[i])] != x[i] || m[x[i]] != string(pattern[i]) {
//             return false
//         }
//     }
//     return true
// }