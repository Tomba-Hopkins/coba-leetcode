package difficultyeasy

// Question :

// Given three integer arrays nums1, nums2, and nums3, return a distinct array containing all the values that are present in at least two out of the three arrays. You may return the values in any order.
// Example 1:
// Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
// Output: [3,2]
// Explanation: The values that are present in at least two arrays are:
// - 3, in all three arrays.
// - 2, in nums1 and nums2.
// Example 2:
// Input: nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
// Output: [2,3,1]
// Explanation: The values that are present in at least two arrays are:
// - 2, in nums2 and nums3.
// - 3, in nums1 and nums2.
// - 1, in nums1 and nums3.
// Example 3:
// Input: nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
// Output: []
// Explanation: No value is present in at least two arrays.
// Constraints:
// 1 <= nums1.length, nums2.length, nums3.length <= 100
// 1 <= nums1[i], nums2[j], nums3[k] <= 100

// Answer :
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (r []int) {
	b := map[int]bool{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] && !b[nums2[j]] {
				r = append(r, nums2[j])
				b[nums1[i]] = true
			}
			for k := 0; k < len(nums3); k++ {
				if nums3[k] == nums1[i] && !b[nums3[k]] {
					r = append(r, nums3[k])
					b[nums1[i]] = true
				}
				if nums3[k] == nums2[j] && !b[nums3[k]] {
					r = append(r, nums3[k])
					b[nums2[j]] = true
				}
			}
		}
	}
	return
}