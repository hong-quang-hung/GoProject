package easy

import "fmt"

// Reference: https://leetcode.com/problems/sqrtx/
func init() {
	Solutions[69] = func() {
		fmt.Println("Input: x = 4")
		fmt.Println("Output:", mySqrt(4))
		fmt.Println("Input: x = 8")
		fmt.Println("Output:", mySqrt(8))
	}
}

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		m := (l + r + 1) / 2
		if m*m > x {
			r = m - 1
		} else {
			l = m
		}

	}
	return l
}
