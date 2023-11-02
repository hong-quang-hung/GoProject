package medium

import "fmt"

// Reference: https://leetcode.com/problems/divide-two-integers/
func init() {
	Solutions[29] = func() {
		fmt.Println("Input: dividend = 10, divisor = 3")
		fmt.Println("Output:", divide(10, 3))
		fmt.Println("Input: dividend = 7, divisor = -3")
		fmt.Println("Output:", divide(7, -3))
	}
}

func divide(dividend int, divisor int) int {
	return max(min(dividend/divisor, 1<<31-1), -(1 << 31))
}
