package difficultyeasy

// Question :

// Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.
// Example 1:
// Input: arr = [1,2,2,6,6,6,6,7,10]
// Output: 6
// Example 2:
// Input: arr = [1,1]
// Output: 1
// Constraints:
// 1 <= arr.length <= 104
// 0 <= arr[i] <= 105

// Answer :
func ElAppearMoreThan25PercInSortedArr(arr []int) int {
	m := map[int]int{}
	x, r := 0, 0
	for _, a := range arr {
		m[a]++
		if m[a] > r {
			r = m[a]
			x = a
		}
	}
	return x
}