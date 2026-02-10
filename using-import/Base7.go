package usingimportpackage

import "strconv"

// Question :

// Given an integer num, return a string of its base 7 representation.
// Example 1:
// Input: num = 100
// Output: "202"
// Example 2:
// Input: num = -7
// Output: "-10"
// Constraints:
// -107 <= num <= 107

// Answer :

func Base7(num int) string {
	r := strconv.FormatInt(int64(num), 7)
	return r
}
