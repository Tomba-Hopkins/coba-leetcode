package difficultyeasy

// Question :

// You are given an array of strings words (0-indexed).
// In one operation, pick two distinct indices i and j, where words[i] is a non-empty string, and move any character from words[i] to any position in words[j].
// Return true if you can make every string in words equal using any number of operations, and false otherwise.
// Example 1:
// Input: words = ["abc","aabc","bc"]
// Output: true
// Explanation: Move the first 'a' in words[1] to the front of words[2],
// to make words[1] = "abc" and words[2] = "abc".
// All the strings are now equal to "abc", so return true.
// Example 2:
// Input: words = ["ab","a"]
// Output: false
// Explanation: It is impossible to make all the strings equal using the operation.
// Constraints:
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] consists of lowercase English letters.

// Answer :
func RedistributeCharToMakeAllStrEqual(words []string) bool {
	n := map[string]int{}
	m := map[string]int{}
	for _, x := range words {
		for _, v := range x {
			n[string(v)] = 1
			m[string(v)]++
		}
	}
	t := ""
	for _, w := range words {
		for _, v := range w {
			if n[string(v)] > 0 {
				t += string(v)
				n[string(v)] = 0
			}
		}
	}
	for i := 0; i < len(t); i++ {
		if m[string(t[i])]%len(words) != 0 {
			return false
		}
	}
	return true
}
