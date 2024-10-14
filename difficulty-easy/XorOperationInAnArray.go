package difficultyeasy

// Question :

// You are given an integer n and an integer start.
// Define an array nums where nums[i] = start + 2 * i (0-indexed) and n == nums.length.
// Return the bitwise XOR of all elements of nums.
// Example 1:
// Input: n = 5, start = 0
// Output: 8
// Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
// Where "^" corresponds to bitwise XOR operator.
// Example 2:
// Input: n = 4, start = 3
// Output: 8
// Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.
// Constraints:
// 1 <= n <= 1000
// 0 <= start <= 1000
// n == nums.length

// Answer :
func XorOperationInAnArray(n int, start int) (r int) {
	for i := 0; i < n; i++ {
		r ^= start
		start += 2
	}
	return
}


// My Manual experiment :
// func XorOperationInAnArray(n int, start int) (r int) {
//     for i := 0; i < n; i++{
//         // r ^= start
//         binStart := strconv.FormatInt(int64(start), 2)
//         binTemp := strconv.FormatInt(int64(r), 2)
//         for len(binTemp) < len(binStart) {
//             binTemp += "0"
//         }
//         revTemp := ""
//         for j := len(binTemp) - 1; j >= 0; j--{
//             revTemp += string(binTemp[j])
//         }
//         fmt.Println(revTemp, binStart)
//         temp := ""
//         for j := 0; j < len(binStart); j++{
//             if binStart[j] == revTemp[j] {
//                 temp += "0"
//             } else {
//                 temp += "1"
//             }
//         }
//         fmt.Println(temp)
//         tempInt, _ := strconv.ParseInt(temp, 2, 64)
//         r += int(tempInt)
//         start += 2
//     }
//     return
// }