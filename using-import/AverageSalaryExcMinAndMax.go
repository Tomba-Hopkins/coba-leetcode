package usingimportpackage

import "slices"

// Question :

// You are given an array of unique integers salary where salary[i] is the salary of the ith employee.
// Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.
// Example 1:
// Input: salary = [4000,3000,1000,2000]
// Output: 2500.00000
// Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
// Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500
// Example 2:
// Input: salary = [1000,2000,3000]
// Output: 2000.00000
// Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
// Average salary excluding minimum and maximum salary is (2000) / 1 = 2000
// Constraints:
// 3 <= salary.length <= 100
// 1000 <= salary[i] <= 106
// All the integers of salary are unique.

// Answer :
func AverageSalaryExcMinAndMax(salary []int) float64 {
	a, z := []int{}, 0
	x, y := slices.Max(salary), slices.Min(salary)
	for _, s := range salary {
		if s != x && s != y {
			a = append(a, s)
		}
	}
	for _, v := range a {
		z += v
	}
	return float64(z) / float64(len(a))
}