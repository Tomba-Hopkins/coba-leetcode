package difficultyeasy

// Question :

// You are given a string s consisting of n characters which are either 'X' or 'O'.
// A move is defined as selecting three consecutive characters of s and converting them to 'O'. Note that if a move is applied to the character 'O', it will stay the same.
// Return the minimum number of moves required so that all the characters of s are converted to 'O'.
// Example 1:
// Input: s = "XXX"
// Output: 1
// Explanation: XXX -> OOO
// We select all the 3 characters and convert them in one move.
// Example 2:
// Input: s = "XXOX"
// Output: 2
// Explanation: XXOX -> OOOX -> OOOO
// We select the first 3 characters in the first move, and convert them to 'O'.
// Then we select the last 3 characters and convert them so that the final string contains all 'O's.
// Example 3:
// Input: s = "OOOO"
// Output: 0
// Explanation: There are no 'X's in s to convert.
// Constraints:
// 3 <= s.length <= 1000
// s[i] is either 'X' or 'O'.

// Answer :
func MinimumMovesToConvStr(s string) (r int) {
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			r++
			i += 3
		} else {
			i++
		}

	}
	return
}

// func MinimumMovesToConvStr(s string) (r int) {
// 	for i := 0; i < len(s) - 1; i++{
// 		if s[i] == 'X' {
// 			s= s[i:]
// 			break
// 		}
// 	}
// 	for i := len(s) - 1; i > 1; i-- {
// 		if s[i] == 'X' {
// 			s = s[:i + 1]
// 			break
// 		}
// 	}
// 	for i := 0; i <= len(s); i += 3 {
// 		if s[i] == 'O' {
// 			continue
// 		}
// 		if i + 3 < len(s) {
// 			a := s[i: i+3]
// 			if strings.Contains(a, "X") {
// 				r++
// 			}
// 		} else {
// 			a := s[i:]
// 			if strings.Contains(a, "X") {
// 				r++
// 			}
// 		}
// 	}
// 	return
// }