package hard

import "fmt"

// Reference: https://leetcode.com/problems/k-inverse-pairs-array/
func init() {
	Solutions[629] = func() {
		fmt.Println(`Input: n = 3, k = 0`)
		fmt.Println(`Output:`, kInversePairs(3, 0))
		fmt.Println(`Input: n = 3, k = 1`)
		fmt.Println(`Output:`, kInversePairs(3, 1))
	}
}

func kInversePairs(n int, k int) int {
	prevDP := make([]int, k+1)
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 2; i <= n; i++ {
		prevDP, dp = dp, prevDP
		sum := 0
		for j := range dp {
			sum += prevDP[j]
			if j >= i {
				sum -= prevDP[j-i]
			}
			dp[j] = sum % mod
		}
	}
	return dp[k]
}
