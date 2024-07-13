package difficultyeasy

// Question :

// Given a valid (IPv4) IP address, return a defanged version of that IP address.
// A defanged IP address replaces every period "." with "[.]".
// Example 1:
// Input: address = "1.1.1.1"
// Output: "1[.]1[.]1[.]1"
// Example 2:
// Input: address = "255.100.50.0"
// Output: "255[.]100[.]50[.]0"
// Constraints:
// The given address is a valid IPv4 address.

// Answer :
func DefangIPaddr(address string) (res string) {

	payload := "[.]"
	newArr := []string{}

	temp := ""
	for i, a := range address {

		if i == len(address) - 1 {
			temp += string(a)
			newArr = append(newArr, temp)
			break
		}
		
		if a == '.'{
			newArr = append(newArr, temp)
			temp = ""
			continue
		}
		temp += string(a)
	}

	for i, r := range newArr{
		res += r 
		if i != len(newArr) - 1 {
			res += payload
		}
	}
	
	return 

}