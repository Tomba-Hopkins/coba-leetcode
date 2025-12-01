package usingimportpackage

import "strconv"

// Question :

// Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.
// Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.
// Example 1:
// Input: path = "NES"
// Output: false
// Explanation: Notice that the path doesn't cross any point more than once.
// Example 2:
// Input: path = "NESWW"
// Output: true
// Explanation: Notice that the path visits the origin twice.
// Constraints:
// 1 <= path.length <= 104
// path[i] is either 'N', 'S', 'E', or 'W'.

// Answer :
// takut gini misal -> x = 11 y = 1 -> 111 atau x = 1, y = 11 -> 111, makanya bypass pake x11y1 or x1y11
func PathCrossing(path string) bool {
	x, y := 0, 0
	m := map[string]bool{}
	m["x0y0"] = true
	for _, v := range path {
		switch v {
		case 'N':
			y++
		case 'E':
			x++
		case 'S':
			y--
		case 'W':
			x--
		}
		s1, s2 := strconv.Itoa(x), strconv.Itoa(y)
		if m["x"+s1+"y"+s2]{
			return true
		}
		m["x"+s1+"y"+s2] = true
	}
	return false
}