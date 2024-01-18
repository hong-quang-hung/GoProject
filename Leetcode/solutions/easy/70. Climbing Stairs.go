package easy

import "fmt"

// Reference: https://leetcode.com/problems/climbing-stairs/
func init() {
	Solutions[70] = func() {
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, climbStairs(2))
		fmt.Println(`Input: n = 3`)
		fmt.Println(`Output:`, climbStairs(3))
	}
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	res, next := 2, 3
	for i := 3; i < n+1; i++ {
		res, next = next, res+next
	}
	return res
}
