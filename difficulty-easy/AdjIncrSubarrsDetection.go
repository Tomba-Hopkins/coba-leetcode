package difficultyeasy

// Question :

// Answer :
// gpt fixed
func AdjIncrSubarrsDetection(nums []int, k int) bool {
	for i := 0; i + k*2 <= len(nums); i++{
		x, b := nums[i:i+k], true
		for j := 0; j < k - 1 ;j++{
			if x[j + 1] <= x[j] {
				b = false
				break
			}
		}
		if !b {
			continue
		}
		y, c := nums[i+k:i+k*2], true
		for j := 0; j < k - 1; j++{
			if y[j + 1] <= y[j] {
				c = false
				break
			}
		}
		if b && c {
			return  true
		}
	}
	return false
}

// my wrong ans
// func AdjIncrSubarrsDetection(nums []int, k int) bool {
// 	d := 0
// 	for i := 0; i < len(nums); i++{
// 		var x []int
// 		if i + k <= len(nums) {
// 			x = nums[i:i+k]
// 		} else {
// 			break
// 		}
// 		b := true
// 		for j := 0; j < len(x) - 1; j++{
// 			if x[j + 1] <= x[j] {
// 				b = false
// 			}
// 		}
// 		if b {
// 			d++
// 			i += k - 1
// 		}
		
// 	}
// 	return d >= 2
// }