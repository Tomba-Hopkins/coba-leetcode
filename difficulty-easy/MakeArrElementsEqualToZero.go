package difficultyeasy

// Question :

// You are given an integer array nums.
// Start by selecting a starting position curr such that nums[curr] == 0, and choose a movement direction of either left or right.
// After that, you repeat the following process:
// If curr is out of the range [0, n - 1], this process ends.
// If nums[curr] == 0, move in the current direction by incrementing curr if you are moving right, or decrementing curr if you are moving left.
// Else if nums[curr] > 0:
// Decrement nums[curr] by 1.
// Reverse your movement direction (left becomes right and vice versa).
// Take a step in your new direction.
// A selection of the initial position curr and movement direction is considered valid if every element in nums becomes 0 by the end of the process.
// Return the number of possible valid selections.
// Example 1:
// Input: nums = [1,0,2,0,3]
// Output: 2
// Explanation:
// The only possible valid selections are the following:
// Choose curr = 3, and a movement direction to the left.
// [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,0].
// Choose curr = 3, and a movement direction to the right.
// [1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2] -> [1,0,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0].
// Example 2:
// Input: nums = [2,3,4,0,4,1,0]
// Output: 0
// Explanation:
// There are no possible valid selections.
// Constraints:
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
// There is at least one element i where nums[i] == 0.

// Answer :
func MakeArrElementsEqualToZero(nums []int) (r int) {
	for i := 0; i < len(nums); i++{
	    b := false
		if nums[i] != 0 {
			continue
		}
		dst1, dst2 := make([]int, len(nums)), make([]int, len(nums))
		copy(dst1, nums)
		copy(dst2, nums)
		d := i
		for "car" != "cat" {
			if d == len(nums) || d == -1 {
				break
			}
			if dst1[d] == 0 {
				if !b {
					d++
				} else {
					d--
				}
				continue
			} else {
				dst1[d]--
				b = !b
			}
			if !b {
				d++
			} else {
				d--
			}
		}
		b = true
		d = i
		for "car" != "cat" {
			if d == len(nums) || d == -1 {
				break
			}
			if dst2[d] == 0 {
				if !b {
					d++
				} else {
					d--
				}
				continue
			} else {
				dst2[d]--
				b = !b
			}
			if !b {
				d++
			} else {
				d--
			}
		}
		x, x2 := true, true
		for _, v := range dst1 {
			if v != 0 {
				x = false
				break
			}
		}
		for _, v := range dst2 {
			if v != 0 {
				x2 = false
				break
			}
		}
		if x {
			r++
		}
		if x2 {
			r++
		}
	}
	return
}

