package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-pivot-integer/
func init() {
	Solutions[2485] = func() {
		fmt.Println(`Input: n = 8`)
		fmt.Println(`Output:`, pivotInteger(8))
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, pivotInteger(1))
		fmt.Println(`Input: n = 4`)
		fmt.Println(`Output:`, pivotInteger(4))
	}
}

func pivotInteger(n int) int {
	y := (n + 1) * n / 2
	for i := 1; i <= n; i++ {
		if (i+1)*i == y+i {
			return i
		}
	}
	return -1
}
