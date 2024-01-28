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
	m1 := make([]int, 2)
	for i := range s1 {
		m1[int(s1[i]-'x')]++
	}

	fmt.Println(m1)

	for i := range s2 {
		_ = i
	}
	return 0
}
