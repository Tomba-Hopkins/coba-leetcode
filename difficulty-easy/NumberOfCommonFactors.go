package difficultyeasy

// Question :

// Given two positive integers a and b, return the number of common factors of a and b.
// An integer x is a common factor of a and b if x divides both a and b.
// Example 1:
// Input: a = 12, b = 6
// Output: 4
// Explanation: The common factors of 12 and 6 are 1, 2, 3, 6.
// Example 2:
// Input: a = 25, b = 30
// Output: 2
// Explanation: The common factors of 25 and 30 are 1, 5.
// Constraints:
// 1 <= a, b <= 1000

// Answer :
func NumberOfCommonFactors(a int, b int) (r int) {
	x := min(a, b)
	for i := 1; i <= x; i++{
		if a % i == 0 && b % i == 0 {
			r++
		}
	}
	return
}