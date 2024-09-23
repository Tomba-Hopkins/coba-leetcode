package difficultyeasy

// // Question :
// Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.
// A string is represented by an array if the array elements concatenated in order forms the string.
// Example 1:
// Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
// Output: true
// Explanation:
// word1 represents string "ab" + "c" -> "abc"
// word2 represents string "a" + "bc" -> "abc"
// The strings are the same, so return true.
// Example 2:
// Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
// Output: false
// Example 3:
// Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
// Output: true
// Constraints:
// 1 <= word1.length, word2.length <= 103
// 1 <= word1[i].length, word2[i].length <= 103
// 1 <= sum(word1[i].length), sum(word2[i].length) <= 103
// word1[i] and word2[i] consist of lowercase letters.

// Answer :
func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	x, y := "", ""
	for _, w := range word1 {
		x += w
	}
	for _, w := range word2 {
		y += w
	}
	return x == y
}

// simple ans : 
// func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
// 	return strings.Join(word1, "") == strings.Join(word2, "")
// }

// my old ans (lol) :
// func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
// 	k := map[string]int{}
// 	k2 := map[string]int{}
// 	for i := 0; i < len(word1); i++ {
// 		for j := 0; j < len(word1[i]); j++ {
// 			v := string(word1[i][j])
// 			if k[v] == 0 {
// 				k[v] = len(k) + 1
// 			}
// 		}
// 	}
// 	for i := 0; i < len(word2); i++ {
// 		for j := 0; j < len(word2[i]); j++ {
// 			v := string(word2[i][j])
// 			if k2[v] == 0 {
// 				k2[v] = len(k2) + 1
// 			}
// 		}

// 	}
// 	t := ""
// 	t2 := ""

// 	for i := 0; i < len(word1); i++ {
// 		for j := 0; j < len(word1[i]); j++ {
// 			t += string(word1[i][j])
// 		}

// 	}
// 	for i := 0; i < len(word2); i++ {
// 		for j := 0; j < len(word2[i]); j++ {
// 			t2 += string(word2[i][j])
// 		}

// 	}

// 	if len(t) > len(t2) {
// 		for _, r := range t {
// 			if k[string(r)] != k2[string(r)] {
// 				return false
// 			}
// 		}
// 	} else {
// 		for _, r := range t2 {
// 			if k[string(r)] != k2[string(r)] {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }