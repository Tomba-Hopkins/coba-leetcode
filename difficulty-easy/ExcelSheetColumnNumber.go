package difficultyeasy

// Question :

// Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.
// For example:
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
// Example 1:
// Input: columnTitle = "A"
// Output: 1
// Example 2:
// Input: columnTitle = "AB"
// Output: 28
// Example 3:
// Input: columnTitle = "ZY"
// Output: 701
// Constraints:
// 1 <= columnTitle.length <= 7
// columnTitle consists only of uppercase English letters.
// columnTitle is in the range ["A", "FXSHRXW"].

// Answer :
func ExcelSheetColumnNumber(columnTitle string) (r int) {
	x, m := len(columnTitle) - 1, map[string]int{}
	for c := 'A'; c <= 'Z'; c++{
		m[string(c)] = int(c - 'A') + 1
	}
	for i := 0; i < len(columnTitle); i++{
		d := 1; // d := 26 * x -> awalnya pake gini salah
		for j := 0; j < x; j++{
			d *= 26
		}
		r += m[string(columnTitle[i])] * d
		x--	
	}
	return
}