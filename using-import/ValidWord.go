package usingimportpackage

import "strings"

// Question :

// A word is considered valid if:
// It contains a minimum of 3 characters.
// It contains only digits (0-9), and English letters (uppercase and lowercase).
// It includes at least one vowel.
// It includes at least one consonant.
// You are given a string word.
// Return true if word is valid, otherwise, return false.
// Notes:
// 'a', 'e', 'i', 'o', 'u', and their uppercases are vowels.
// A consonant is an English letter that is not a vowel.
// Example 1:
// Input: word = "234Adas"
// Output: true
// Explanation:
// This word satisfies the conditions.
// Example 2:
// Input: word = "b3"
// Output: false
// Explanation:
// The length of this word is fewer than 3, and does not have a vowel.
// Example 3:
// Input: word = "a3$e"
// Output: false
// Explanation:
// This word contains a '$' character and does not have a consonant.
// Constraints:
// 1 <= word.length <= 20
// word consists of English uppercase and lowercase letters, digits, '@', '#', and '$'.

// Answer :
func ValidWord(word string) bool {
	if len(word) < 3 {
		return  false
	}
	vow, conson, etc := 0, 0, 0
	for _, w := range word {
		l := strings.ToLower(string(w))
		if l == "a" || l == "i" || l == "u" || l == "e" || l == "o" {
			vow++
            etc++
		} else if l >= "a" && l <= "z" {
			if l != "a" && l != "i" && l != "u" && l != "e" && l != "o"{
				conson++
                etc++
			}
		} else if w >= '0' && w <= '9' {
			etc++
		} else {
            return false
        }
	}
	return conson > 0 && vow > 0 && etc >= 3
}