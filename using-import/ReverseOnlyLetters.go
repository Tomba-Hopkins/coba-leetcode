package usingimportpackage

import "strings"

// Question :

// Given a string s, reverse the string according to the following rules:
// All the characters that are not English letters remain in the same position.
// All the English letters (lowercase or uppercase) should be reversed.
// Return s after reversing it.
// Example 1:
// Input: s = "ab-cd"
// Output: "dc-ba"
// Example 2:
// Input: s = "a-bC-dEf-ghIj"
// Output: "j-Ih-gfE-dCba"
// Example 3:
// Input: s = "Test1ng-Leet=code-Q!"
// Output: "Qedo1ct-eeLg=ntse-T!"
// Constraints:
// 1 <= s.length <= 100
// s consists of characters with ASCII values in the range [33, 122].
// s does not contain '\"' or '\\'.

// Answer :
func ReverseOnlyLetters(s string) (r string) {
	w := "abcdefghijklmnopqrstuvwxyz"
    m := map[string]bool{}
    for _, r := range w {
        m[strings.ToLower(string(r))] = true
    }
    t := ""
    d := 0
    for i := len(s) - 1; i >= 0; i--{
        if m[strings.ToLower(string(s[i]))]{
            t += string(s[i])
        }
    }
    for _, x := range s {
        if m[strings.ToLower(string(x))]{
            r += string(t[d])
            d++
        } else {
            r += string(x)
        }
    }
    return
}