package usingimportpackage

import "sort"

// Question :

// Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.
// Return the sorted array.
// Example 1:
// Input: nums = [1,1,2,2,2,3]
// Output: [3,1,1,2,2,2]
// Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
// Example 2:
// Input: nums = [2,3,1,3,2]
// Output: [1,3,3,2,2]
// Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
// Example 3:
// Input: nums = [-1,1,-6,4,5,-6,1,4,1]
// Output: [5,-1,4,4,-6,-6,1,1,1]
// Constraints:
// 1 <= nums.length <= 100
// -100 <= nums[i] <= 100

// Answer :
func FrequencySort(nums []int) (r []int) {
	n := map[int]int{}
	for _, v := range nums {
		n[v]++
	}
	min, max := 10, 0
	for _, v := range nums {
		if n[v] > max {
			max = n[v]
		}

		if n[v] < min {
			min = n[v]
		}
	}
	for i := min; i <= max; i++ {
        temp := []int{}
		for k := 0; k < len(nums); k++ {
			if n[nums[k]] == i {
				temp = append(temp, nums[k])
			}
		}
        sort.Sort(sort.Reverse(sort.IntSlice(temp)))
        r = append(r, temp...)
	}
	return
}