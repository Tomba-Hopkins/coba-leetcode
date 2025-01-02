package difficultyeasy

// Question :

// Given an m x n matrix of distinct numbers, return all lucky numbers in the matrix in any order.
// A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.
// Example 1:
// Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
// Output: [15]
// Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column.
// Example 2:
// Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
// Output: [12]
// Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.
// Example 3:
// Input: matrix = [[7,8],[1,2]]
// Output: [7]
// Explanation: 7 is the only lucky number since it is the minimum in its row and the maximum in its column.
// Constraints:
// m == mat.length
// n == mat[i].length
// 1 <= n, m <= 50
// 1 <= matrix[i][j] <= 105.
// All elements in the matrix are distinct.

// Answer :
func LuckyNumbersInMatrix(matrix [][]int) (r []int) {
	x := []int{}
	y := []int{}
	for i := 0; i < len(matrix); i++ {
		t := matrix[i][0]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < t {
				t = matrix[i][j]
			}
		}
		x = append(x, t)
	}
	for i := 0; i < len(matrix[0]); i++ {
		t := matrix[0][i]

		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > t {
				t = matrix[j][i]
			}
		}
		y = append(y, t)
	}
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				r = append(r, x[i])
			}
		}
	}
	return
}