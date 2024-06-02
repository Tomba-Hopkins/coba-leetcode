package difficultyeasy

// Question :
// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Example 1:
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]
// Example 2:
// Input: s = ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
// Constraints:
// 1 <= s.length <= 105

// Answer :
func ReverseString(s []byte) {

	a, b := 0, len(s) - 1
	for a < b {
		s[a], s[b] = s[b], s[a]
		a++
		b--
	}

}

// testcase:
// text := []byte("hello")
