package usingimportpackage

import "strconv"

// Question :

// Given a positive integer num, return the number of positive integers less than or equal to num whose digit sums are even.
// The digit sum of a positive integer is the sum of all its digits.
// Example 1:
// Input: num = 4
// Output: 2
// Explanation:
// The only integers less than or equal to 4 whose digit sums are even are 2 and 4.
// Example 2:
// Input: num = 30
// Output: 14
// Explanation:
// The 14 integers less than or equal to 30 whose digit sums are even are
// 2, 4, 6, 8, 11, 13, 15, 17, 19, 20, 22, 24, 26, and 28.
// Constraints:
// 1 <= num <= 1000

// Answer :
func CountIntWithEvenDigitSum(num int) (r int) {
    for i := 1; i <= num; i++{
        if i >= 10 {
            t := 0
            s := strconv.Itoa(i)
            for j := 0; j < len(s); j++{
                n, _ := strconv.Atoi(string(s[j]))
                t += n   
            }
            if t % 2 == 0 {
                r++
            }
        } else {
            if i % 2 == 0 {
                r++
            }
        }
    }
	return
}