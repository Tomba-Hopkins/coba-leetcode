package difficultyeasy

// Question :
// Given a string s which consists of lowercase or uppercase letters, return the length of the longest
// palindrome
//  that can be built with those letters.
// Letters are case sensitive, for example, "Aa" is not considered a palindrome.
// Example 1:
// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
// Example 2:
// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.

// Answer :
func LongestPalindrome(s string) (res int) {

	mapS := map[string]int{}

	for _, r := range s{
		mapS[string(r)]++
	}
	adaYangGaGenapCoy := false

	for _, isi := range mapS{
		if isi % 2 == 0 {
			res += isi
		} else {
			res += isi - 1
			adaYangGaGenapCoy = true
		}
	}

	if adaYangGaGenapCoy{
		res++
	}
	return
}