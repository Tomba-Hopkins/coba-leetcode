package usingimportpackage

import "sort"

// Question :

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
// Example 1:
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
// Example 3:
// Input: nums = [1,0,1,2]
// Output: 3
// Constraints:
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

// Answer :
func LongestConsecutiveSeq(nums []int) (r int) {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			continue
		} else if nums[i] == nums[i - 1] + 1 {
			c++
		} else {
			if c > r { 
				r = c
			}
			c = 1
		}
	}
	if c > r {
		r = c 
	}
	return
}

// my wrong ans before :
// func LongestConsecutiveSeq(nums []int) (r int) {
// 	sort.Ints(nums)
// 	b := []int{}
// 	c := 0
// 	d := 0
// 	z := 0
// 	for i := 0; i < len(nums); i++{
// 		if i == 0 {
// 			d = nums[i]
// 			c++
// 			continue
// 		} 
// 		if nums[i] == d && nums[i] != z  {
// 			continue
// 		}
// 		if nums[i] ==  d + 1{
// 			d++
// 			c++
// 		} else {
// 			b  = append(b, c)
// 			c = 1
// 			d = nums[i]
// 			z = nums[i]
// 		}
// 		if i == len(nums) - 1 {
// 			b  = append(b, c)
// 		}
// 	}
// 	for _, x := range b {
// 		if x > r {
// 			r = x
// 		}
// 	}
// 	return
// }
