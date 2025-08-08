package difficultyeasy

// Question :

// Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
// Note that after backspacing an empty text, the text will continue empty.
// Example 1:
// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".
// Example 2:
// Input: s = "ab##", t = "c#d#"
// Output: true
// Explanation: Both s and t become "".
// Example 3:
// Input: s = "a#c", t = "b"
// Output: false
// Explanation: s becomes "c" while t becomes "b".
// Constraints:
// 1 <= s.length, t.length <= 200
// s and t only contain lowercase letters and '#' characters.
// Follow up: Can you solve it in O(n) time and O(1) space?

// Answer :
func BackspaceStrCompare(s string, t string) bool {
	x, y := "", ""
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(x) > 0 {
				x = x[:len(x)-1]
			} else {
				continue
			}
		} else {
			x += string(s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(y) > 0 {
				y = y[:len(y)-1]
			} else {
				continue
			}
		} else {
			y += string(t[i])
		}
	}
	return x == y
}