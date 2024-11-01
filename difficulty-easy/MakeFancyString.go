package difficultyeasy

// Question :

// A fancy string is a string where no three consecutive characters are equal.
// Given a string s, delete the minimum possible number of characters from s to make it fancy.
// Return the final string after the deletion. It can be shown that the answer will always be unique.
// Example 1:
// Input: s = "leeetcode"
// Output: "leetcode"
// Explanation:
// Remove an 'e' from the first group of 'e's to create "leetcode".
// No three consecutive characters are equal, so return "leetcode".
// Example 2:
// Input: s = "aaabaaaa"
// Output: "aabaa"
// Explanation:
// Remove an 'a' from the first group of 'a's to create "aabaaaa".
// Remove two 'a's from the second group of 'a's to create "aabaa".
// No three consecutive characters are equal, so return "aabaa".
// Example 3:
// Input: s = "aab"
// Output: "aab"
// Explanation: No three consecutive characters are equal, so return "aab".
// Constraints:
// 1 <= s.length <= 105
// s consists only of lowercase English letters.

// Answer :
func MakeFancyString(s string) string {
    var r []rune
	x := ""
	d := 0
	for i := 0; i < len(s); i++ {
		if x != string(s[i]) {
			d = 0
			x = string(s[i])
			r = append(r, rune(s[i]))
		} else if x == string(s[i]) && d < 1 {
			r = append(r, rune(s[i]))
			d++
		} else {
			d++
		}
	}
	return string(r)
}

// My ans before [time limit exceeded] :
// func MakeFancyString(s string) (r string) {
//     x := ""
//     d := 0
//     for i := 0; i < len(s); i++{
//         if x != string(s[i]){
//             d = 0
//             x = string(s[i])
//             r += string(s[i])
//         } else if x == string(s[i]) && d < 1 {
//             r += string(s[i])
//             d++
//         } else {
//             d++
//         }
//     }
//     return
// }
