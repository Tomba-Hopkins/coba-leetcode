#include <stdio.h>
#include <string.h>
#include <stdlib.h>
// desc :

// Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.
// Note that after backspacing an empty text, the text will continue empty.
// Example 1:
// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".
// Example 2:
// Input: s = "ab##", t = "c#d#"
// Output: true
// Explanation: Both s and t become "".
// Example 3:
// Input: s = "a#c", t = "b"
// Output: false
// Explanation: s becomes "c" while t becomes "b".
// Constraints:
// 1 <= s.length, t.length <= 200
// s and t only contain lowercase letters and '#' characters.
// Follow up: Can you solve it in O(n) time and O(1) space?

// answer :

char* process(const char* str){
    int len = strlen(str);
    char* res = (char*)malloc(len + 1); // alokasi memori ke str.length + 1, biar cukup lah y kek make([]byte, panjang len)
    int dex = 0;
    for(int i = 0; i < len; i++){
        if (str[i] == '#'){
            if(dex > 0){
                dex--;
            }
        } else {
            res[dex++] = str[i]; // res[dex] = str[i] -> lalu dex++

            // sama kek gini
            // res[dex] = str[i];
            // dex++;
        }
    }
    res[dex] = '\0'; // tanda mengakhiri string, biar ga ketimpa ato gmn gapaham
    return res;

}

bool backspaceCompare(const char* s, const char* t) {
    char* x = process(s);
    char* y = process(t);
    bool result = strcmp(x, y) == 0;
    free(x); // hapus x dari memori,kan alloc tadi nambah ngatur memori, ini hapus
    free(y);
    return result;
}



int main(){


    printf("%s", backspaceCompare("ab#c", "ad#c") ? "True" : "False"); //mustbe true

}