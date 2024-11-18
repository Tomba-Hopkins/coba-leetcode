package usingimportpackage

import "strings"

// Question :

// Given a string s, reverse only all the vowels in the string and return it.
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
// Example 1:
// Input: s = "IceCreAm"
// Output: "AceCreIm"
// Explanation:
// The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".
// Example 2:
// Input: s = "leetcode"
// Output: "leotcede"
// Constraints:
// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.

// Answer :

func ReverseVowelsOfAString(s string) (r string) {
    vow :=  "aiueo"
    x := []string{}
    for i := len(s) - 1; i >= 0; i--{
        l := strings.ToLower(string(s[i]))
        if strings.Contains(vow, l){
            x = append(x, string(s[i]))
        }
    }
    d := 0
    for i := 0; i < len(s); i++{
        l := strings.ToLower(string(s[i]))
        if strings.Contains(vow, l){
            r += x[d]
            d++
        } else {
            r += string(s[i])
        }
    }
    return
}


