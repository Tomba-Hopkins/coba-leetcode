package difficultyeasy

// Question :
// We define the usage of capitals in a word to be right when one of the following cases holds:
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Given a string word, return true if the usage of capitals in it is right.
// Example 1:
// Input: word = "USA"
// Output: true
// Example 2:
// Input: word = "FlaG"
// Output: false
// Constraints:
// 1 <= word.length <= 100
// word consists of lowercase and uppercase English letters.

// Answer :
func DetectCapital(word string) bool {
    cap := map[string]bool{}
    x := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for _, c := range x {
        cap[string(c)] = true
    }
    d := 0
    for i := 0; i < len(word); i++{
        if cap[string(word[i])]{
            d++
        }
    }
	return d == len(word) || cap[string(word[0])] && d == 1 || d == 0
}