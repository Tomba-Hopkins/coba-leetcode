package difficultyeasy

// Question :

// Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.
// Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.
// Example 1:
// Input: nums = [1,2,2,3,1]
// Output: 2
// Explanation:
// The input array has a degree of 2 because both elements 1 and 2 appear twice.
// Of the subarrays that have the same degree:
// [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
// The shortest length is 2. So return 2.
// Example 2:
// Input: nums = [1,2,2,3,1,4,2]
// Output: 6
// Explanation:
// The degree is 3 because the element 2 is repeated 3 times.
// So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.
// Constraints:
// nums.length will be between 1 and 50,000.
// nums[i] will be an integer between 0 and 49,999.

// Answer :
// my long ans
func DegreeOfAnArr(nums []int) (r int) {
	m, h, a := map[int]int{}, 0, []int{}
	for _, n := range nums {
		m[n]++
	}
	for _, n := range nums {
		if m[n] > h {
			h = m[n]
		}
	}
	for k, v := range m {
		if v == h {
			a = append(a, k)
		}
	}
	f, e := map[int]int{}, map[int]int{}
	for i := 0; i < len(a); i++{
		for j := 0; j < len(nums); j++{
			if nums[j] == a[i] {
				f[a[i]] = j + 1
				break
			}
		}
		for j := len(nums) - 1; j >= 0; j--{
			if nums[j] == a[i] {
				e[a[i]] = j + 1
				break
			}
		}
		
	}
	for _, v := range a {
		l := (e[v] - f[v]) + 1
		if  l < r || r == 0 {
			r = l
		}
	}
	return
}

// TLE ans
// func DegreeOfAnArr(nums []int) (r int) {
// 	m, m2, h := map[int]int{}, map[int]bool{}, 0
// 	for _, n := range nums {
// 		m[n]++
// 	}
// 	for _, n := range nums {
// 		if m[n] > h {
// 			h = m[n]
// 		}
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		x := nums[i]
// 		for j := len(nums) - 1; j >= i; j--{
// 			y := nums[j]
// 			if x == y && m[x] == h && !m2[x] {
// 				a := nums[i:j + 1]
// 				if r == 0 {
// 					r = len(a)
// 				} else {
// 					if len(a) < r {
// 						r = len(a)
// 					}
// 				}
// 				m2[x] = true
// 				break
// 			}
// 		}
// 	}
// 	return
// }