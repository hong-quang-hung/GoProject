package medium

import "fmt"

// Reference: https://leetcode.com/problems/ugly-number-ii/
func init() {
	Solutions[264] = func() {
		fmt.Println("Input: n = 10")
		fmt.Println("Output:", nthUglyNumber(10))
		fmt.Println("Input: n = 1000")
		fmt.Println("Output:", nthUglyNumber(1000))
	}
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	i, x, y, z := 2, 1, 1, 1
	for i <= n {
		dp[i] = min(min(dp[x]*2, dp[y]*3), dp[z]*5)
		if dp[i] == dp[x]*2 {
			x++
		}
		if dp[i] == dp[y]*3 {
			y++
		}
		if dp[i] == dp[z]*5 {
			z++
		}
		i++
	}
	return dp[n]
}
