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
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	var f func(i int) int
	f = func(i int) int {
		if i > n {
			return 0
		}

		if i == n {
			return 1
		}

		if dp[i] != -1 {
			return dp[i]
		}

		dp[i] = f(i+1) + f(i+2)
		return dp[i]
	}
	return f(0)
}
