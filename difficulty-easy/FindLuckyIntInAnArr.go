package difficultyeasy

// Question :

// Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.
// Return the largest lucky integer in the array. If there is no lucky integer return -1.
// Example 1:
// Input: arr = [2,2,3,4]
// Output: 2
// Explanation: The only lucky number in the array is 2 because frequency[2] == 2.
// Example 2:
// Input: arr = [1,2,2,3,3,3]
// Output: 3
// Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
// Example 3:
// Input: arr = [2,2,2,3,3]
// Output: -1
// Explanation: There are no lucky numbers in the array.
// Constraints:
// 1 <= arr.length <= 500
// 1 <= arr[i] <= 500

// Answer :
func FindLuckyIntInAnArr(arr []int) int {
	m := map[int]int{}
	x := 0
	for _, v := range arr {
		m[v]++
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == m[arr[i]] {
			if x < arr[i] {
				x = arr[i]
			}
		}
	}
	if x > 0 {
		return x
	}
	return -1
}