package difficultyeasy

// Question :

// You are given an integer array nums and an integer k.
// An integer h is called valid if all values in the array that are strictly greater than h are identical.
// For example, if nums = [10, 8, 10, 8], a valid integer is h = 9 because all nums[i] > 9 are equal to 10, but 5 is not a valid integer.
// You are allowed to perform the following operation on nums:
// Select an integer h that is valid for the current values in nums.
// For each index i where nums[i] > h, set nums[i] to h.
// Return the minimum number of operations required to make every element in nums equal to k. If it is impossible to make all elements equal to k, return -1.
// Example 1:
// Input: nums = [5,2,5,4,5], k = 2
// Output: 2
// Explanation:
// The operations can be performed in order using valid integers 4 and then 2.
// Example 2:
// Input: nums = [2,1,2], k = 2
// Output: -1
// Explanation:
// It is impossible to make all the values equal to 2.
// Example 3:
// Input: nums = [9,7,5,3], k = 1
// Output: 4
// Explanation:
// The operations can be performed using valid integers in the order 7, 5, 3, and 1.
// Constraints:
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
// 1 <= k <= 100

// Answer :
func MinOperationsToMakeArrValEqualToK(nums []int, k int) int {
	same := true
	for _, n := range nums {
		if n != k {
			same = false
		}
	}
	if same {
		return 0
	}
	for _, n := range nums {
		if n < k {
			return -1
		}
	}
	m := map[int]bool{}
	a := []int{}
	for _, n := range nums {
		if !m[n] {
			a = append(a, n)
			m[n] = true
		}
	}
	r := 0
	for _, x := range a {
		if x > k {
			r++
		}
	}
	return r
}


// gpt ans :
// func MinOperationsToMakeArrValEqualToK(nums []int, k int) int {
// 	same := true
// 	for _, n := range nums {
// 		if n != k {
// 			same = false
// 		}
// 	}
// 	if same {
// 		return 0
// 	}
// 	for _, n := range nums {
// 		if n < k {
// 			return -1
// 		}
// 	}
// 	m := map[int]bool{}
// 	r := 0
// 	for _, x := range nums {
// 		if x > k && !m[x] {
// 			m[x] = true
// 			r++
// 		}
// 	}
// 	return r
// }



// my first ans :
// func MinOperationsToMakeArrValEqualToK(nums []int, k int) int {
// 	r := 0
// 	m := map[int]bool{}
// 	a := []int{}
// 	for _, n := range nums {
// 		if !m[n] {
// 			a = append(a, n)
// 			m[n] = true
// 		}
// 	}
// 	for i := 0; i < len(a); i++{
// 		if a[i] > k {
// 			r++
// 		}
// 	}
// 	if len(m) == 1 {
// 		return 0
// 	}
// 	if r == 0 {
// 		return -1
// 	}
// 	return r
// }
