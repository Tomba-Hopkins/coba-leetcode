package usingimportpackage

import "sort"

// Question :

// You are given an integer array nums and an integer k. You want to find a subsequence of nums of length k that has the largest sum.
// Return any such subsequence as an integer array of length k.
// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
// Example 1:
// Input: nums = [2,1,3,3], k = 2
// Output: [3,3]
// Explanation:
// The subsequence has the largest sum of 3 + 3 = 6.
// Example 2:
// Input: nums = [-1,-2,3,4], k = 3
// Output: [-1,3,4]
// Explanation:
// The subsequence has the largest sum of -1 + 3 + 4 = 6.
// Example 3:
// Input: nums = [3,4,3,3], k = 2
// Output: [3,4]
// Explanation:
// The subsequence has the largest sum of 3 + 4 = 7.
// Another possible subsequence is [4, 3].
// Constraints:
// 1 <= nums.length <= 1000
// -105 <= nums[i] <= 105
// 1 <= k <= nums.length

// Answer :
func FindSubseqOfLenKeyWithTheLargeSum(nums []int, k int) (r []int) {
	a := make([]int, len(nums))
	copy(a, nums)
	sort.Ints(nums)
	m, d := map[int]int{}, 0
	for i := len(nums) - 1; i >= 0; i-- {
		m[nums[i]]++
		d++
		if d == k {
			break
		}
	}
	for i := 0; i < len(a); i++{
		if m[a[i]] > 0 {
			r = append(r, a[i])
			m[a[i]]--
		}
		if len(r) == k {
			break
		}
	}
	return
}