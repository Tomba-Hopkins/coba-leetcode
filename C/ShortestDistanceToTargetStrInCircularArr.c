#include <stdio.h>
#include <stdbool.h>
#include <string.h>


// desc

// You are given a 0-indexed circular string array words and a string target. A circular array means that the array's end connects to the array's beginning.
// Formally, the next element of words[i] is words[(i + 1) % n] and the previous element of words[i] is words[(i - 1 + n) % n], where n is the length of words.
// Starting from startIndex, you can move to either the next word or the previous word with 1 step at a time.
// Return the shortest distance needed to reach the string target. If the string target does not exist in words, return -1.
// Example 1:
// Input: words = ["hello","i","am","leetcode","hello"], target = "hello", startIndex = 1
// Output: 1
// Explanation: We start from index 1 and can reach "hello" by
// - moving 3 units to the right to reach index 4.
// - moving 2 units to the left to reach index 4.
// - moving 4 units to the right to reach index 0.
// - moving 1 unit to the left to reach index 0.
// The shortest distance to reach "hello" is 1.
// Example 2:
// Input: words = ["a","b","leetcode"], target = "leetcode", startIndex = 0
// Output: 1
// Explanation: We start from index 0 and can reach "leetcode" by
// - moving 2 units to the right to reach index 2.
// - moving 1 unit to the left to reach index 2.
// The shortest distance to reach "leetcode" is 1.
// Example 3:
// Input: words = ["i","eat","leetcode"], target = "ate", startIndex = 0
// Output: -1
// Explanation: Since "ate" does not exist in words, we return -1.
// Constraints:
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] and target consist of only lowercase English letters.
// 0 <= startIndex < words.length

// ans

int ShortestDistanceToTargetStrInCircularArr(char** words, int wordsSize, char* target, int startIndex) {
    bool d = true;
    for(int i = 0; i < wordsSize; i++){
        if(strcmp(words[i], target) == 0) {
            d = false;
            break;
        }
    }
    if(d) {
        return -1;
    }
    int x = startIndex;
    int y = startIndex;
    int xCount = 0;
    int yCount = 0;
    while (true) {
        if(strcmp(words[x], target) != 0) {
            x--;
            if(x == -1) {
                x = wordsSize - 1;
            }
            xCount++;
        } else {
            break;
        }
    }
    while (true) {
        if(strcmp(words[y], target) != 0) {
            y++;
            if(y == wordsSize) {
                y = 0;
            }
            yCount++;
        } else {
            break;
        }
    }
    if(xCount < yCount) {
        return xCount;
    }
    return yCount;
}