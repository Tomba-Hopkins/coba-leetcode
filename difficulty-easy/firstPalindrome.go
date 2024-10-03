package difficultyeasy

// Question :

// Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".
// A string is palindromic if it reads the same forward and backward.
// Example 1:
// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.
// Example 2:
// Input: words = ["notapalindrome","racecar"]
// Output: "racecar"
// Explanation: The first and only string that is palindromic is "racecar".
// Example 3:
// Input: words = ["def","ghi"]
// Output: ""
// Explanation: There are no palindromic strings, so the empty string is returned.
// Constraints:
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] consists only of lowercase English letters.

// Answer :
func FirstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		invalid := false
		f := 0
		e := len(word) - 1
		for j := 0; j < len(word); j++ {
			if word[f] != word[e] {
				invalid = true
			}
			f++
			e--
		}
		if invalid {
			continue
		}
		return word
	}
	return ""
}


// my ans before :
// func firstPalindrome(words []string) string {
// 	for i := 0; i < len(words); i++{
// 		if words[i][0] == words[i][len(words[i]) - 1] {
// 			return words[i]
// 		}
		
// 	}
// 	return ""
// }