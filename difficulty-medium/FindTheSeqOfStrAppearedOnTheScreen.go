package difficultymedium

// Question :

// You are given a string target.
// Alice is going to type target on her computer using a special keyboard that has only two keys:
// Key 1 appends the character "a" to the string on the screen.
// Key 2 changes the last character of the string on the screen to its next character in the English alphabet. For example, "c" changes to "d" and "z" changes to "a".
// Note that initially there is an empty string "" on the screen, so she can only press key 1.
// Return a list of all strings that appear on the screen as Alice types target, in the order they appear, using the minimum key presses.
// Example 1:
// Input: target = "abc"
// Output: ["a","aa","ab","aba","abb","abc"]
// Explanation:
// The sequence of key presses done by Alice are:
// Press key 1, and the string on the screen becomes "a".
// Press key 1, and the string on the screen becomes "aa".
// Press key 2, and the string on the screen becomes "ab".
// Press key 1, and the string on the screen becomes "aba".
// Press key 2, and the string on the screen becomes "abb".
// Press key 2, and the string on the screen becomes "abc".
// Example 2:
// Input: target = "he"
// Output: ["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]
// Constraints:
// 1 <= target.length <= 400
// target consists only of lowercase English letters.

// Answer :
func FindTheSeqOfStrAppearedOnTheScreen(target string) (r []string) {
	d := ""
	for i := 0; i < len(target); i++{
		x := 'a'
		for x < rune(target[i]) {
			if len(d) > 0 {
				r = append(r, d + string(x))
			} else {
				r = append(r, string(x))
			}
			x++
		}
		if len(d) > 0 {
			r = append(r, d + string(x))
		} else {
			r = append(r, string(x))
		}
		d += string(x)
		
	}
	return
}