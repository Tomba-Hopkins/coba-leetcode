package usingimportpackage

import "strings"

// Question :

// You are given a string caption representing the caption for a video.
// The following actions must be performed in order to generate a valid tag for the video:
// Combine all words in the string into a single camelCase string prefixed with '#'. A camelCase string is one where the first letter of all words except the first one is capitalized. All characters after the first character in each word must be lowercase.
// Remove all characters that are not an English letter, except the first '#'.
// Truncate the result to a maximum of 100 characters.
// Return the tag after performing the actions on caption.
// Example 1:
// Input: caption = "Leetcode daily streak achieved"
// Output: "#leetcodeDailyStreakAchieved"
// Explanation:
// The first letter for all words except "leetcode" should be capitalized.
// Example 2:
// Input: caption = "can I Go There"
// Output: "#canIGoThere"
// Explanation:
// The first letter for all words except "can" should be capitalized.
// Example 3:
// Input: caption = "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
// Output: "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
// Explanation:
// Since the first word has length 101, we need to truncate the last two letters from the word.
// Constraints:
// 1 <= caption.length <= 150
// caption consists only of English letters and ' '.

// Answer :
func GenerateTagForVidCaption(caption string) (r string) {
	r += "#"
	c := ""
	for _, v := range caption {
		if v == ' ' && len(c) == 0 {
			continue
		} 
		c += string(v)
	}
	a := strings.Split(c, " ")
	for i := 0; i < len(a); i++{
		s := ""
		if i == 0 {
			s = strings.ToLower(a[i])
		} else {
			for j := 0; j < len(a[i]); j++{
				var x string
				if j == 0 {
					x = strings.ToUpper(string(a[i][j]))
					s += x
				} else {
					x = strings.ToLower(string(a[i][j]))
					s += x
				}
			}
		}
		r += s	
	}
	if len(r) > 99 {
		r = r[:100]
	}
	return
}
