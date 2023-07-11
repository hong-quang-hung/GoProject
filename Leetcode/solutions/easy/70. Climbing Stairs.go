package easy

import "fmt"

// Reference: https://leetcode.com/problems/climbing-stairs/
func init() {
	Solutions[70] = func() {
		fmt.Println("Input: n = 2")
		fmt.Println("Output:", climbStairs(2))
		fmt.Println("Input: n = 3")
		fmt.Println("Output:", climbStairs(3))
	}
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return climbStairsSolved(n, dp, 0)
}

func climbStairsSolved(n int, dp []int, i int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = climbStairsSolved(n, dp, i+1) + climbStairsSolved(n, dp, i+2)
	return dp[i]
}
