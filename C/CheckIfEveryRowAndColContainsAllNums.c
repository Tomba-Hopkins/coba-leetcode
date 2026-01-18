#include <stdio.h>

// desc

// An n x n matrix is valid if every row and every column contains all the integers from 1 to n (inclusive).
// Given an n x n integer matrix matrix, return true if the matrix is valid. Otherwise, return false.
// Example 1:
// Input: matrix = [[1,2,3],[3,1,2],[2,3,1]]
// Output: true
// Explanation: In this case, n = 3, and every row and column contains the numbers 1, 2, and 3.
// Hence, we return true.
// Example 2:
// Input: matrix = [[1,1,1],[1,2,3],[1,2,3]]
// Output: false
// Explanation: In this case, n = 3, but the first row and the first column do not contain the numbers 2 or 3.
// Hence, we return false.
// Constraints:
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// 1 <= matrix[i][j] <= n

// ans
bool CheckIfEveryRowAndColContainsAllNums(int** matrix, int matrixSize, int* matrixColSize) {
    for(int point = 1; point <= matrixSize; point++){
        for(int i = 0; i < matrixColSize[0]; i++){
            bool isRow = false;
            for(int j = 0; j < matrixSize; j++){
                if(matrix[j][i] == point) {
                    isRow = true;
                    break;
                }
            }
            if(!isRow) {
                return false;
            }
        }
        for(int i = 0; i < matrixSize; i++){
            bool isCol = false;
            for(int j = 0; j < matrixColSize[i]; j++){
                if(matrix[i][j] == point) {
                    isCol = true;
                    break;
                }
            }
            if(!isCol) {
                return false;
            }
        }
    }
    return true;
}