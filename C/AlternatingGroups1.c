#include <stdio.h>

// desc :

// There is a circle of red and blue tiles. You are given an array of integers colors. The color of tile i is represented by colors[i]:
// colors[i] == 0 means that tile i is red.
// colors[i] == 1 means that tile i is blue.
// Every 3 contiguous tiles in the circle with alternating colors (the middle tile has a different color from its left and right tiles) is called an alternating group.
// Return the number of alternating groups.
// Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.
// Example 1:
// Input: colors = [1,1,1]
// Output: 0
// Explanation:
// Example 2:
// Input: colors = [0,1,0,0,1]
// Output: 3
// Explanation:
// Alternating groups:
// Constraints:
// 3 <= colors.length <= 100
// 0 <= colors[i] <= 1

// answers :

int numberOfAlternatingGroups(int* colors, int colorsSize){
    int result = 0;
    for(int i = 0; i < colorsSize; i++){
        int a = 0, b = 0;
        if(i - 1 < 0){
            a = colorsSize - 1;
        } else {
            a = i - 1;
        }
        if(i + 1 >= colorsSize){
            b = 0;
        } else {
            b = i + 1;
        }
        if (colors[i] != colors[a] && colors[i] != colors[b]){
            result++;
        }
    }
    return result;
}


int main(){
    {
        int colors[] = {1,1,1}; // must be 0
        int size = sizeof(colors) / sizeof(colors[0]);
        int result = numberOfAlternatingGroups(colors, size);
        printf("%d\n", result);
    }
    {
        int colors[] = {0,1,0,0,1}; // must be 3
        int size = sizeof(colors) / sizeof(colors[0]);
        int result = numberOfAlternatingGroups(colors, size);
        printf("%d\n", result);
    }
}