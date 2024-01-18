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
	if n <= 2 {
		return n
	}
	// To avoid using a slice, use two pointers to keep track
	// of the previous two numbers.
	vals := make([]int, n)
	vals[0] = 1
	vals[1] = 2
	for i := 2; i < n; i++ {
		vals[i] = vals[i-1] + vals[i-2]
	}
	return vals[n-1]
}
