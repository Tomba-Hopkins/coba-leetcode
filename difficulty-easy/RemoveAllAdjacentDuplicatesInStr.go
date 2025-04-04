package difficultyeasy

// Question :

// You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.
// We repeatedly make duplicate removals on s until we no longer can.
// Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.
// Example 1:
// Input: s = "abbaca"
// Output: "ca"
// Explanation:
// For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
// Example 2:
// Input: s = "azxxzy"
// Output: "ay"
// Constraints:
// 1 <= s.length <= 105
// s consists of lowercase English letters.

// Answer :
func RemoveAllAdjacentDuplicatesInStr(s string) (r string) {
	for i := 0; i < len(s); i++ {
		if len(r) > 0 && r[len(r)-1] == s[i] {
			r = r[:len(r)-1]
		} else {
			r += string(s[i])
		}
	}
	return
}


// lebih hemat
// func RemoveAllAdjacentDuplicatesInStr(s string) string {
//     var r []rune
// 	for _, x := range s {
//         if len(r) > 0 && r[len(r) - 1] == x {
//             r = r[:len(r) - 1]
//         } else {
//             r = append(r, x)
//         }
//     }
// 	return string(r)
// }
