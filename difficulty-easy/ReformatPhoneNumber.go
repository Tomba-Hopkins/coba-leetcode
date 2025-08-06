package difficultyeasy

// Question :

// You are given a phone number as a string number. number consists of digits, spaces ' ', and/or dashes '-'.
// You would like to reformat the phone number in a certain manner. Firstly, remove all spaces and dashes. Then, group the digits from left to right into blocks of length 3 until there are 4 or fewer digits. The final digits are then grouped as follows:
// 2 digits: A single block of length 2.
// 3 digits: A single block of length 3.
// 4 digits: Two blocks of length 2 each.
// The blocks are then joined by dashes. Notice that the reformatting process should never produce any blocks of length 1 and produce at most two blocks of length 2.
// Return the phone number after formatting.
// Example 1:
// Input: number = "1-23-45 6"
// Output: "123-456"
// Explanation: The digits are "123456".
// Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
// Step 2: There are 3 digits remaining, so put them in a single block of length 3. The 2nd block is "456".
// Joining the blocks gives "123-456".
// Example 2:
// Input: number = "123 4-567"
// Output: "123-45-67"
// Explanation: The digits are "1234567".
// Step 1: There are more than 4 digits, so group the next 3 digits. The 1st block is "123".
// Step 2: There are 4 digits left, so split them into two blocks of length 2. The blocks are "45" and "67".
// Joining the blocks gives "123-45-67".
// Example 3:
// Input: number = "123 4-5678"
// Output: "123-456-78"
// Explanation: The digits are "12345678".
// Step 1: The 1st block is "123".
// Step 2: The 2nd block is "456".
// Step 3: There are 2 digits left, so put them in a single block of length 2. The 3rd block is "78".
// Joining the blocks gives "123-456-78".
// Constraints:
// 2 <= number.length <= 100
// number consists of digits and the characters '-' and ' '.
// There are at least two digits in number.

// Answer :
func ReformatPhoneNumber(number string) (r string) {
	x, d := "", 0
	for _, n := range number {
		if n >= '0' && n <= '9'{
			x += string(n)
		}
	}
	a := []string{}
	for i := 0; i < len(x); i += d{
		z := len(x[i + 1:])
		if len(x) - i <= 3 {
			a = append(a, x[i:])
			break
		}
		if z > 3 {
			d = 3
			a = append(a, x[i:i+3])
		} else if z > 2 {
			d = 2 
			a = append(a, x[i:i+2])
		}
	}
	for i, v := range a {
		if i == len(a) - 1 {
			r += v
			break
		}
		r += v + "-"
	}
	return
}