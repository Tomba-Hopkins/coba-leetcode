package difficultymedium

// Question :

// You are given an integer array values where values[i] represents the value of the ith sightseeing spot. Two sightseeing spots i and j have a distance j - i between them.
// The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j: the sum of the values of the sightseeing spots, minus the distance between them.
// Return the maximum score of a pair of sightseeing spots.
// Example 1:
// Input: values = [8,1,5,2,6]
// Output: 11
// Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
// Example 2:
// Input: values = [1,2]
// Output: 2
// Constraints:
// 2 <= values.length <= 5 * 104
// 1 <= values[i] <= 1000

// Answer:

func BestSightseeingPair(values []int) int {
	max := 0
    bef := values[0]
    for j := 1; j < len(values); j++{
        if bef + values[j] - j > max {
            max = bef + values[j] - j
        }
        if values[j] + j > bef {
            bef = values[j] + j
        }
    }
	return max
}


// my ans before
// func BestSightseeingPair(values []int) int {
// 	x := 0
//     for i := 0; i < len(values) - 1; i++{
//         for j := i + 1; j < len(values); j++{
//             if values[i] + values[j] + i - j > x {
//                 x = values[i] + values[j] + i - j
//             }
//         }
//     }
// 	return x
// }