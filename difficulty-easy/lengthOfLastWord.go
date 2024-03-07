package difficultyeasy

// Question :
// Example 1:
// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:
// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:
// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

// Answer :

func LengthOfLastWord(s string) int {
    arr := []string{}

    for string(s[len(s) - 1]) == " "{
        s = s[:len(s) - 1]
    }

    for i := len(s) - 1; i >= 0; i--{
        if string(s[i]) == " "{
            break
        }
        arr = append(arr, string(s[i]))
    }

    return len(arr)
}