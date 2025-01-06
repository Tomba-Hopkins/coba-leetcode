package difficultymedium

// Question :

// You are given a 0-indexed array of strings words and a 2D array of integers queries.
// Each query queries[i] = [li, ri] asks us to find the number of strings present in the range li to ri (both inclusive) of words that start and end with a vowel.
// Return an array ans of size queries.length, where ans[i] is the answer to the ith query.
// Note that the vowel letters are 'a', 'e', 'i', 'o', and 'u'.
// Example 1:
// Input: words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
// Output: [2,3,0]
// Explanation: The strings starting and ending with a vowel are "aba", "ece", "aa" and "e".
// The answer to the query [0,2] is 2 (strings "aba" and "ece").
// to query [1,4] is 3 (strings "ece", "aa", "e").
// to query [1,1] is 0.
// We return [2,3,0].
// Example 2:
// Input: words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
// Output: [3,2,1]
// Explanation: Every string satisfies the conditions, so we return [3,2,1].
// Constraints:
// 1 <= words.length <= 105
// 1 <= words[i].length <= 40
// words[i] consists only of lowercase English letters.
// sum(words[i].length) <= 3 * 105
// 1 <= queries.length <= 105
// 0 <= li <= ri < words.length

// Answer :
func CountVowelStringsInRange(words []string, queries [][]int) (res []int) {
    vow := map[byte]bool{}
    x := "aiueo"
    for _, v := range x {
        vow[byte(v)] = true
    }
    prefix := make([]int, len(words) + 1)
    for i := 0; i < len(words); i++{
        l := words[i][0]
        r := words[i][len(words[i]) - 1]
        if vow[l] && vow[r] {
            prefix[i + 1] = prefix[i] + 1
        } else {
            prefix[i + 1] = prefix[i]
        }
    }
    for _, q := range queries {
        l, r := q[0], q[1]
        res = append(res, prefix[r + 1] - prefix[l])
    }
	return
}



// Time limit excedeed
// func CountVowelStringsInRange(words []string, queries [][]int) (r []int) {
//     vow := map[byte]bool{}
//     x := "aiueo"
//     for _, v := range x {
//         vow[byte(v)] = true
//     }
//     for i := 0; i < len(queries); i++{
//         d := 0
//         for j := queries[i][0]; j <= queries[i][1]; j++{
//             l := words[j][0]
//             r := words[j][len(words[j]) - 1]
//             if vow[l] && vow[r] || vow[l] && len(words[j]) == 1 {
//                 d++
//             }
//         }
//         r = append(r, d)
//     }
// 	return
// }