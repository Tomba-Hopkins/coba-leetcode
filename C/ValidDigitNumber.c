#include <stdbool.h>
#include <stdio.h>
#include <string.h>
// desc

// You are given an integer n and a digit x.
// A number is considered valid if:
// It contains at least one occurrence of digit x, and
// It does not start with digit x.
// Return true if n is valid, otherwise return false.
// Example 1:
// Input: n = 101, x = 0
// Output: true
// Explanation:
// The number contains digit 0 at index 1. It does not start with 0, so it satisfies both conditions. Thus, the answer is true​​​​​​​.
// Example 2:
// Input: n = 232, x = 2
// Output: false
// Explanation:
// The number starts with 2, which violates the condition. Thus, the answer is false.
// Example 3:
// Input: n = 5, x = 1
// Output: false
// Explanation:
// The number does not contain digit 1. Thus, the answer is false.
// Constraints:
// 0 <= n <= 105​​​​​​​
// 0 <= x <= 9

// ans


bool ValidDigitNumber(int n, int x) {
    char s1[20];
    char s2[20];
    sprintf(s1, "%d", n);
    sprintf(s2, "%d", x);
    if (s1[0] == s2[0]) {
        return false;
    }
    return strstr(s1, s2) != NULL;
}