package difficultyeasy

// Question :

// Example 1:
// Input: nums = [-1,2,-3,3]
// Output: 3
// Explanation: 3 is the only valid k we can find in the array.
// Example 2:
// Input: nums = [-1,10,6,7,-7,1]
// Output: 7
// Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.
// Example 3:
// Input: nums = [-10,8,6,7,-2,-3]
// Output: -1
// Explanation: There is no a single valid k, we return -1.

// Answer :

func FindMaxK(nums []int) int {
	max := -1
 
	 isi := map[int]int{}
	 for _, i := range nums{
		 isi[i]++
	 }
 
	 for _, n := range nums{
		 if  isi[-n] > 0 && n > max{
			 max = n
		 }
	 }
 
	 return max
 }