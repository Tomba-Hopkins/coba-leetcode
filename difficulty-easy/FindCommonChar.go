package difficultyeasy

// Question :

// Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.
// Example 1:
// Input: words = ["bella","label","roller"]
// Output: ["e","l","l"]
// Example 2:
// Input: words = ["cool","lock","cook"]
// Output: ["c","o"]
// Constraints:
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] consists of lowercase English letters.

// Answer :

// gpt fixed
func FindCommonChar(words []string) (r []string) {
	h, w := map[string]int{}, ""
	for _, v := range words[0] {
		h[string(v)]++
	}
	for k := range h {
		w += k
	}
	for i := 0; i < len(w); i++ {
		target, count := string(w[i]), h[string(w[i])]
		for j := 1; j < len(words); j++ {
			m := map[string]int{}
			for _, v := range words[j] {
				m[string(v)]++
			}
			if m[target] < count {
				count = m[target]
			}
		}
		for j := 0; j < count; j++ {
			r = append(r, target)
		}
	}
	return
}

// my ans before
// func FindCommonChar(words []string) (r []string) {
// 	h, w := map[string]bool{}, ""
// 	for _, v := range words[0] {
// 		if !h[string(v)] {
// 			w += string(v)
// 		}
// 		h[string(v)] = true
// 	}
// 	for i := 0; i < len(w); i++ {
// 		target, count := string(w[i]), len(words)
// 		for j := 0; j < len(words); j++ {
// 			m := map[string]int{}
// 			for _, v := range words[j] {
// 				m[string(v)]++
// 			}
// 			if m[target] < count {
// 				count = m[target]
// 			}
// 		}
// 		for j := 0; j < count; j++ {
// 			r = append(r, target)
// 		}
// 	}
// 	return
// }
