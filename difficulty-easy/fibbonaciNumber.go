package difficultyeasy

// Question :

// Example 1:
// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
// Example 2:
// Input: n = 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
// Example 3:
// Input: n = 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

// Answer :

func Fib(n int) int {

	a := 0
	b := 1
	var c int

	if(n == 0) {
		return a
	}

	for i := 2; i <= n; i++{
		c = a + b
		a = b
		b = c
	}

	return b
    
}
