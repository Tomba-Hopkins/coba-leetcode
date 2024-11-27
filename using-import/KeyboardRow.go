package usingimportpackage

import "strings"

// Question :

// Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.
// Note that the strings are case-insensitive, both lowercased and uppercased of the same letter are treated as if they are at the same row.
// In the American keyboard:
// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".
// Example 1:
// Input: words = ["Hello","Alaska","Dad","Peace"]
// Output: ["Alaska","Dad"]
// Explanation:
// Both "a" and "A" are in the 2nd row of the American keyboard due to case insensitivity.
// Example 2:
// Input: words = ["omk"]
// Output: []
// Example 3:
// Input: words = ["adsdf","sfd"]
// Output: ["adsdf","sfd"]
// Constraints:
// 1 <= words.length <= 20
// 1 <= words[i].length <= 100
// words[i] consists of English letters (both lowercase and uppercase).

// Answer :

func KeyboardRow(words []string) (r []string) {
    key := []string{"qwertyuiop","asdfghjkl", "zxcvbnm"}
    m := map[string]int{}
    for i, v := range key {
        for _, w := range v {
            m[string(w)] = i + 1
        }
    }
    for _, x := range words {
        d := 0
        low := strings.ToLower(x)
        f := m[string(low[0])]
        for _, y := range low {
            if m[string(y)] == f {
                d++
            }
        }
        if d == len(x) {
            r = append(r, x)
        }
    }
	return
}