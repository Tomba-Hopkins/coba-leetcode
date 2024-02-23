package difficultyeasy

// Question :
// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Answer :

func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++{ // ini iterasi lagi kalau j udah mentok ampe nums.length
		for j := i + 1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				// result = append(result,nums[i], nums[j])
				result = append(result, i, j)
			}
		}
	}
	return result
}
