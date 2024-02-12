package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-swaps-to-make-strings-equal/
func init() {
	Solutions[1247] = func() {
		fmt.Println(`Input: s1 = "xx", s2 = "yy"`)
		fmt.Println(`Output:`, minimumSwap("xx", "yy"))
		fmt.Println(`Input: s1 = "xy", s2 = "yx"`)
		fmt.Println(`Output:`, minimumSwap("xy", "yx"))
		fmt.Println(`Input: s1 = "xx", s2 = "yx"`)
		fmt.Println(`Output:`, minimumSwap("xx", "yx"))
	}
}

func minimumSwap(s1 string, s2 string) int {
	n := len(s1)
	x, y := 0, 0
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				x++
			} else {
				y++
			}
		}
	}

	if (x+y)%2 == 1 {
		return -1
	}
	return x/2 + y/2 + x%2 + y%2
}
