package difficultyhard

import (
	"fmt"
	"sort"
)

// belum kelar
func maxScoreWords(words []string, letters []byte, score []int) int {

    b := map[byte]int{}
    m := map[string]int{}
    q := map[byte]int{}

    for _, l := range letters {
        q[l]++
    }



    newLetter := []string{}

    for _, nl := range letters {
        b[nl] = 1
    }
    for _, nl := range letters {
        if b[nl] > 0 {
            newLetter = append(newLetter, string(nl))
            b[nl]--
        }
    }

    sort.Strings(newLetter)

    fmt.Println(newLetter)
    o := 0
    
 

    for _, s := range score {
        if o == len(newLetter) {
            break
        }
        det := newLetter[o]
        if s > 0 {
            m[string(det)] = s
            o++
        }
    }



   t := []int{}

   for i := 0; i < len(words); i++{

    no := false


    g := map[rune]int{}
    for _, n := range words[i] {
        g[n]++
    }

    for j := 0; j < len(words[i]); j++{
        if q[words[i][j]] < g[rune(words[i][j])] {
            no = true
        }
    }

    if !no {
        z := 0
        for _, j := range words[i] {
            z += m[string(j)]

        }
        t = append(t, z)
    }
    
    
   }


   sort.Sort(sort.Reverse(sort.IntSlice(t)))



    
   if len(t) > 2 {
        return t[0] + t[1]
   }



    return 0
}