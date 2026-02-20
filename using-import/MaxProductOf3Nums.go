package usingimportpackage

import "sort"

// Question :

// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.
// Example 1:
// Input: nums = [1,2,3]
// Output: 6
// Example 2:
// Input: nums = [1,2,3,4]
// Output: 24
// Example 3:
// Input: nums = [-1,-2,-3]
// Output: -6
// Constraints:
// 3 <= nums.length <= 104
// -1000 <= nums[i] <= 1000

// Answer :
func MaxProductOf3Nums(nums []int) int {
	neg, pos := []int{}, []int{}
	for _, v := range nums {
		if v < 0 {
			neg = append(neg, v)
		} else {
			pos = append(pos, v)
		}
	}
	sort.Ints(neg)
	sort.Ints(pos)
	sort.Ints(nums)
	x, y := nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3], 0
	if len(pos) > 0 && len(neg) > 1 {
		y = neg[0] * neg[1] * pos[len(pos)-1]
	}
	if y == 0 {
		return x
	}
	if x > y {
		return x
	}
	return y
}

// func MaxProductOf3Nums(nums []int) (r int) {
// 	for i := 0; i < len(nums)-2; i++ {
// 		x := nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			y := nums[j]
// 			for k := j + 1; k < len(nums); k++ {
// 				z := nums[k]
// 				if i == 0 && j == 1 && k == 2 {
// 					r = x * y * z
// 				} else {
// 					if x*y*z > r {
// 						r = x * y * z
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return
// }
