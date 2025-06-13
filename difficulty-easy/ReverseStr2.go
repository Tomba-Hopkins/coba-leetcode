package difficultyeasy

// Question :

// Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.
// If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.
// Example 1:
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"
// Example 2:
// Input: s = "abcd", k = 2
// Output: "bacd"
// Constraints:
// 1 <= s.length <= 104
// s consists of only lowercase English letters.
// 1 <= k <= 104

// // Answer :
func ReverseStr2(s string, k int) (r string) {
	a := []string{}
	for i := 0; i < len(s); i += k * 2 {
		x := ""
		if i + (k * 2) < len(s) {
			x = s[i : i + (k * 2)]
			a = append(a, x)
		} else {
			x = s[i:]
			a = append(a, x)
		}
	}

	for i := 0; i < len(a); i++{
		t, rev := "", ""
		if len(a[i]) < k {
			rev = a[i]
		} else {
            rev = a[i][:k]
        }
		for j := len(rev) - 1; j >= 0; j-- {
			t += string(rev[j])
		}
		for j := k; j < len(a[i]); j++{
			t += string(a[i][j])
		}
		r += t	
	}
	return
}