package difficultyeasy

// Question :

// Given a string s containing only lowercase English letters and the '?' character, convert all the '?' characters into lowercase letters such that the final string does not contain any consecutive repeating characters. You cannot modify the non '?' characters.
// It is guaranteed that there are no consecutive repeating characters in the given string except for '?'.
// Return the final string after all the conversions (possibly zero) have been made. If there is more than one solution, return any of them. It can be shown that an answer is always possible with the given constraints.
// Example 1:
// Input: s = "?zs"
// Output: "azs"
// Explanation: There are 25 solutions for this problem. From "azs" to "yzs", all are valid. Only "z" is an invalid modification as the string will consist of consecutive repeating characters in "zzs".
// Example 2:
// Input: s = "ubv?w"
// Output: "ubvaw"
// Explanation: There are 24 solutions for this problem. Only "v" and "w" are invalid modifications as the strings will consist of consecutive repeating characters in "ubvvw" and "ubvww".
// Constraints:
// 1 <= s.length <= 100
// s consist of lowercase English letters and '?'.

// Answer :
func ReplaceAllTandaTanya(s string) (r string) {
	x := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			for d := 0; d < 26; d++ {
				c := x[d]
				if (i > 0 && r[len(r)-1] == c) || (i < len(s)-1 && s[i+1] == c) {
					continue
				}
				r += string(c)
				break
			}
		} else {
			r += string(s[i])
		}
	}
	return
}

// func ReplaceAllTandaTanya(s string) (r string) {
// 	x := "abcdefghijklmnopqrstuvwxyz"
// 	for i := 0; i < len(s); i++{
// 		if s[i] == '?'{
// 			d := 0
// 			if len(r) > 0 {
// 				for r[len(r) - 1] == x[d] {
// 					d++
// 				}
// 			}
// 			if i + 1 < len(s) {
// 				for s[i + 1] == x[d] {
// 					d++
// 				}
// 			}
// 			r += string(x[d])
// 		} else {
// 			r += string(s[i])
// 		}
// 	}
// 	return
// }