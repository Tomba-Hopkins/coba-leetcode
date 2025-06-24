package difficultyeasy

// Question :

// You are given an 8 x 8 matrix representing a chessboard. There is exactly one white rook represented by 'R', some number of white bishops 'B', and some number of black pawns 'p'. Empty squares are represented by '.'.
// A rook can move any number of squares horizontally or vertically (up, down, left, right) until it reaches another piece or the edge of the board. A rook is attacking a pawn if it can move to the pawn's square in one move.
// Note: A rook cannot move through other pieces, such as bishops or pawns. This means a rook cannot attack a pawn if there is another piece blocking the path.
// Return the number of pawns the white rook is attacking.
// Example 1:
// Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 3
// Explanation:
// In this example, the rook is attacking all the pawns.
// Example 2:
// Input: board = [[".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 0
// Explanation:
// The bishops are blocking the rook from attacking any of the pawns.
// Example 3:
// Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 3
// Explanation:
// The rook is attacking the pawns at positions b5, d6, and f5.
// Constraints:
// board.length == 8
// board[i].length == 8
// board[i][j] is either 'R', '.', 'B', or 'p'
// There is exactly one cell with board[i][j] == 'R'

// Answer :
func AvailableCapturesForRook(board [][]byte) (r int) {
	row, col := 0, 0
	for i := 0; i < len(board); i++{
		d := false
		for j, b := range board[i] {
			if b == 'R'{
				row = i
				col = j
				d = true
				break
			}
		}
		if d {
			break
		}
	}
	for i := row; i < len(board); i++{
		if board[i][col] == 'p' {
			r++
			break
		} else if board[i][col] == 'B'{
			break
		}
	}
	for i := row; i >= 0; i--{
		if board[i][col] == 'p' {
			r++
			break
		}else if board[i][col] == 'B'{
			break
		}
	}

	for j := col; j >= 0; j--{
		if board[row][j] == 'p' {
			r++
			break
		} else if board[row][j] == 'B'{
			break
		}
	}
	for j := col; j < len(board[row]); j++{
		if board[row][j] == 'p' {
			r++
			break
		} else if board[row][j] == 'B'{
			break
		}
	}
	return
}