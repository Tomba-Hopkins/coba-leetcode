package weeklycontest

import "strconv"

func countMonobit(n int) (r int) {
	for i := 0; i <= n; i++ {
		s, x, b := strconv.FormatInt(int64(i), 2), "", true
		for j := 0; j < len(s); j++ {
			if j == 0 {
				x += string(s[j])
			}
			if string(s[j]) != x {
				b = false
			}
		}
		if b {
			r++
		}
	}
	return
}
