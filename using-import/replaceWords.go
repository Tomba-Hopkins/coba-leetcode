package usingimportpackage

import (
	"sort"
	"strings"
)

// Question :
// In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word derivative. For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".
// Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the derivatives in the sentence with the root forming it. If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.
// Return the sentence after the replacement.
// Example 1:
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"
// Example 2:
// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"

// Answer :
func ReplaceWords(dictionary []string, sentence string) string {

	sArr := strings.Split(sentence, " ")
	newArr := []string{}

	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})


	for i := 0; i < len(sArr); i++{

		breakin := false
		
		for j := 0; j < len(dictionary); j++{
			determine := 0
			for k := 0; k < len(dictionary[j]); k++{

				if len(dictionary[j]) > len(sArr[i]){
					continue
				}
				
				if dictionary[j][k] == sArr[i][k]{
					determine++
				}
				
			}

			if determine == len(dictionary[j]){
				breakin = true
				newArr = append(newArr, dictionary[j])
				break
			}
		}

		if breakin{
			continue
		}

		newArr = append(newArr, sArr[i])
	}

	return strings.Join(newArr, " ")

}