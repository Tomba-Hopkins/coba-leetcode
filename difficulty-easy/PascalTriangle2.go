package difficultyeasy

// Question :

// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
// Example 1:
// Input: rowIndex = 3
// Output: [1,3,3,1]
// Example 2:
// Input: rowIndex = 0
// Output: [1]
// Example 3:
// Input: rowIndex = 1
// Output: [1,1]
// Constraints:
// 0 <= rowIndex <= 33
// Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?

// Answer :
func PascalTriangle2(rowIndex int) []int {
	t := [][]int{}
	for i := 1; i <= rowIndex + 1; i++{
		a := []int{}
		for j := 0; j < i; j++{
			if j == 0 || j == i - 1 {
				a = append(a, 1)
			} else {
				a = append(a, t[len(t) - 1][j] + t[len(t) - 1][j - 1])
			}
		}
		t = append(t, a)
	} 
	return t[rowIndex]
}