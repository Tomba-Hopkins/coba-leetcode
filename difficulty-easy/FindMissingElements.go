package difficultyeasy

// Question :

// You are given an integer array nums consisting of unique integers.
// Originally, nums contained every integer within a certain range. However, some integers might have gone missing from the array.
// The smallest and largest integers of the original range are still present in nums.
// Return a sorted list of all the missing integers in this range. If no integers are missing, return an empty list.
// Example 1:
// Input: nums = [1,4,2,5]
// Output: [3]
// Explanation:
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. Among these, only 3 is missing.
// Example 2:
// Input: nums = [7,8,6,9]
// Output: []
// Explanation:
// The smallest integer is 6 and the largest is 9, so the full range is [6,7,8,9]. All integers are already present, so no integer is missing.
// Example 3:
// Input: nums = [5,1]
// Output: [2,3,4]
// Explanation:
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. The missing integers are 2, 3, and 4.
// Constraints:
// 2 <= nums.length <= 100
// 1 <= nums[i] <= 100

// Answer :
func FindMissingElements(nums []int) (r []int) {
	h, l := 0, 0
	for i := 0; i < len(nums); i++{
		if i == 0 {
			h = nums[i]
			l = nums[i]
		}
		if nums[i] > h {
			h = nums[i]
		}
		if nums[i] < l {
			l = nums[i]
		}
	}
	for i := l; i < h; i++{
		d := false
		for _, v := range nums {
			if v == i {
				d = true
				break
			}
		}
		if !d {
			r = append(r, i)
		}
	}
	return
}

// with package
// func FindMissingElements(nums []int) (r []int) {
// 	h, l := slices.Max(nums), slices.Min(nums)
// 	for i := l; i < h; i++{
// 		if !slices.Contains(nums, i){
// 			r = append(r, i)
// 		}
// 	}
// 	return
// }