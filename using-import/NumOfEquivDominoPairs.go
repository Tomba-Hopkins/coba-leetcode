package usingimportpackage

import "sort"

// Question :

// Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a == c and b == d), or (a == d and b == c) - that is, one domino can be rotated to be equal to another domino.
// Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].
// Example 1:
// Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
// Output: 1
// Example 2:
// Input: dominoes = [[1,2],[1,2],[1,1],[1,2],[2,2]]
// Output: 3
// Constraints:
// 1 <= dominoes.length <= 4 * 104
// dominoes[i].length == 2
// 1 <= dominoes[i][j] <= 9

// Answer :
func NumOfEquivDominoPairs(dominoes [][]int) (r int) {
	for _, d := range dominoes {
		sort.Ints(d)
	}
	m := map[int]bool{}
	for i := 0; i < len(dominoes) - 1; i++{
		c := 0
		if !m[i] {
			for j := i + 1; j < len(dominoes); j++{
				if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1] {
					m[j] = true
					m[i] = true
					c++
				}
			}
		}
		if c > 0 {
			c++
			r += c * (c - 1) / 2
		}
	}
	return
}