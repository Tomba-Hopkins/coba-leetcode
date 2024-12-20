package usingimportpackage

import "strconv"

// Question :

// You are given a positive integer num consisting only of digits 6 and 9.
// Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).
// Example 1:
// Input: num = 9669
// Output: 9969
// Explanation:
// Changing the first digit results in 6669.
// Changing the second digit results in 9969.
// Changing the third digit results in 9699.
// Changing the fourth digit results in 9666.
// The maximum number is 9969.
// Example 2:
// Input: num = 9996
// Output: 9999
// Explanation: Changing the last digit 6 to 9 results in the maximum number.
// Example 3:
// Input: num = 9999
// Output: 9999
// Explanation: It is better not to apply any change.
// Constraints:
// 1 <= num <= 104
// num consists of only 6 and 9 digits.

// Answer :

func Maximum69Number(num int) int {
	p := 1
	s := strconv.Itoa(num)
    t := ""
    for i := 0; i < len(s); i++ {
        if s[i] == '6' && p > 0 {
            t += "9"
            p--
        } else {
            t += string(s[i])
        }
    }
    n, _ := strconv.Atoi(t)
    return n
}


// my ans before :

// func maximum69Number(num int) int {
// 	s := strconv.Itoa(num)
// 	t := ""
// 	r := []string{}
// 	for i := 0; i < len(s); i++{
// 		t += string(s[i])
// 		if i % 2 != 0 {
// 			r = append(r, t) 
// 			t = ""
// 			continue
// 		}
// 	}
// 	for _, x := range r {
// 		n, _ := strconv.Atoi(x)
// 		if n > 69 {
// 			n = 99
// 		} else {
// 			n = 69
// 		}
// 		l := strconv.Itoa(n)
// 		t += l
// 	}
// 	v, _ := strconv.Atoi(t)
// 	return v
// }