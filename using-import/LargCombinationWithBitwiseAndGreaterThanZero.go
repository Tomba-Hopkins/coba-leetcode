package usingimportpackage

import "strconv"

// Question :

// The bitwise AND of an array nums is the bitwise AND of all integers in nums.
// For example, for nums = [1, 5, 3], the bitwise AND is equal to 1 & 5 & 3 = 1.
// Also, for nums = [7], the bitwise AND is 7.
// You are given an array of positive integers candidates. Compute the bitwise AND for all possible combinations of elements in the candidates array.
// Return the size of the largest combination of candidates with a bitwise AND greater than 0.
// Example 1:
// Input: candidates = [16,17,71,62,12,24,14]
// Output: 4
// Explanation: The combination [16,17,62,24] has a bitwise AND of 16 & 17 & 62 & 24 = 16 > 0.
// The size of the combination is 4.
// It can be shown that no combination with a size greater than 4 has a bitwise AND greater than 0.
// Note that more than one combination may have the largest size.
// For example, the combination [62,12,24,14] has a bitwise AND of 62 & 12 & 24 & 14 = 8 > 0.
// Example 2:
// Input: candidates = [8,8]
// Output: 2
// Explanation: The largest combination [8,8] has a bitwise AND of 8 & 8 = 8 > 0.
// The size of the combination is 2, so we return 2.
// Constraints:
// 1 <= candidates.length <= 105
// 1 <= candidates[i] <= 107

// Answer :
func LargCombinationWithBitwiseAndGreaterThanZero(candidates []int) (r int) {
	m := map[int]int{}
	for i := 0; i < len(candidates); i++ {
		s, revS := strconv.FormatInt(int64(candidates[i]), 2), ""
		for j := len(s) - 1; j >= 0; j-- {
			revS += string(s[j])
		}
		for j, v := range revS {
			if v == '1' {
				m[j]++
			}
		}
	}
	d, h := 0, 0
	for k, v := range m {
		if v > h {
			h = v
			d = k
		}
		if v == h {
			if k > d {
				d = k
			}
		}
	}
	for i := 0; i < len(candidates); i++ {
		s, revS := strconv.FormatInt(int64(candidates[i]), 2), ""
		for j := len(s) - 1; j >= 0; j-- {
			revS += string(s[j])
		}
		if len(revS) > d {
			if revS[d] == '1' {
				r++
			}
		}
	}
	return
}
