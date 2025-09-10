package difficultyeasy

// Question :

// You are given an integer n. Your task is to compute the GCD (greatest common divisor) of two values:
// sumOdd: the sum of the first n odd numbers.
// sumEven: the sum of the first n even numbers.
// Return the GCD of sumOdd and sumEven.
// Example 1:
// Input: n = 4
// Output: 4
// Explanation:
// Sum of the first 4 odd numbers sumOdd = 1 + 3 + 5 + 7 = 16
// Sum of the first 4 even numbers sumEven = 2 + 4 + 6 + 8 = 20
// Hence, GCD(sumOdd, sumEven) = GCD(16, 20) = 4.
// Example 2:
// Input: n = 5
// Output: 5
// Explanation:
// Sum of the first 5 odd numbers sumOdd = 1 + 3 + 5 + 7 + 9 = 25
// Sum of the first 5 even numbers sumEven = 2 + 4 + 6 + 8 + 10 = 30
// Hence, GCD(sumOdd, sumEven) = GCD(25, 30) = 5.
// Constraints:
// 1 <= n <= 10​​​​​​​00

// Answer :
func GcdOfOddAndEvenSums(n int) int {
	x, y := 0, 0
	d1, d2 := 0, 0
	c := 1
	for d2 < n {
		if c%2 != 0 {
			x += c
			d1++
		} else {
			y += c
			d2++
		}
		c++
	}
	return y - x
}