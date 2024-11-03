package difficultyeasy

// Question :

// Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
// A shift on s consists of moving the leftmost character of s to the rightmost position.
// For example, if s = "abcde", then it will be "bcdea" after one shift.
// Example 1:
// Input: s = "abcde", goal = "cdeab"
// Output: true
// Example 2:
// Input: s = "abcde", goal = "abced"
// Output: false
// Constraints:
// 1 <= s.length, goal.length <= 100
// s and goal consist of lowercase English letters.

// Answer :
func RotateString(s string, goal string) bool {
	newS := s + s
	r := []int{}
	for i := 0; i < len(newS); i++ {
		if goal[0] == newS[i] {
			r = append(r, i)
		}
	}
	for i := 0; i < len(r); i++ {
		d := r[i]
		if ((len(newS) - d) > len(goal)) && len(goal) == len(s) {
			if newS[d:d+len(goal)] == goal {
				return true
			}
		}
	}
	return false
}