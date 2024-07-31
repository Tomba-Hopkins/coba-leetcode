package difficultyeasy

// Question :

// Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.
// Example 1:
// Input: s = "Hello"
// Output: "hello"
// Example 2:
// Input: s = "here"
// Output: "here"
// Example 3:
// Input: s = "LOVELY"
// Output: "lovely"
// Constraints:
// 1 <= s.length <= 100
// s consists of printable ASCII characters.

// Answer :
func ToLowerCase(s string) (r string) {

	c := map[string]string{}
	
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(a); i++{
		c[string(a[i])] = string(b[i])
	}

	for _, x := range s {
		if c[string(x)] != "" {
			r += c[string(x)]
			continue
		}
		r += string(x)
	}
	return
}