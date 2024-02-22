package difficultyeasy

// Question :
// Example 1:
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
// Example 2:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
// Example 3:
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

// Answer :
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	temp := 0
	for i := 0; i < len(s); i++ {
		now := roman[s[i]]
		if i+1 < len(s) {
			next := roman[s[i+1]]
			if now < next {
				temp -= now
			} else {
				temp += now
			}
		} else {
			temp += now
		}
	}
	return temp
}