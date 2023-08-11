package medium

import "fmt"

// Reference: https://leetcode.com/problems/coin-change-ii/
func init() {
	Solutions[518] = func() {
		fmt.Println("Input: amount = 5, coins = [1,2,5]")
		fmt.Println("Output:", change(5, []int{1, 2, 5}))
		fmt.Println("Input: amount = 3, coins = [2]")
		fmt.Println("Output:", change(3, []int{2}))
		fmt.Println("Input: amount = 10, coins = [10]")
		fmt.Println("Output:", change(10, []int{10}))
	}
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[i][j] = dp[i+1][j] + dp[i][j-coins[i]]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][amount]
}
