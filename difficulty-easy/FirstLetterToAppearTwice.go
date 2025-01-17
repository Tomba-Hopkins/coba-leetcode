package difficultyeasy

// Question :

// Given a string s consisting of lowercase English letters, return the first letter to appear twice.
// Note:
// A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
// s will contain at least one letter that appears twice.
// Example 1:
// Input: s = "abccbaacz"
// Output: "c"
// Explanation:
// The letter 'a' appears on the indexes 0, 5 and 6.
// The letter 'b' appears on the indexes 1 and 4.
// The letter 'c' appears on the indexes 2, 3 and 7.
// The letter 'z' appears on the index 8.
// The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.
// Example 2:
// Input: s = "abcdd"
// Output: "d"
// Explanation:
// The only letter that appears twice is 'd' so we return 'd'.
// Constraints:
// 2 <= s.length <= 100
// s consists of lowercase English letters.
// s has at least one repeated letter.

// Answer :

func FirstLetterToAppearTwice(s string) byte {
	x := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if x[s[i]] {
			return s[i]
		}
		x[s[i]] = true
	}
	return '0'
}


// before:
// func repeatedCharacter(s string) byte {
//     x := "x"
//     d := len(s)
//     for i := 0; i < len(s) - 1; i++{
//         for j := i + 1; j < len(s); j++{
//             if s[i] == s[j]{
//                 if j - i == 1 {     
//                     return s[i]
//                 }     
//                 if j - i < d {
//                     d = j - 1
//                     x = string(s[i])
//                 }       
//             }
//         }
//     }
// 	return x[0]
// }