package difficultyeasy

// Question :

// Given an array of integers arr, and three integers a, b and c. You need to find the number of good triplets.
// A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:
// 0 <= i < j < k < arr.length
// |arr[i] - arr[j]| <= a
// |arr[j] - arr[k]| <= b
// |arr[i] - arr[k]| <= c
// Where |x| denotes the absolute value of x.
// Return the number of good triplets.
// Example 1:
// Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
// Output: 4
// Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].
// Example 2:
// Input: arr = [1,1,2,2,3], a = 0, b = 0, c = 1
// Output: 0
// Explanation: No triplet satisfies all conditions.
// Constraints:
// 3 <= arr.length <= 100
// 0 <= arr[i] <= 1000
// 0 <= a, b, c <= 1000

// Answer :
func CountGoodTriplets(arr []int, a int, b int, c int) (r int) {
	for i := 0; i < len(arr)-2; i++ {
		x, y, z := false, false, false
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				if arr[i]-arr[j] <= a {
					x = true
				} else {
					x = false
				}
			} else {
				if arr[j]-arr[i] <= a {
					x = true
				} else {
					x = false
				}
			}
			for k := j + 1; k < len(arr); k++ {
				if arr[j] > arr[k] {
					if arr[j]-arr[k] <= b {
						y = true
					} else {
						y = false
					}
				} else {
					if arr[k]-arr[j] <= b {
						y = true
					} else {
						y = false
					}
				}
				if arr[i] > arr[k] {
					if arr[i]-arr[k] <= c {
						z = true
					} else {
						z = false
					}
				} else {
					if arr[k]-arr[i] <= c {
						z = true
					} else {
						z = false
					}
				}
				if x && y && z {
					r++
					z = false
				}
			}
		}
	}
	return
}