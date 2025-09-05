package usingimportpackage

import (
	"slices"
	"strconv"
)

// Question :

// Given an integer n, find the digit that occurs least frequently in its decimal representation. If multiple digits have the same frequency, choose the smallest digit.
// Return the chosen digit as an integer.
// The frequency of a digit x is the number of times it appears in the decimal representation of n.
// Example 1:
// Input: n = 1553322
// Output: 1
// Explanation:
// The least frequent digit in n is 1, which appears only once. All other digits appear twice.
// Example 2:
// Input: n = 723344511
// Output: 2
// Explanation:
// The least frequent digits in n are 7, 2, and 5; each appears only once.
// Constraints:
// 1 <= n <= 231​​​​​​​ - 1

// Answer :
func FindTheLeastFrequentDigit(n int) int {
	s,m, d := strconv.Itoa(n), map[string]int{}, 0
	for _, v := range s {
		m[string(v)]++
	}
	for _, v := range m {
		if d == 0 || v < d {
			d = v
		}
	}
	t := []int{}
	for k, v := range m {
		x, _ := strconv.Atoi(string(k))
		if  v == d {
			t = append(t, x)
		}
	}
	return slices.Min(t)
}