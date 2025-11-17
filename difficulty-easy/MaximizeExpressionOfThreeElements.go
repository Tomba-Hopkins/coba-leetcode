package difficultyeasy

// Question :

// You are given an integer array nums.
// Choose three elements a, b, and c from nums at distinct indices such that the value of the expression a + b - c is maximized.
// Return an integer denoting the maximum possible value of this expression.
// Example 1:
// Input: nums = [1,4,2,5]
// Output: 8
// Explanation:
// We can choose a = 4, b = 5, and c = 1. The expression value is 4 + 5 - 1 = 8, which is the maximum possible.
// Example 2:
// Input: nums = [-2,0,5,-2,4]
// Output: 11
// Explanation:
// We can choose a = 5, b = 4, and c = -2. The expression value is 5 + 4 - (-2) = 11, which is the maximum possible.
// Constraints:
// 3 <= nums.length <= 100
// -100 <= nums[i] <= 100

// Answer :
func MaximizeExpressionOfThreeElements(nums []int) int {
	h, h2, l, iH := 0, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			h = nums[i]
			l = nums[i]
		}
		if nums[i] > h {
			h = nums[i]
			iH = i
		}
		if nums[i] < l {
			l = nums[i]
		}
	}
	d := false
	for i := 0; i < len(nums); i++ {
		if iH != i && !d {
			h2 = nums[i]
			d = true
		}
		if nums[i] > h2 && i != iH {
			h2 = nums[i]
		}
	}
	return h + h2 - l
}

// with package
// func MaximizeExpressionOfThreeElements(nums []int) int {
// 	sort.Ints(nums)
// 	return  nums[len(nums) - 1] + nums[len(nums) - 2] - nums[0]
// }