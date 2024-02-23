package main

// func isValid(s string) bool {

// 	stack := []byte{}

// 	kotak := map[byte]byte{
// 		'(' : ')',
// 		'{' : '}',
// 		'[' : ']',
// 	}

// 	for i := 0; i < len(s); i++{
// 		karakter := s[i]

// 		if _, ok := kotak[karakter]; ok{
// 			stack = append(stack, karakter)
// 		} else {
// 			if len(s) == 0 ||karakter != kotak[stack[len(stack) - 1]] {
// 				return false
// 			}
// 			stack = stack[:len(stack) - 1]
// 		}
// 	}
// 	return len(stack) == 0
// }

func main() {
	// text := "(}"
	// fmt.Println(isValid(text))
}