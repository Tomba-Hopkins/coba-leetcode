package usingimportpackage

import "sort"

// Question :

// You are given a 0-indexed integer array nums. Rearrange the values of nums according to the following rules:
// Sort the values at odd indices of nums in non-increasing order.
// For example, if nums = [4,1,2,3] before this step, it becomes [4,3,2,1] after. The values at odd indices 1 and 3 are sorted in non-increasing order.
// Sort the values at even indices of nums in non-decreasing order.
// For example, if nums = [4,1,2,3] before this step, it becomes [2,1,4,3] after. The values at even indices 0 and 2 are sorted in non-decreasing order.
// Return the array formed after rearranging the values of nums.
// Example 1:
// Input: nums = [4,1,2,3]
// Output: [2,3,4,1]
// Explanation:
// First, we sort the values present at odd indices (1 and 3) in non-increasing order.
// So, nums changes from [4,1,2,3] to [4,3,2,1].
// Next, we sort the values present at even indices (0 and 2) in non-decreasing order.
// So, nums changes from [4,1,2,3] to [2,3,4,1].
// Thus, the array formed after rearranging the values is [2,3,4,1].
// Example 2:
// Input: nums = [2,1]
// Output: [2,1]
// Explanation:
// Since there is exactly one odd index and one even index, no rearrangement of values takes place.
// The resultant array formed is [2,1], which is the same as the initial array.
// Constraints:
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100

// Answer :
func SortEvenAndOddIndicesIndependently(nums []int) (r []int) {
	o, e, d1, d2 := []int{}, []int{}, 0, 0
	for i, n := range nums {
		if i % 2 == 0 {
			e = append(e, n)
		} else {
			o = append(o, n)
		}
	}
	sort.Slice(o, func(i, j int) bool {
		return o[i] > o[j]
	})
	sort.Slice(e, func(i, j int) bool {
		return e[i] < e[j]
	})
	for len(r) < len(nums){
		if d1 < len(e){
			r = append(r, e[d1])
		}
		if d2 < len(o){
			r = append(r, o[d2])
		}
		d1++
		d2++
	}
	return
}