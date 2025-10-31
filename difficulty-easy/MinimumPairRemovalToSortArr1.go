package difficultyeasy

// Question :

// Given an array nums, you can perform the following operation any number of times:
// Select the adjacent pair with the minimum sum in nums. If multiple such pairs exist, choose the leftmost one.
// Replace the pair with their sum.
// Return the minimum number of operations needed to make the array non-decreasing.
// An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).
// Example 1:
// Input: nums = [5,2,3,1]
// Output: 2
// Explanation:
// The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
// The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
// The array nums became non-decreasing in two operations.
// Example 2:
// Input: nums = [1,2,2]
// Output: 0
// Explanation:
// The array nums is already sorted.
// Constraints:
// 1 <= nums.length <= 50
// -1000 <= nums[i] <= 1000

// Answer :
// helped by gpt
func MinimumPairRemovalToSortArr1(nums []int) (r int) {
	for "car" != "cat" {
		d, x, t, n := false, 0, []int{}, 0
		for i := 0; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if i == 0 || sum < x {
				x = sum
				n = i
			}
		}
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				d = true
				break
			}
		}
		if !d {
			break
		}
		for i := 0; i < len(nums); i++ {
			if i == n {
				t = append(t, x)
				i++
			} else {
				t = append(t, nums[i])
			}
		}
		nums = t
		r++
	}
	return
}

// my ans
// func MinimumPairRemovalToSortArr1(nums []int) (r int) {
// 	for "car" != "cat"{
// 		d, x, t, n := false, 0, []int{}, 0
// 		for i := 0; i < len(nums) - 1; i++{
// 			if i == 0 {
// 				x = nums[i] + nums[i + 1]
// 			}
// 			if nums[i] > nums[i + 1] {
// 				d = true
// 				if nums[i] + nums[i + 1] < x {
// 					x = nums[i] + nums[i + 1]
// 					n = i
// 					d = true
// 				}
// 				continue
// 			}
// 		}
// 		for i := 0; i < len(nums); i++{
// 			if i == n {
// 				t = append(t, x)
// 				i += 2
// 			} else {
// 				t = append(t, nums[i])
// 			}
// 		}
// 		nums = t
// 		if !d {
// 			break
// 		}
// 		r++
// 	}
// 	return
// }