package difficultymedium

// Question :

// You are given an n x n integer matrix. You can do the following operation any number of times:
// Choose any two adjacent elements of matrix and multiply each of them by -1.
// Two elements are considered adjacent if and only if they share a border.
// Your goal is to maximize the summation of the matrix's elements. Return the maximum sum of the matrix's elements using the operation mentioned above.
// Example 1:
// Input: matrix = [[1,-1],[-1,1]]
// Output: 4
// Explanation: We can follow the following steps to reach sum equals 4:
// - Multiply the 2 elements in the first row by -1.
// - Multiply the 2 elements in the first column by -1.
// Example 2:
// Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
// Output: 16
// Explanation: We can follow the following step to reach sum equals 16:
// - Multiply the 2 last elements in the second row by -1.
// Constraints:
// n == matrix.length == matrix[i].length
// 2 <= n <= 250
// -105 <= matrix[i][j] <= 105

// Answer :

// gpt help me for explained the description
func MaxMatrixSum(matrix [][]int) int64 {
	sum, negative, min := 0, 0, 0
	for i, mat := range matrix {
		for j, m := range mat {
			if m < 0 {
				negative++
			}
			abs := m
			if m < 0 {
				abs = -m
			}
			sum += abs
			if i == 0 && j == 0 {
				min = abs
			}
			if abs < min {
				min = abs
			}
		}
	}
	if negative%2 == 0 {
		return int64(sum)
	}
	dist := 2 * min
	return int64(sum) - int64(dist)
}
