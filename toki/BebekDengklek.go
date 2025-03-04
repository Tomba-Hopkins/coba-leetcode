package toki

import (
	"fmt"
)

func BebekDengklek(umurBebek []int) {

	t := []int{}
	for i := 0; i < len(umurBebek); i++{
		
		
		for j := 0; j < len(t); j++{
			t[j]++
		}
		
		
		t = append(t, umurBebek[i])

		min, max, sum := t[0], t[0], 0
		
		for _, r := range t {
			if r > max {
				max = r
			}
			if r < min {
				min = r
			}
			sum += r
		}


		med := float64(sum) / float64(len(t))


		fmt.Printf("%d %d %.4f\n", min, max, med)
	}
	
	
}

// func main() {

// 	var N int
// 	fmt.Scan(&N)

// 	umurBebek := make([]int, N)

// 	for i := 0; i < N; i++{
// 		fmt.Scan(&umurBebek[i])
// 	}

// 	BebekDengklek(umurBebek)

// }
