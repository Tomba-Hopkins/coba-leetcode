package difficultyeasy

// Question :

// You are given an integer array nums of length n.
// An array is trionic if there exist indices 0 < p < q < n − 1 such that:
// nums[0...p] is strictly increasing,
// nums[p...q] is strictly decreasing,
// nums[q...n − 1] is strictly increasing.
// Return true if nums is trionic, otherwise return false.
// Example 1:
// Input: nums = [1,3,5,4,2,6]
// Output: true
// Explanation:
// Pick p = 2, q = 4:
// nums[0...2] = [1, 3, 5] is strictly increasing (1 < 3 < 5).
// nums[2...4] = [5, 4, 2] is strictly decreasing (5 > 4 > 2).
// nums[4...5] = [2, 6] is strictly increasing (2 < 6).
// Example 2:
// Input: nums = [2,1,3]
// Output: false
// Explanation:
// There is no way to pick p and q to form the required three segments.
// Constraints:
// 3 <= n <= 100
// -1000 <= nums[i] <= 1000

// Answer :
func TrionicArr1(nums []int) bool {
	a, newA := []bool{}, []bool{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			a = append(a, true)
		} else if nums[i] > nums[i+1] {
			a = append(a, false)
		} else {
			return false
		}
	}
	for i := 0; i < len(a); i++ {
		if i == 0 {
			newA = append(newA, a[i])
		} else {
			if newA[len(newA)-1] != a[i] {
				newA = append(newA, a[i])
			}
		}
	}
	if len(newA) != 3 {
		return false
	}
	return newA[0] && !newA[1] && newA[2]
}
