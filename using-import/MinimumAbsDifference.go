package usingimportpackage

import "sort"

// Question :

// Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.
// Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows
// a, b are from arr
// a < b
// b - a equals to the minimum absolute difference of any two elements in arr
// Example 1:
// Input: arr = [4,2,1,3]
// Output: [[1,2],[2,3],[3,4]]
// Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.
// Example 2:
// Input: arr = [1,3,6,10,15]
// Output: [[1,3]]
// Example 3:
// Input: arr = [3,8,-10,23,19,-4,-14,27]
// Output: [[-14,-10],[19,23],[23,27]]
// Constraints:
// 2 <= arr.length <= 105
// -106 <= arr[i] <= 106

// Answer :
func MinimumAbsDifference(arr []int) (r [][]int) {
	sort.Ints(arr)
	min := 0
	for i := 0; i < len(arr) - 1; i++{
		if i == 0 {
			min = arr[i + 1] - arr[i]
		}
		if arr[i + 1] - arr[i] < min {
			min = arr[i + 1] - arr[i]
		}
	}
	for i := 0; i < len(arr) - 1; i++{
		t := []int{}
		if arr[i + 1] - arr[i] == min {
			t = append(t, arr[i])
			t = append(t, arr[i + 1])
		}
		if len(t) > 0 {
			r = append(r, t)
		}
	}
	return
}