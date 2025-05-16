package difficultyeasy

// Question :

// You are given an array nums, where each number in the array appears either once or twice.
// Return the bitwise XOR of all the numbers that appear twice in the array, or 0 if no number appears twice.
// Example 1:
// Input: nums = [1,2,1,3]
// Output: 1
// Explanation:
// The only number that appears twice in nums is 1.
// Example 2:
// Input: nums = [1,2,3]
// Output: 0
// Explanation:
// No number appears twice in nums.
// Example 3:
// Input: nums = [1,2,2,1]
// Output: 3
// Explanation:
// Numbers 1 and 2 appeared twice. 1 XOR 2 == 3.
// Constraints:
// 1 <= nums.length <= 50
// 1 <= nums[i] <= 50
// Each number in nums appears either once or twice.

// Answer :
func FindTheXorNumbersWhichAppearTwice(nums []int) int {
	m := map[int]int{}
	t := []int{}
	for _, n := range nums {
		m[n]++
	}
	for _, n := range nums {
		if m[n] > 1 {
			t = append(t, n)
			m[n] = 0
		}
	}
	r := 0
	for i := 0; i < len(t); i++ {
		if i == 0 {
			r = t[i]
			continue
		}
		r ^= t[i]
	}
	return r
}