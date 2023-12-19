package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-total-number-of-colored-cells/
func init() {
	Solutions[2579] = func() {
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, coloredCells(1))
	}
}

func coloredCells(n int) int64 {
	return int64(2*n*n - 2*n + 1)
}
