package usingimportpackage

import (
	"slices"
	"sort"
)

// Question :

// You are given a positive integer array nums and an integer k.
// Choose at most k elements from nums so that their sum is maximized. However, the chosen numbers must be distinct.
// Return an array containing the chosen numbers in strictly descending order.
// Example 1:
// Input: nums = [84,93,100,77,90], k = 3
// Output: [100,93,90]
// Explanation:
// The maximum sum is 283, which is attained by choosing 93, 100 and 90. We rearrange them in strictly descending order as [100, 93, 90].
// Example 2:
// Input: nums = [84,93,100,77,93], k = 3
// Output: [100,93,84]
// Explanation:
// The maximum sum is 277, which is attained by choosing 84, 93 and 100. We rearrange them in strictly descending order as [100, 93, 84]. We cannot choose 93, 100 and 93 because the chosen numbers must be distinct.
// Example 3:
// Input: nums = [1,1,1,2,2,2], k = 6
// Output: [2,1]
// Explanation:
// The maximum sum is 3, which is attained by choosing 1 and 2. We rearrange them in strictly descending order as [2, 1].
// Constraints:
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 109
// 1 <= k <= nums.length

// Answer :
func MaximizeSumOfAtMostKeDistinctEl(nums []int, k int) (r []int) {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if len(r) > 0 {
			if !slices.Contains(r, nums[i]) {
				r = append(r, nums[i])
			}
		} else {
			r = append(r, nums[i])
		}
		if len(r) == k {
			break
		}	
	}
	return
}