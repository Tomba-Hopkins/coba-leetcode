package usingimportpackage

import (
	"fmt"
	"strconv"
)

// Question :

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given two integers x and y, return the Hamming distance between them.
// Example 1:
// Input: x = 1, y = 4
// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑
// The above arrows point to positions where the corresponding bits are different.
// Example 2:
// Input: x = 3, y = 1
// Output: 1
// Constraints:
// 0 <= x, y <= 231 - 1

// Answer :

func HammingDistance(x int, y int) (r int) {
    n1 := strconv.FormatInt(int64(x), 2)
    n2 := strconv.FormatInt(int64(y), 2)
    mlen := len(n1)
    if len(n2) > mlen{
        mlen = len(n2)
    }
    n1 = fmt.Sprintf("%0*s", mlen, n1)
    n2 = fmt.Sprintf("%0*s", mlen, n2)
    for i := 0; i < len(n1); i++{
        if n1[i] != n2[i] {
            r++
        }
    }
    return
}
