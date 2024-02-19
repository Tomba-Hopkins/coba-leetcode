package difficultyeasy

// Question :
// Example 1:
// Input: n = 1
// Output: true
// Explanation: 20 = 1
// Example 2:
// Input: n = 16
// Output: true
// Explanation: 24 = 16
// Example 3:
// Input: n = 3
// Output: false

// Answer :
func isPowerOfTwo(n int) bool {
    pow := 1
    for pow < n {
        pow *= 2
    }
    return pow == n
}