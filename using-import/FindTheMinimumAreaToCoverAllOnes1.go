package usingimportpackage

import "slices"

// Question :

// You are given a 2D binary array grid. Find a rectangle with horizontal and vertical sides with the smallest area, such that all the 1's in grid lie inside this rectangle.
// Return the minimum possible area of the rectangle.
// Example 1:
// Input: grid = [[0,1,0],[1,0,1]]
// Output: 6
// Explanation:
// The smallest rectangle has a height of 2 and a width of 3, so it has an area of 2 * 3 = 6.
// Example 2:
// Input: grid = [[1,0],[0,0]]
// Output: 1
// Explanation:
// The smallest rectangle has both height and width 1, so its area is 1 * 1 = 1.
// Constraints:
// 1 <= grid.length, grid[i].length <= 1000
// grid[i][j] is either 0 or 1.
// The input is generated such that there is at least one 1 in grid.

// Answer :
func FindTheMinimumAreaToCoverAllOnes1(grid [][]int) int {
	y1, y2 := 0, 0
	x1, x2 := 0, 0
	for i := 0; i < len(grid); i++{
		if y1 == 0 && slices.Contains(grid[i], 1){
			y1 = i +1
		}
		if y1 != 0 && slices.Contains(grid[i], 1){
			y2 = i + 1
		}
		for j := 0; j < len(grid[i]); j++{
			if x1 == 0 && grid[i][j] == 1 {
				x1 = j + 1
			}
			if j < x1 && grid[i][j] == 1 {
				x1 = j + 1
			}
			if j + 1  > x2 && grid[i][j] == 1 {
				x2 = j + 1
			}
		}
	}
	return (x2 - x1 + 1) * (y2 - y1 + 1)
}