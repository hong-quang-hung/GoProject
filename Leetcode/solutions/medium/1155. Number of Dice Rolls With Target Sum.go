package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
func init() {
	Solutions[1155] = func() {
		fmt.Println(`Input: n = 1, k = 6, target = 3`)
		fmt.Println(`Output:`, numRollsToTarget(1, 6, 3))
		fmt.Println(`Input: n = 2, k = 6, target = 7`)
		fmt.Println(`Output:`, numRollsToTarget(2, 6, 7))
		fmt.Println(`Input: n = 30, k = 30, target = 500`)
		fmt.Println(`Output:`, numRollsToTarget(30, 30, 500))
	}
}

func numRollsToTarget(n int, k int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= k && i <= target; i++ {
		dp[i] = 1
	}

	for x := 2; x <= n; x++ {
		next := make([]int, target+1)
		for i := x; i <= x*k && i <= target; i++ {
			for j := i - 1; j >= i-k && j >= 0; j-- {
				next[i] = (next[i] + dp[j]) % mod
			}
		}
		dp = next
	}
	return dp[target]
}
