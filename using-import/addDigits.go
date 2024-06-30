package usingimportpackage

import "strconv"

// Question :

// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
// Example 1:
// Input: num = 38
// Output: 2
// Explanation: The process is
// 38 --> 3 + 8 --> 11
// 11 --> 1 + 1 --> 2
// Since 2 has only one digit, return it.
// Example 2:
// Input: num = 0
// Output: 0
// Constraints:
// 0 <= num <= 231 - 1

// Answer :
func AddDigits(num int) int {

	strNum := strconv.Itoa(num)

	for len(strNum) > 1 {

		temp := 0

		for i := 0; i < len(strNum); i++{
			num, _ := strconv.Atoi(string(strNum[i]))
			temp += num
		}

		strNum = strconv.Itoa(temp)
		
	}

	res, _ := strconv.Atoi(strNum)

	return res
}