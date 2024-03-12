package difficultyeasy

// Question :

// Example 1:
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Answer :

func RemoveDuplicates(nums []int) int {

	res := 0


	for i := 1; i < len(nums); i++{
		if nums[i] != nums[res] {
			res++
			nums[res] = nums[i]
		}
	}

	return res + 1

}


// awalnya gini :
// func removeDuplicates(nums []int) int {
// 	table := map[int]int{}
// 	for _, num := range nums {
// 		table[num]++
// 	}
// 	return len(table)
// }