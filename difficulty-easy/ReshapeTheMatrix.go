package difficultyeasy

// Question :

// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
// You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
// The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
// If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
// Example 1:
// Input: mat = [[1,2],[3,4]], r = 1, c = 4
// Output: [[1,2,3,4]]
// Example 2:
// Input: mat = [[1,2],[3,4]], r = 2, c = 4
// Output: [[1,2],[3,4]]
// Constraints:
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 100
// -1000 <= mat[i][j] <= 1000
// 1 <= r, c <= 300

// Answer :
func ReshapeTheMatrix(mat [][]int, r int, c int) (res [][]int) {
	cN, rN := len(mat[0]), len(mat)
	if cN*rN != r*c { // awalnya pake < salah
		return mat
	}
	t := []int{}
	for _, m := range mat {
		t = append(t, m...)
	}
	a := []int{}
	for i := 0; i < len(t); i++ {
		if len(a) == c {
			res = append(res, a)
			a = []int{}
			a = append(a, t[i])
		} else {
			a = append(a, t[i])
		}
	}
	if len(a) > 0 {
		res = append(res, a)
	}
	return
}