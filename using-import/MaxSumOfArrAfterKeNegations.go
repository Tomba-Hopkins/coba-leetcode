package usingimportpackage

import "sort"

// Question :

// Given an integer array nums and an integer k, modify the array in the following way:
// choose an index i and replace nums[i] with -nums[i].
// You should apply this process exactly k times. You may choose the same index i multiple times.
// Return the largest possible sum of the array after modifying it in this way.
// Example 1:
// Input: nums = [4,2,3], k = 1
// Output: 5
// Explanation: Choose index 1 and nums becomes [4,-2,3].
// Example 2:
// Input: nums = [3,-1,0,2], k = 3
// Output: 6
// Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
// Example 3:
// Input: nums = [2,-3,-1,5,-4], k = 2
// Output: 13
// Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
// Constraints:
// 1 <= nums.length <= 104
// -100 <= nums[i] <= 100
// 1 <= k <= 104

// Answer :

// gpt ans :
func MaxSumOfArrAfterKeNegations(nums []int, k int) (r int) {
    sort.Ints(nums)
    for ; k > 0; k-- {
        nums[0] = -nums[0]
        sort.Ints(nums)
    }
    for _, n := range nums {
        r += n
    }
    return
}

// my wrong answer :

// func MaxSumOfArrAfterKeNegations(nums []int, k int) (r int) {
// 	sort.Ints(nums)
// 	for k > 0 {
// 		for i := 0; i < len(nums); i++{
// 			x := nums[i]
// 			if k == 0 {
// 				break
// 			}
// 			if x < 0 && k > 0 {
// 				x = -x
// 				k--
// 			} else if x >= 0 && k > 0 {
// 				for k > 0 {
// 					x = -x
// 					k--
// 				}
// 			}
// 			nums[i] = x
// 		}
// 	}
// 	for _, n := range nums {
// 		r += n
// 	}
// 	return
// }