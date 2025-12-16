package usingimportpackage

import "strings"

// Question :

// You are given a string s consisting of lowercase English words, each separated by a single space.
// Determine how many vowels appear in the first word. Then, reverse each following word that has the same vowel count. Leave all remaining words unchanged.
// Return the resulting string.
// Vowels are 'a', 'e', 'i', 'o', and 'u'.
// Example 1:
// Input: s = "cat and mice"
// Output: "cat dna mice"
// Explanation:​​​​​​​
// The first word "cat" has 1 vowel.
// "and" has 1 vowel, so it is reversed to form "dna".
// "mice" has 2 vowels, so it remains unchanged.
// Thus, the resulting string is "cat dna mice".
// Example 2:
// Input: s = "book is nice"
// Output: "book is ecin"
// Explanation:
// The first word "book" has 2 vowels.
// "is" has 1 vowel, so it remains unchanged.
// "nice" has 2 vowels, so it is reversed to form "ecin".
// Thus, the resulting string is "book is ecin".
// Example 3:
// Input: s = "banana healthy"
// Output: "banana healthy"
// Explanation:
// The first word "banana" has 3 vowels.
// "healthy" has 2 vowels, so it remains unchanged.
// Thus, the resulting string is "banana healthy".
// Constraints:
// 1 <= s.length <= 105
// s consists of lowercase English letters and spaces.
// Words in s are separated by a single space.
// s does not contain leading or trailing spaces.

// Answer :
func ReverseWordsWithSameVowelCount(s string) string {
	a, t, n := strings.Split(s, " "), []string{}, 0
	for i := 0; i < len(a); i++{
		x, k := 0, ""
		for _, v := range a[i] {
			if v == 'a' || v == 'i' || v == 'u' || v == 'e' || v == 'o'{
				x++
			}
		}
		if i == 0 {
			n = x
			k = revOrNo(a[i], false)
		} else {
			if x == n {
				k = revOrNo(a[i], true)
			} else {
				k = revOrNo(a[i], false)
			}
		}		
		t = append(t, k)
	}
	return strings.Join(t, " ")
}

func revOrNo(s string, isRev bool) string {
	var r strings.Builder
	r.Grow(len(s))
	if !isRev {
		for i := 0; i < len(s); i++{
			r.WriteByte(s[i])
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			r.WriteByte(s[i])
		}
	}
	return r.String()
}