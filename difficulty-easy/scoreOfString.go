package difficultyeasy

// Question :

// You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.
// Return the score of s.
// Example 1:
// Input: s = "hello"
// Output: 13
// Explanation:
// The ASCII values of the characters in s are: 'h' = 104, 'e' = 101, 'l' = 108, 'o' = 111. So, the score of s would be |104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13.
// Example 2:
// Input: s = "zaz"
// Output: 50
// Explanation:
// The ASCII values of the characters in s are: 'z' = 122, 'a' = 97. So, the score of s would be |122 - 97| + |97 - 122| = 25 + 25 = 50.

// Answer :
func ScoreOfString(s string) (res int) {

	text := "abcdefghijklmnopqrstuvwxyz"
	d := 97
	abjad := map[string]int{}

	for _, r := range text {
		abjad[string(r)] = d
		d++
	}

	for i := 0; i < len(s) - 1; i++{
		str := string(s[i])
		plus := (abjad[str] - abjad[string(s[i + 1])])
		if plus < 0{
			plus = -plus
		}
		res += plus
	}

	return 

}