package usingimportpackage

import "strings"

// Question :

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.
// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
// Constraints:
// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.

// Answer :

func ValidPalindrome(s string) bool {
	x := ""
	for i := 0; i < len(s); i++{
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			l := strings.ToLower(string(s[i]))
			x += l
		}
		
	}
	a, b := 0, len(x) - 1
	for i := a; i < len(x); i++{
		if x[a] != x[b] {
			return false
		} else {
			a++
			b--
		}
	}
	return true
}


// func ValidPalindrome(s string) bool {
// 	x := ""
// 	for i := 0; i < len(s); i++{
// 		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
// 			l := strings.ToLower(string(s[i]))
// 			x += l
// 		}
// 	}
// 	a, b := 0, len(x) - 1
// for a < b {
// 	if x[a] != x[b] {
// 		return false
// 	} else {
// 		a++
// 		b--
// 	}
// }
// 	return true
// }