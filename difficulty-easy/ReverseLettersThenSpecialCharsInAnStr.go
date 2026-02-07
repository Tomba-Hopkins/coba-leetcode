package difficultyeasy

// Question :

// You are given a string s consisting of lowercase English letters and special characters.
// Your task is to perform these in order:
// Reverse the lowercase letters and place them back into the positions originally occupied by letters.
// Reverse the special characters and place them back into the positions originally occupied by special characters.
// Return the resulting string after performing the reversals.
// Example 1:
// Input: s = ")ebc#da@f("
// Output: "(fad@cb#e)"
// Explanation:
// The letters in the string are ['e', 'b', 'c', 'd', 'a', 'f']:
// Reversing them gives ['f', 'a', 'd', 'c', 'b', 'e']
// s becomes ")fad#cb@e("
// ​​​​​​​The special characters in the string are [')', '#', '@', '(']:
// Reversing them gives ['(', '@', '#', ')']
// s becomes "(fad@cb#e)"
// Example 2:
// Input: s = "z"
// Output: "z"
// Explanation:
// The string contains only one letter, and reversing it does not change the string. There are no special characters.
// Example 3:
// Input: s = "!@#$%^&*()"
// Output: ")(*&^%$#@!"
// Explanation:
// The string contains no letters. The string contains all special characters, so reversing the special characters reverses the whole string.
// Constraints:
// 1 <= s.length <= 100
// s consists only of lowercase English letters and the special characters in "!@#$%^&*()".

// Answer :
func ReverseLettersThenSpecialCharsInAnStr(s string) (r string) {
	chars, specials := []string{}, []string{}
	for _, v := range s {
		if v <= 'z' && v >= 'a' {
			chars = append(chars, string(v))
		} else {
			specials = append(specials, string(v))
		}
	}
	idxCh, idxSp := len(chars)-1, len(specials)-1
	for _, v := range s {
		if v <= 'z' && v >= 'a' {
			r += chars[idxCh]
			idxCh--
		} else {
			r += specials[idxSp]
			idxSp--
		}
	}
	return
}
