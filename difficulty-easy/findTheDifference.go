package difficultyeasy

// Question :

// You are given two strings s and t.
// String t is generated by random shuffling string s and then add one more letter at a random position.
// Return the letter that was added to t.
// Example 1:
// Input: s = "abcd", t = "abcde"
// Output: "e"
// Explanation: 'e' is the letter that was added.
// Example 2:
// Input: s = "", t = "y"
// Output: "y"
// Constraints:
// 0 <= s.length <= 1000
// t.length == s.length + 1
// s and t consist of lowercase English letters.

// Answer :

func FindTheDifference(s string, t string) byte {

	a := map[string]int{}
	b := map[string]int{}

	for _, i := range s {
		a[string(i)]++
	}
	for _, i := range t {
		b[string(i)]++
	}

	for i := 0; i < len(t); i++ {
		if a[string(t[i])] != b[string(t[i])] {
			return t[i]
		}
	}

	return '0'

}