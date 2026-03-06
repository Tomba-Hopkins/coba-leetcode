package difficultyeasy

// Question :

// Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.
// Example 1:
// Input: s = "1001"
// Output: false
// Explanation: The ones do not form a contiguous segment.
// Example 2:
// Input: s = "110"
// Output: true
// Constraints:
// 1 <= s.length <= 100
// s[i]​​​​ is either '0' or '1'.
// s[0] is '1'.

// Answer :
func CheckIfBinStrHasAtMostOneSegOfOnes(s string) bool {
	b, b2 := false, false
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			b = true
			if b2 && b {
				return false
			}
		} else {
			if b {
				b2 = true
			}
		}
	}
	return true
}
