package difficultyeasy

// Question :

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

// Answer :


func Intersection(nums1 []int, nums2 []int) []int {

	table := map[int]bool{}

	for _, n1 := range nums1 {
		table[n1] = true
	}

	result := []int{}
	for _, n2 := range nums2 {
		if table[n2] {
			result = append(result, n2)
			table[n2] = false
		}
	}

	return result

}