package difficultymedium

import "sort"

// Question :

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Explanation:
// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:
// Input: strs = [""]
// Output: [[""]]
// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]
// Constraints:
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

// Answer :
func GroupAnagrams(strs []string) (r [][]string) {
    t := []string{}

    for _, s := range strs {
        rs := []rune(s)
        sort.Slice(rs, func(i, j int) bool {
            return rs[i] < rs[j]
        })
        t = append(t, string(rs))
    }
	
    m := map[int]bool{}
    for i := 0; i < len(t); i++{
        temp := []string{}
        for j := 0; j < len(t); j++{
            if t[i] == t[j] && !m[j] {
                m[j] = true
                temp = append(temp, strs[j])
            }   
        }
        if len(temp) > 0 {
            r = append(r, temp)
        }   
    }
    for _, v := range r {
        sort.Strings(v)   
    }
    sort.Slice(r, func(i, j int) bool {
        return len(r[i]) < len(r[j])
    })
    return
}