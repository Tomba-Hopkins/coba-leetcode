package usingimportpackage

import (
	"sort"
	"strconv"
)

// Question :
// Example 1:
// Input: score = [5,4,3,2,1]
// Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
// Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].
// Example 2:
// Input: score = [10,3,8,9,4]
// Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
// Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].

// Answer :
func FindRelativeRanks(score []int) (res []string) {

	sortScore := make([]int, len(score))
	copy(sortScore, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sortScore)))

	peringkat := map[int]string{}

	for i := 0; i < len(sortScore); i++{
		if i == 0 {
			peringkat[sortScore[i]] = "Gold Medal"
		} else if i == 1 {
			peringkat[sortScore[i]] = "Silver Medal"
		} else if i == 2 {
			peringkat[sortScore[i]] = "Bronze Medal"
		}
	}

	if len(score) > 3 {
		for i := 3; i < len(sortScore); i++{
			strI := strconv.Itoa(i + 1)
			peringkat[sortScore[i]] = strI
		}
	}

	for _, r := range score{
		res = append(res, peringkat[r])
	}


	return 

}