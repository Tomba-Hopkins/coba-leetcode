package usingimportpackage

import (
	"sort"
)

// Question :
// You are given an array of strings names, and an array heights that consists of distinct positive integers. Both arrays are of length n.
// For each index i, names[i] and heights[i] denote the name and height of the ith person.
// Return names sorted in descending order by the people's heights.
// Example 1:
// Input: names = ["Mary","John","Emma"], heights = [180,165,170]
// Output: ["Mary","Emma","John"]
// Explanation: Mary is the tallest, followed by Emma and John.
// Example 2:
// Input: names = ["Alice","Bob","Bob"], heights = [155,185,150]
// Output: ["Bob","Alice","Bob"]
// Explanation: The first Bob is the tallest, followed by Alice and the second Bob.
// Constraints:
// n == names.length == heights.length
// 1 <= n <= 103
// 1 <= names[i].length <= 20
// 1 <= heights[i] <= 105
// names[i] consists of lower and upper case English letters.
// All the values of heights are distinct.

// Answer :
func SortPeople(names []string, heights []int) (r []string) {
	m := map[int]string{}
	for i := 0; i < len(names); i++{
		m[heights[i]] = names[i]
	}
	sort.Ints(heights)
	for i := len(heights) - 1; i >= 0; i--{
		r = append(r, m[heights[i]])
	}
	return
}

// My ans before :

// #1 :
// func sortPeople(names []string, heights []int) (r []string) {
// 	d := map[int]bool{}
// 	for i := 0; i < len(names); i++{
// 		if d[i]{
// 			continue
// 		}
// 		push := ""
// 		stop := false
// 		for j := 0; j < len(names); j++{
// 			if heights[j] > heights[i] && !d[j] {
// 				push = names[j]
// 				d[j] = true
// 				stop = true
// 			}
// 		}
// 		if !stop {
// 			d[i] = true
// 			push = names[i]
// 		}
// 		r = append(r, push)
// 	}
// 	for k := 0; k < len(names); k++{
// 		if !d[k] {
// 			r = append(r, names[k])
// 		}
// 	}
// 	return
// }

// #2 :
// func sortPeople(names []string, heights []int) (r []string) {
// 	x := map[int]bool{}
// 	for i := 0; i < len(heights); i++{
// 		a := heights[i]
// 		for j := 0; j < len(heights); j++{
// 			if !x[j] && heights[j] > a {
// 				fmt.Println(names[j], heights[j])
// 				r = append(r, names[j])
// 				x[j] = true
// 			}
// 		}
// 	}
// 	for i := 0; i < len(heights); i++{
// 		if !x[i] {
// 			r = append(r, names[i])
// 		}
// 	}
// 	return
// }
