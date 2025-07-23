package difficultyeasy

// Question :

// A square triple (a,b,c) is a triple where a, b, and c are integers and a2 + b2 = c2.
// Given an integer n, return the number of square triples such that 1 <= a, b, c <= n.
// Example 1:
// Input: n = 5
// Output: 2
// Explanation: The square triples are (3,4,5) and (4,3,5).
// Example 2:
// Input: n = 10
// Output: 4
// Explanation: The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).
// Constraints:
// 1 <= n <= 250

// Answer :
func CountSquareSumTriples(n int) (r int) {
	for i := 0; i <= n; i++ {
		c2 := i * i
		for j := 1; j <= i; j++ {
			b2 := j * j
			for k := 1; k <= i; k++ {
				a2 := k * k
				if a2+b2 == c2 {
					r++

				}
			}
		}
	}
	return
}