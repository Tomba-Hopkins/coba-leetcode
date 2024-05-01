package difficultyeasy

// Question :

// Example 1:
// Input: word = "abcdefd", ch = "d"
// Output: "dcbaefd"
// Explanation: The first occurrence of "d" is at index 3.
// Reverse the part of word from 0 to 3 (inclusive), the resulting string is "dcbaefd".
// Example 2:
// Input: word = "xyxzxe", ch = "z"
// Output: "zxyxxe"
// Explanation: The first and only occurrence of "z" is at index 3.
// Reverse the part of word from 0 to 3 (inclusive), the resulting string is "zxyxxe".
// Example 3:
// Input: word = "abcd", ch = "z"
// Output: "abcd"
// Explanation: "z" does not exist in word.
// You should not do any reverse operation, the resulting string is "abcd".

// Answer :


func ReversePrefix(word string, ch byte) (res string) {
    indexCH := 0
	
	for i := 0; i < len(word); i++{
		if word[i] == ch{
			indexCH = i
			break
		}
	}

	for j := indexCH; j >= 0; j--{
		res += string(word[j])
	}

	for k := indexCH; k < len(word); k++{

		if k == indexCH{
			continue
		}
		
		res += string(word[k])
	}
	
	
	return 
}