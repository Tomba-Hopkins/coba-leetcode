package usingimportpackage

import "strconv"

// Question :

// You are given coordinates, a string that represents the coordinates of a square of the chessboard. Below is a chessboard for your reference.
// Return true if the square is white, and false if the square is black.
// The coordinate will always represent a valid chessboard square. The coordinate will always have the letter first, and the number second.
// Example 1:
// Input: coordinates = "a1"
// Output: false
// Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.
// Example 2:
// Input: coordinates = "h3"
// Output: true
// Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.
// Example 3:
// Input: coordinates = "c7"
// Output: false
// Constraints:
// coordinates.length == 2
// 'a' <= coordinates[0] <= 'h'
// '1' <= coordinates[1] <= '8'

// Answer :
func DetermineColorOfAChessboardSquare(coordinates string) bool {
	m := map[string]bool{}
    w := "abcdefgh"
    for i := 0; i < len(w); i++{
        if i == 0 || i % 2 == 0 {
            x := false
            for j := 1; j <= 8; j++{
                n := strconv.Itoa(j)
                if j % 2 == 0 {
                    m[string(w[i]) + n] = !x
                } else {
                    m[string(w[i]) + n] = x
                }
            }
        } else {
            x := true
            for j := 1; j <= 8; j++{
                n := strconv.Itoa(j)
                if  j % 2 == 0 {
                    m[string(w[i]) + n] = !x
                } else {
                    m[string(w[i]) + n] = x
                }
            }
        }
    }
    return m[coordinates]
}
