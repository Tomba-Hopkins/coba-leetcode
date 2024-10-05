package difficultyeasy

// Question :

// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
// Example 1:
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Example 2:
// Input: s = "Mr Ding"
// Output: "rM gniD"
// Constraints:
// 1 <= s.length <= 5 * 104
// s contains printable ASCII characters.
// s does not contain any leading or trailing spaces.
// There is at least one word in s.
// All the words in s are separated by a single space.

// Answer :
func ReverseWords3(s string) (r string) {

	temp := []string{}
	space := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			space = append(space, i)
		}
	}

	if len(space) == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			r += string(s[i])
		}
		return
	}

	for i := 0; i < len(space); i++ {
		if i < 1 {
			temp = append(temp, s[:space[i]])
		} else {
			temp = append(temp, s[space[i-1]+1:space[i]])
		}
	}
	temp = append(temp, s[space[len(space)-1]+1:])
	for i := 0; i < len(temp); i++ {
		for j := len(temp[i]) - 1; j >= 0; j-- {
			r += string(temp[i][j])
		}
		if i != len(temp)-1 {
			r += " "
		}
	}
	return
}


// Use package :
// func ReverseWords3(s string) (r string){
//     t := strings.Split(s, " ")
//     t2 := []string{}
//     for i := 0; i < len(t); i++{
//         w := ""
//         for j := len(t[i]) -1; j >= 0; j-- {
//             w += string(t[i][j])
//         }
//         t2 = append(t2, w)
//     }
//     return strings.Join(t2, " ")
// }