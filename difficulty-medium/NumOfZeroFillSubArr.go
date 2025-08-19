package difficultymedium

// Question :

// Given an integer array nums, return the number of subarrays filled with 0.
// A subarray is a contiguous non-empty sequence of elements within an array.
// Example 1:
// Input: nums = [1,3,0,0,2,0,0,4]
// Output: 6
// Explanation:
// There are 4 occurrences of [0] as a subarray.
// There are 2 occurrences of [0,0] as a subarray.
// There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.
// Example 2:
// Input: nums = [0,0,0,2,0,0]
// Output: 9
// Explanation:
// There are 5 occurrences of [0] as a subarray.
// There are 3 occurrences of [0,0] as a subarray.
// There is 1 occurrence of [0,0,0] as a subarray.
// There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.
// Example 3:
// Input: nums = [2,10,2019]
// Output: 0
// Explanation: There is no subarray filled with 0. Therefore, we return 0.
// Constraints:
// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109

// Answer :
// gpt answer :
func NumOfZeroFillSubArr(nums []int) (r int64) {
	temp := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == 0 {
			temp++
			r += int64(temp)
		} else {
			temp = 0
		}
	}
	return 
}

// my answer :
// func NumOfZeroFillSubArr(nums []int) (r int64) {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums) - i; j++{
// 			a, t := nums[j: j + i + 1], 0
// 			for _, v := range a {
// 				if v != 0 {
// 					t = 1
// 					break
// 				}
// 			}
// 			if t == 0 {
// 				r++
// 			}	
// 		}
// 	}
// 	return
// }