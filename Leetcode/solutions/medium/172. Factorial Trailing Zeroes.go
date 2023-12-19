package medium

import "fmt"

// Reference: https://leetcode.com/problems/factorial-trailing-zeroes/
func init() {
	Solutions[172] = func() {
		fmt.Println(`Input: n = 3`)
		fmt.Println(`Output:`, trailingZeroes(3))
		fmt.Println(`Input: n = 5`)
		fmt.Println(`Output:`, trailingZeroes(5))
		fmt.Println(`Input: n = 30`)
		fmt.Println(`Output:`, trailingZeroes(30))
	}
}

func trailingZeroes(n int) int {
	return n/5 + n/25 + n/125 + n/625 + n/3125
}
