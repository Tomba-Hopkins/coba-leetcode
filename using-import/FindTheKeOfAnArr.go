package usingimportpackage

import "strconv"

// Question :
// You are given an integer array nums, and an integer k. Let's introduce K-or operation by extending the standard bitwise OR. In K-or, a bit position in the result is set to 1 if at least k numbers in nums have a 1 in that position.
// Return the K-or of nums.
// Example 1:
// Input: nums = [7,12,9,8,9,15], k = 4
// Output: 9
// Explanation:
// Represent numbers in binary:
// Number	Bit 3	Bit 2	Bit 1	Bit 0
// 7	0	1	1	1
// 12	1	1	0	0
// 9	1	0	0	1
// 8	1	0	0	0
// 9	1	0	0	1
// 15	1	1	1	1
// Result = 9	1	0	0	1
// Bit 0 is set in 7, 9, 9, and 15. Bit 3 is set in 12, 9, 8, 9, and 15.
// Only bits 0 and 3 qualify. The result is (1001)2 = 9.
// Example 2:
// Input: nums = [2,12,1,11,4,5], k = 6
// Output: 0
// Explanation: No bit appears as 1 in all six array numbers, as required for K-or with k = 6. Thus, the result is 0.
// Example 3:
// Input: nums = [10,8,5,9,11,6,8], k = 1
// Output: 15
// Explanation: Since k == 1, the 1-or of the array is equal to the bitwise OR of all its elements. Hence, the answer is 10 OR 8 OR 5 OR 9 OR 11 OR 6 OR 8 = 15.
// Constraints:
// 1 <= nums.length <= 50
// 0 <= nums[i] < 231
// 1 <= k <= nums.length

// Answer :
// gpt fixed
func FindTheKeOfAnArr(nums []int, k int) int {
	m := map[int]int{}
	maxBits := 0
	for _, num := range nums {
		bitLen := len(strconv.FormatInt(int64(num), 2))
		if bitLen > maxBits {
			maxBits = bitLen
		}
	}
	for i := 0; i < len(nums); i++ {
		s := strconv.FormatInt(int64(nums[i]), 2)
		for len(s) < maxBits {
			s = "0" + s
		}
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				m[len(s)-1-j]++
			}
		}
	}
	t := ""
	for i := maxBits - 1; i >= 0; i-- {
		if m[i] >= k {
			t += "1"
		} else {
			t += "0"
		}
	}
	n, _ := strconv.ParseInt(t, 2, 64)
	return int(n)
}

// my ans
// func FindTheKeOfAnArr(nums []int, k int) int {
// 	m := map[int]int{}
// 	for i := 0; i < len(nums); i++{
// 		s := strconv.FormatInt(int64(nums[i]), 2)
// 		for len(s) < k {
// 			s = "0" + s
// 		}
// 		for j := 0; j < len(s); j++{
// 			if s[j] == '1' {
// 				m[len(s) - 1 - j]++
// 			}
// 		}
// 	}
// 	t := ""
// 	for i := k - 1; i >= 0; i-- {
// 		if m[i] >= k {
// 			t += "1"
// 		} else {
// 			t += "0"
// 		}
// 	}
// 	n, _ := strconv.ParseInt(t, 2, 64)
// 	return int(n)
// }