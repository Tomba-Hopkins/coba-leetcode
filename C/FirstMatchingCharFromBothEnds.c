#include <stdio.h>
#include <string.h>


// desc

// You are given a string s of length n consisting of lowercase English letters.
// Return the smallest index i such that s[i] == s[n - i - 1].
// If no such index exists, return -1.
//
// Example 1:
// Input: s = "abcacbd"
// Output: 1
// Explanation:
// At index i = 1, s[1] and s[5] are both 'b'.
// No smaller index satisfies the condition, so the answer is 1.
// Example 2:
// Input: s = "abc"
// Output: 1
// Explanation:
// ​​​​​​​At index i = 1, the two compared positions coincide, so both characters are 'b'.
// No smaller index satisfies the condition, so the answer is 1.
// Example 3:
// Input: s = "abcdab"
// Output: -1
// Explanation:
// ​​​​​​​For every index i, the characters at positions i and n - i - 1 are different.
// Therefore, no valid index exists, so the answer is -1.
//
// Constraints:
// 1 <= n == s.length <= 100
// s consists of lowercase English letters

// ans

int FirstMatchingCharFromBothEnds(char* s) {
    int n = strlen(s);
    for (int i = 0; i < strlen(s); i++){
        if (s[i] == s[n - i - 1]) {
            return i;
        }
    }
    return -1;
}