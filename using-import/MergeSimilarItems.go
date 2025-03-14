package usingimportpackage

import "sort"

// Question :

// You are given two 2D integer arrays, items1 and items2, representing two sets of items. Each array items has the following properties:
// items[i] = [valuei, weighti] where valuei represents the value and weighti represents the weight of the ith item.
// The value of each item in items is unique.
// Return a 2D integer array ret where ret[i] = [valuei, weighti], with weighti being the sum of weights of all items with value valuei.
// Note: ret should be returned in ascending order by value.
// Example 1:
// Input: items1 = [[1,1],[4,5],[3,8]], items2 = [[3,1],[1,5]]
// Output: [[1,6],[3,9],[4,5]]
// Explanation:
// The item with value = 1 occurs in items1 with weight = 1 and in items2 with weight = 5, total weight = 1 + 5 = 6.
// The item with value = 3 occurs in items1 with weight = 8 and in items2 with weight = 1, total weight = 8 + 1 = 9.
// The item with value = 4 occurs in items1 with weight = 5, total weight = 5.
// Therefore, we return [[1,6],[3,9],[4,5]].
// Example 2:
// Input: items1 = [[1,1],[3,2],[2,3]], items2 = [[2,1],[3,2],[1,3]]
// Output: [[1,4],[2,4],[3,4]]
// Explanation:
// The item with value = 1 occurs in items1 with weight = 1 and in items2 with weight = 3, total weight = 1 + 3 = 4.
// The item with value = 2 occurs in items1 with weight = 3 and in items2 with weight = 1, total weight = 3 + 1 = 4.
// The item with value = 3 occurs in items1 with weight = 2 and in items2 with weight = 2, total weight = 2 + 2 = 4.
// Therefore, we return [[1,4],[2,4],[3,4]].
// Example 3:
// Input: items1 = [[1,3],[2,2]], items2 = [[7,1],[2,2],[1,4]]
// Output: [[1,7],[2,4],[7,1]]
// Explanation:
// The item with value = 1 occurs in items1 with weight = 3 and in items2 with weight = 4, total weight = 3 + 4 = 7.
// The item with value = 2 occurs in items1 with weight = 2 and in items2 with weight = 2, total weight = 2 + 2 = 4.
// The item with value = 7 occurs in items2 with weight = 1, total weight = 1.
// Therefore, we return [[1,7],[2,4],[7,1]].
// Constraints:
// 1 <= items1.length, items2.length <= 1000
// items1[i].length == items2[i].length == 2
// 1 <= valuei, weighti <= 1000
// Each valuei in items1 is unique.
// Each valuei in items2 is unique.

// Answer :
func MergeSimilarItems(items1 [][]int, items2 [][]int) (r [][]int) {
	for _, i1 := range items1 {
		t := false
		for _, j := range items2 {
			if i1[0] == j[0]{
				t = true
			}
		}
		if !t {
			r = append(r, i1)
		}	
	}
	for _, i2 := range items2 {
		t := false
		for _, j := range items1 {
			if i2[0] == j[0]{
				t = true
			}
		}
		if !t {
			r = append(r, i2)
		}	
	}
	for i := 0; i < len(items1); i++{
		d := items1[i][0]
		for j := 0; j < len(items2); j++{
			if d == items2[j][0]{
				x := items1[i][1] + items2[j][1]
				r = append(r, []int{d,x})
				break
			}
		}
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i][0] < r[j][0]
	})
	return
}


// alternative
// func MergeSimilarItems(items1 [][]int, items2 [][]int) (r [][]int) {
// 	m := map[int]int{}
// 	for _, n := range items1 {
// 		m[n[0]] += n[1]
// 	}
// 	for _, n := range items2 {
// 		m[n[0]] += n[1]
// 	}
// 	for key, val := range m {
// 		r = append(r, []int{key,val})
// 	}
// 	sort.Slice(r, func(i, j int) bool {
// 		return r[i][0] < r[j][0]
// 	})
// 	return
// }