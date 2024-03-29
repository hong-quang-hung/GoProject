package medium

import "fmt"

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func init() {
	Solutions[714] = func() {
		fmt.Println(`Input: prices = [1,3,2,8,4,9], fee = 2`)
		fmt.Println(`Output:`, maxProfit3([]int{1, 3, 2, 8, 4, 9}, 2))
		fmt.Println(`Input: prices = [1,3,7,5,10,3], fee = 3`)
		fmt.Println(`Output:`, maxProfit3([]int{1, 3, 7, 5, 10, 3}, 3))
	}
}

func maxProfit3(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{-1, -1}
	}
	return maxProfit3Solved(prices, dp, 0, 0, fee)
}

func maxProfit3Solved(prices []int, dp [][]int, i, state int, fee int) int {
	if i >= len(prices) {
		return 0
	}

	if dp[i][state] != -1 {
		return dp[i][state]
	}

	if state == 0 {
		dp[i][state] = max(maxProfit3Solved(prices, dp, i+1, 1, fee)-prices[i], maxProfit3Solved(prices, dp, i+1, state, fee))
	} else {
		dp[i][state] = max(maxProfit3Solved(prices, dp, i+1, 0, fee)+prices[i]-fee, maxProfit3Solved(prices, dp, i+1, state, fee))
	}
	return dp[i][state]
}
