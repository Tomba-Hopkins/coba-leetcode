package difficultyeasy

// Question :

// You are given an integer array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
// Example 1:
// Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// Output: true
// Example 2:
// Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
// Output: false
// Constraints:
// 2 <= coordinates.length <= 1000
// coordinates[i].length == 2
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
// coordinates contains no duplicate point.

// Answer :
// formula from gpt: y2 - y1 / x2 - x1 -> x1, y1 | x2, y2
func CheckIfItIsAnStraightLine(coordinates [][]int) bool {
	dx := coordinates[1][0] - coordinates[0][0]
	dy := coordinates[1][1] - coordinates[0][1]
	for i := 0; i < len(coordinates)-1; i++ {
		x1, y1 := coordinates[i][0], coordinates[i][1]
		x2, y2 := coordinates[i+1][0], coordinates[i+1][1]
		if dy*(x2-x1) != dx*(y2-y1) {
			return false
		}
	}
	return true
}
