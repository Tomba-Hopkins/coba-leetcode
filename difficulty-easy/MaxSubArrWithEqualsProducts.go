package difficultyeasy

// Question :

// You are given an array of positive integers nums.
// An array arr is called product equivalent if prod(arr) == lcm(arr) * gcd(arr), where:
// prod(arr) is the product of all elements of arr.
// gcd(arr) is the GCD of all elements of arr.
// lcm(arr) is the LCM of all elements of arr.
// Return the length of the longest product equivalent subarray of nums.
// Example 1:
// Input: nums = [1,2,1,2,1,1,1]
// Output: 5
// Explanation:
// The longest product equivalent subarray is [1, 2, 1, 1, 1], where prod([1, 2, 1, 1, 1]) = 2, gcd([1, 2, 1, 1, 1]) = 1, and lcm([1, 2, 1, 1, 1]) = 2.
// Example 2:
// Input: nums = [2,3,4,5,6]
// Output: 3
// Explanation:
// The longest product equivalent subarray is [3, 4, 5].
// Example 3:
// Input: nums = [1,2,3,1,4,5,1]
// Output: 5
// Constraints:
// 2 <= nums.length <= 100
// 1 <= nums[i] <= 10

// Answer :
func MaxSubArrWithEqualsProducts(nums []int) int {
	for i := len(nums); i > 1; i--{
		for j := 0; j < len(nums); j++{
			var x []int
			prod := 0
			if j + i <= len(nums) {
				x = nums[j:j+i]
			} else {
				break
			}
			for in, v := range x {
				if in == 0 {
					prod = v
				} else {
					prod *= v
				}
			}
			g, l := x[0], x[0]
			for k := 1; k < len(x); k++{
				g = gcd(g, x[k])
				l = lcm(l, x[k])
			}
			if prod == g * l {
				return  len(x)
			}	
		}	
	}	
	return 0
}

// great common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return  a
}

// least commond multiple
func lcm(a, b int) int {
	return  (a * b) / gcd(a, b)
}
