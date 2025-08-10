#include <stdio.h>
#include <stdbool.h>


// desc :

// You are given an integer array nums.
// Return true if the frequency of any element of the array is prime, otherwise, return false.
// The frequency of an element x is the number of times it occurs in the array.
// A prime number is a natural number greater than 1 with only two factors, 1 and itself.
// Example 1:
// Input: nums = [1,2,3,4,5,4]
// Output: true
// Explanation:
// 4 has a frequency of two, which is a prime number.
// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: false
// Explanation:
// All elements have a frequency of one.
// Example 3:
// Input: nums = [2,2,2,4,4]
// Output: true
// Explanation:
// Both 2 and 4 have a prime frequency.
// Constraints:
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100

// answer :


bool isItPrime(int n){
    if(n <= 1){
        return false;
    } else {
        for(int i = 2; i * i <= n; i++){
            if(n % i == 0){
                return false;
            }
        }
    }
    return true;
    
}


bool checkPrimeFrequency(int* nums, int numsSize) {
    for (int i = 0; i < numsSize; i++){
        int t = 1;
        for(int j = 0; j < numsSize; j++){
            if (j != i) {
                if(nums[i] == nums[j]){
                    t++;
                }
            }
        }
        if (isItPrime(t)){
            return true;
        }
        
    }
    return false;
}



int main(){

    int nums[] = {1,2,3,4,5,4};
    int size = sizeof(nums) / sizeof(nums[0]);

    printf("%s", checkPrimeFrequency(nums, size) ? "True" : "False");
    

}