package difficultyeasy

// Question :

// A pangram is a sentence where every letter of the English alphabet appears at least once.
// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
// Example 1:
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.
// Example 2:
// Input: sentence = "leetcode"
// Output: false
// Constraints:
// 1 <= sentence.length <= 1000
// sentence consists of lowercase English letters.

// Answer :
func CheckIfPangram(sentence string) bool {
    a := "abcdefghijklmnopqrstuvwxyz"
    m := map[string]bool{}
    d := 0
    
    for _, r := range a {
        m[string(r)] = true
    }
    for _, v := range sentence {
        if m[string(v)] {
            m[string(v)] = false
            d++
        }
    }
    
	return d == len(m)
}