package usingimportpackage

import "strconv"

// Question :

// You are given a string num, representing a large integer. Return the largest-valued odd integer (as a string) that is a non-empty substring of num, or an empty string "" if no odd integer exists.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
// Input: num = "52"
// Output: "5"
// Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.
// Example 2:
// Input: num = "4206"
// Output: ""
// Explanation: There are no odd numbers in "4206".
// Example 3:
// Input: num = "35427"
// Output: "35427"
// Explanation: "35427" is already an odd number.
// Constraints:
// 1 <= num.length <= 105
// num only consists of digits and does not contain any leading zeros.

// Answer :
func LargestOddNumberInStr(num string) string {
	l, _ := strconv.Atoi(string(num[len(num) - 1]))
	for len(num) > 1 {
		if l % 2 != 0 {
			return num
		}
		num = num[:len(num) - 1]
		l, _ = strconv.Atoi(string(num[len(num) - 1]))
	
	}
	if l % 2 != 0 {
		return num
	}
	return ""
}


// my ans before
// func LargestOddNumberInStr(num string) string {
//     f, _ := strconv.Atoi(num)
//     if f % 2 != 0 {
//         return num
//     }
//     t := ""
//     n := 0
//     for i := 0; i < len(num); i++{
//         s := num[i:]
//         for j := len(s) - 1; j >= 0; j-- {
//             temp, _ := strconv.Atoi(s[:j])
//             if temp > n && temp % 2 != 0 {
//                 n = temp
//                 t = s[:j]
//             }
            
//         }
        
//     }
// 	return t
// }