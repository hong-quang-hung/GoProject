package medium

import "fmt"

// Reference: https://leetcode.com/problems/integer-break/
func init() {
	Solutions[343] = func() {
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, integerBreak(2))
		fmt.Println(`Input: n = 10`)
		fmt.Println(`Output:`, integerBreak(10))
	}
}

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}

	dp := make([]int, n+1)
	for i := 1; i <= 3; i++ {
		dp[i] = i
	}

	for i := 4; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], j*dp[i-j])
		}
	}
	return dp[n]
}
