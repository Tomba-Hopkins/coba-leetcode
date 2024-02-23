package difficultyeasy

// Question :
// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Answer :
func IsPalindrome(x int) bool {

	if x < 10 && x >= 0 {
		return true
	} else if x < 0 {
		return false
	}

	first := x

	for first >= 10 {
		first /= 10 // 125 -> 125 / 10 = 12,5  -> 12,5 / 10 = 1,25 -> selagi di atas 10, bagi 10 terus
	}

	last := x % 10

	if first == last {
		return true
	} else {
		return false
	}
} 
// masih salah entah kenapa kalau input 1000021 hasilnya harus false, punyaku true -> bukannya itu palindrome ya? soalnya digit satu dan awalnya sama



// gpt -> ini bener:
// func isPalindromeGpt(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	original := x
// 	reversed := 0
// 	for x > 0 {
// 		digit := x % 10
// 		fmt.Println("===========================")
// 		fmt.Println("digit: ", digit)
// 		reversed = reversed*10 + digit
// 		fmt.Println("reverse: ", reversed)
// 		x /= 10
// 		fmt.Println("x: ", x)
// 		fmt.Println("===========================")
// 		fmt.Println("")
// 	}
// 	return original == reversed
// }