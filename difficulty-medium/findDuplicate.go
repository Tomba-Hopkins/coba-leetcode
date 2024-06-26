package difficultymedium

// Question :

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.
// Example 1:
// Input: nums = [1,3,4,2,2]
// Output: 2
// Example 2:
// Input: nums = [3,1,3,4,2]
// Output: 3
// Example 3:
// Input: nums = [3,3,3,3,3]
// Output: 3

// Answer :
func FindDuplicate(nums []int) int {
	box := map[int]int{}

	for _, n := range nums {
		box[n]++
	}

	res := 0

	for _, n := range nums {
		if box[n] > 1 {
			res = n
		}
	}

	return res

}