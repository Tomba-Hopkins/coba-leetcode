package usingimportpackage

import (
	"slices"
	"sort"
	"strconv"
)

// Question :

// You are given an integer array nums. You have to find the maximum sum of a pair of numbers from nums such that the largest digit in both numbers is equal.
// For example, 2373 is made up of three distinct digits: 2, 3, and 7, where 7 is the largest among them.
// Return the maximum sum or -1 if no such pair exists.
// Example 1:
// Input: nums = [112,131,411]
// Output: -1
// Explanation:
// Each numbers largest digit in order is [2,3,4].
// Example 2:
// Input: nums = [2536,1613,3366,162]
// Output: 5902
// Explanation:
// All the numbers have 6 as their largest digit, so the answer is 2536 + 3366 = 5902.
// Example 3:
// Input: nums = [51,71,17,24,42]
// Output: 88
// Explanation:
// Each number's largest digit in order is [5,7,7,4,4].
// So we have only two possible pairs, 71 + 17 = 88 and 24 + 42 = 66.
// Constraints:
// 2 <= nums.length <= 100
// 1 <= nums[i] <= 104

// Answer :
func MaxPairSumInAnArr(nums []int) int {
	d, m := []int{}, map[rune]bool{}
	for i := 0; i < len(nums) - 1; i++{
		a, x, numStr := []int{nums[i]}, '0', strconv.Itoa(nums[i])
		for _, n := range numStr {
			if n > x {
				x = n
			}
		}
		if m[x] {
			continue
		}
		for j := i + 1; j < len(nums); j++{
			x2, numStr2 := '0', strconv.Itoa(nums[j])
			for _, n := range numStr2 {
				if n > x2 {
					x2 = n
				}
			}
			if x2 == x {
				a = append(a, nums[j])
			}
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})
		if len(a) > 1 {
			m[x] = true
			t := 0
			for k := 0; k < 2; k++ {
				t += a[k]
			}
			d = append(d, t)
		}
	}
	if len(d) > 0{
		return slices.Max(d)
	}
	return -1
}