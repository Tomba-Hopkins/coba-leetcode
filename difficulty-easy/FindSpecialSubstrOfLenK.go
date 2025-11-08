package difficultyeasy

// Question :

// You are given a string s and an integer k.
// Determine if there exists a substring of length exactly k in s that satisfies the following conditions:
// The substring consists of only one distinct character (e.g., "aaa" or "bbb").
// If there is a character immediately before the substring, it must be different from the character in the substring.
// If there is a character immediately after the substring, it must also be different from the character in the substring.
// Return true if such a substring exists. Otherwise, return false.
// Example 1:
// Input: s = "aaabaaa", k = 3
// Output: true
// Explanation:
// The substring s[4..6] == "aaa" satisfies the conditions.
// It has a length of 3.
// All characters are the same.
// The character before "aaa" is 'b', which is different from 'a'.
// There is no character after "aaa".
// Example 2:
// Input: s = "abc", k = 2
// Output: false
// Explanation:
// There is no substring of length 2 that consists of one distinct character and satisfies the conditions.
// Constraints:
// 1 <= k <= s.length <= 100
// s consists of lowercase English letters only.

// Answer :
// gpt answer
func FindSpecialSubstrOfLenK(s string, k int) bool {
    if len(s) == 1 {
        return true
    }
    for i := 0; i <= len(s)-k; i++ {
        x := s[i : i+k]
        valid := true
        for _, c := range x {
            if c != rune(x[0]) {
                valid = false
                break
            }
        }
        if !valid {
            continue
        }
        if i > 0 && s[i-1] == x[0] {
            continue
        }
        if i+k < len(s) && s[i+k] == x[0] {
            continue
        }

        return true
    }
    return false
}

// my wrong ans
// func FindSpecialSubstrOfLenK(s string, k int) bool {
// 	if len(s) == 1 {
// 		return true
// 	}
// 	for i := 0; i < len(s)-k; i++ {
// 		x := s[i : i+k]
// 		for j := i + k; j < len(s); j++ {
// 			x2 := ""
// 			if j+k < len(s) {
// 				x2 = s[j : j+k]
// 			} else {
// 				x2 = s[j:]
// 			}
// 			if x == x2 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
