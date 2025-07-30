#include <stdio.h>

int longestSubarray(int* nums, int numsSize){
    int x = 0, result = 0, temp = 0;
    for(int i = 0; i < numsSize; i++){
        if(nums[i] > x){
            x = nums[i];
        }
    }
    for(int i = 0; i < numsSize; i++){
        if(nums[i] == x){
            temp++;
        } else {
            if (temp > result){
                result = temp;
            }
            temp = 0;
        }
    }
    if(temp > result){
        result = temp;
    }
    return result;
}


int main(){
   {
        int nums[] = {1,2,3,3,2,2}; // expected: 2
        int size = sizeof(nums) / sizeof(nums[0]);
        int result = longestSubarray(nums, size);
        printf("%d\n", result);
   }



    {
        int nums[] = {1,2,3,4}; // expected: 1
        int size = sizeof(nums) / sizeof(nums[0]);
        int result = longestSubarray(nums, size);
        printf("%d\n", result);
    }


}