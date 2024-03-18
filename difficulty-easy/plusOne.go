package difficultyeasy

// Question :

// Example 1:
// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Incrementing by one gives 123 + 1 = 124.
// Thus, the result should be [1,2,4].
// Example 2:
// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
// Incrementing by one gives 4321 + 1 = 4322.
// Thus, the result should be [4,3,2,2].
// Example 3:
// Input: digits = [9]
// Output: [1,0]
// Explanation: The array represents the integer 9.
// Incrementing by one gives 9 + 1 = 10.
// Thus, the result should be [1,0].

// Answer :

func PlusOne(digits []int) []int {
    n := len(digits)

    for i := n - 1; i >= 0; i-- {
        digits[i]++
        
        if digits[i] < 10 {
            return digits
        }
        
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}


// Kenapa ini salah :
// func plusOne(digits []int) []int {
// 	result := []int{}
// 	for i := 0; i < len(digits); i++ {
// 		if i == len(digits)-1 {
// 			digits[i] += 1
// 			if digits[i] >= 10 {
// 				str := strconv.Itoa(digits[i])
// 				for _, s := range str {
// 					toInt, _ := strconv.Atoi(string(s))
// 					result = append(result, toInt)
// 				}
// 			} else {
// 				result = append(result, digits[i])
// 			}
// 		} else {
// 			result = append(result, digits[i])
// 		}
// 	}
// 	return result
// }