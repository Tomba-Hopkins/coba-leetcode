package usingimportpackage

import "sort"

// Question :

// Given n points on a 2D plane where points[i] = [xi, yi], Return the widest vertical area between two points such that no points are inside the area.
// A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height). The widest vertical area is the one with the maximum width.
// Note that points on the edge of a vertical area are not considered included in the area.
// Example 1:
// Input: points = [[8,7],[9,9],[7,4],[9,7]]
// Output: 1
// Explanation: Both the red and the blue area are optimal.
// Example 2:
// Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
// Output: 3
// Constraints:
// n == points.length
// 2 <= n <= 105
// points[i].length == 2
// 0 <= xi, yi <= 109

// Answer :
func MaxWidthOfVerticalArea(points [][]int) int {
    t := []int{}
    for i := 0; i < len(points); i++{
        t = append(t, points[i][0])
    }
    sort.Ints(t)
    max := 0
    for i := len(t) - 1; i >= 1; i-- {
        if t[i] - t[i - 1] > max {
            max = t[i] - t[i - 1]
        }
    }
	return max
}