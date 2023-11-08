package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/
func init() {
	Solutions[2849] = func() {
		fmt.Println("Input: sx = 2, sy = 4, fx = 7, fy = 7, t = 6")
		fmt.Println("Output:", isReachableAtTime(2, 4, 7, 7, 6))
		fmt.Println("Input: sx = 1, sy = 1, fx = 2, fy = 2, t = 1")
		fmt.Println("Output:", isReachableAtTime(1, 1, 2, 2, 1))
		fmt.Println("Input: sx = 1, sy = 2, fx = 1, fy = 2, t = 1")
		fmt.Println("Output:", isReachableAtTime(1, 2, 1, 2, 1))
	}
}

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	d := max(abs(sx-fx), abs(sy-fy))
	if d == 0 {
		return t != 1
	}
	return t >= d
}
