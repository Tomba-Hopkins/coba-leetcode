package difficultyeasy

// Question :

// You are given an integer array nums.
// A tuple (i, j, k) of 3 distinct indices is good if nums[i] == nums[j] == nums[k].
// The distance of a good tuple is abs(i - j) + abs(j - k) + abs(k - i), where abs(x) denotes the absolute value of x.
// Return an integer denoting the minimum possible distance of a good tuple. If no good tuples exist, return -1.
// Example 1:
// Input: nums = [1,2,1,1,3]
// Output: 6
// Explanation:
// The minimum distance is achieved by the good tuple (0, 2, 3).
// (0, 2, 3) is a good tuple because nums[0] == nums[2] == nums[3] == 1. Its distance is abs(0 - 2) + abs(2 - 3) + abs(3 - 0) = 2 + 1 + 3 = 6.
// Example 2:
// Input: nums = [1,1,2,3,2,1,2]
// Output: 8
// Explanation:
// The minimum distance is achieved by the good tuple (2, 4, 6).
// (2, 4, 6) is a good tuple because nums[2] == nums[4] == nums[6] == 2. Its distance is abs(2 - 4) + abs(4 - 6) + abs(6 - 2) = 2 + 2 + 4 = 8.
// Example 3:
// Input: nums = [1]
// Output: -1
// Explanation:
// There are no good tuples. Therefore, the answer is -1.
// Constraints:
// 1 <= n == nums.length <= 100
// 1 <= nums[i] <= n

// Answer :
func MinDistBetweenThreeEqualElements1(nums []int) (r int) {
	for i := len(nums) - 1; i > 1; i-- {
		for j := i - 1; j > 0; j-- {
			for k := j - 1; k >= 0; k-- {
				x1, x2, x3 := nums[i], nums[j], nums[k]
				if x1 == x2 && x2 == x3 && x3 == x1 {
					sum := (max(i, j) - min(i, j)) + (max(j, k) - min(j, k)) + (max(k, i) - min(k, i))
					if r == 0 {
						r = sum
					} else {
						if sum < r {
							r = sum
						}
					}

				}
			}
		}
	}
	if r == 0 {
		r = -1
	}
	return
}