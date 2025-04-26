package difficultyeasy

// Question :

// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
// Return any array that satisfies this condition.
// Example 1:
// Input: nums = [3,1,2,4]
// Output: [2,4,3,1]
// Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
// Example 2:
// Input: nums = [0]
// Output: [0]
// Constraints:
// 1 <= nums.length <= 5000
// 0 <= nums[i] <= 5000

// Answer :
func SortArrayByParity(nums []int) (r []int) {
	m := map[int]bool{}
	for i, n := range nums {
		if n % 2 == 0 && !m[i] {
			r = append(r, n)
			m[i] = true
		}
	}
	for i, n := range nums {
		if n % 2 != 0 && !m[i] {
			r = append(r, n)
			m[i] = true
		}
	}
	return
}