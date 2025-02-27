package difficultyeasy

// Question :

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true
// Example 2:
// Input: nums = [1,0,1,1], k = 1
// Output: true
// Example 3:
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false
// Constraints:
// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109
// 0 <= k <= 105

// Answer :

func ContainsDuplicate2(nums []int, k int) bool {
	m := map[int]bool{}
	n := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]] {
			m[nums[i]] = true
			n[nums[i]] = i
		} else {
			if i-n[nums[i]] <= k {
				return true
			}
			n[nums[i]] = i
		}

	}
	return false
}


// my ans before :

// func ContainsDuplicate2(nums []int, k int) bool {
//     x := false
//     for i := 0; i < len(nums) -1 ; i++{
//         for j := i + 1; j < len(nums); j++{
//             if nums[i] == nums[j] && j - i <= k {
//                 x = true
//             }
//         }
//     }
//     return x
// }