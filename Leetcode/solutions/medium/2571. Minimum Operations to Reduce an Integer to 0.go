package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/
func init() {
	Solutions[2571] = func() {
		fmt.Println(`Input: n = 39`)
		fmt.Println(`Output:`, minOperationsII(39))
		fmt.Println(`Input: n = 54`)
		fmt.Println(`Output:`, minOperationsII(54))
	}
}

func minOperationsII(n int) int {
	res := 0
	for n > 0 {
		if n%2 == 0 {
			n >>= 1
			continue
		}

		if n%4 == 1 {
			n--
		} else {
			n++
		}
		res++
	}
	return res
}
