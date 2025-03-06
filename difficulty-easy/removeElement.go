package difficultyeasy

// Question :
// Example 1:
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Answer :
func RemoveElement(nums []int, val int) int {
    result := 0
    for i := 0; i < len(nums); i++{
        if nums[i] != val {
            nums[result] = nums[i]
            result++
        }
    }
    return result
}

// func RemoveElement(nums []int, val int) int { //punyaku di awal, padahal outputnya sama aja
// 	result := 0
// 	for _, num := range nums {
// 		if num != val {
// 			result += 1
// 		}
// 	}
// 	return result
// }




// kok ngerjain lagi jir :
// func removeElement(nums []int, val int) int {
//     d := 0
//     for i := 0; i < len(nums); i++{
//         if nums[i] != val {
//             nums[d] = nums[i]
//             d++
//         }   
//     }
//     return d
// }