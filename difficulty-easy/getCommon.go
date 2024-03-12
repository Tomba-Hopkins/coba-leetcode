package difficultyeasy

// Question :
// Example 1:
// Input: nums1 = [1,2,3], nums2 = [2,4]
// Output: 2
// Explanation: The smallest element common to both arrays is 2, so we return 2.
// Example 2:
// Input: nums1 = [1,2,3,6], nums2 = [2,3,4,5]
// Output: 2
// Explanation: There are two common elements in the array 2 and 3 out of which 2 is the smallest, so 2 is returned.

// Answer :

func GetCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2){
		if nums1[i] == nums2[j] {
            return nums1[i]
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
	}

	return -1
}


// Awalnya gini :

// func getCommon(nums1 []int, nums2 []int) int {
// 	table := map[int]int{}
// 	for i := 0; i < len(nums1); i++ {
// 		table[nums1[i]]++
// 	}
// 	for i := 0; i < len(nums2); i++{
// 		table[nums2[i]]++
// 	}
// 	result := nums1[0]
// 	for _, n := range nums1{
// 		for table[n] > table[result]{
// 			result = n
// 			break
// 		}
// 	}
// 	fmt.Println(table)
// 	return result
// }