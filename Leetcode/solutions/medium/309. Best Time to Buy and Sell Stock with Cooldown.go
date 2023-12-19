package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func init() {
	Solutions[309] = func() {
		fmt.Println(`Input: prices = [1,2,3,0,2]`)
		fmt.Println(`Output:`, maxProfit2([]int{1, 2, 3, 0, 2}))
		fmt.Println(`Input: prices = [1,2,4]`)
		fmt.Println(`Output:`, maxProfit2([]int{1, 2, 4}))
		fmt.Println(`Input: prices = [2,1,4]`)
		fmt.Println(`Output:`, maxProfit2([]int{2, 1, 4}))
	}
}

func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{math.MinInt, math.MinInt}
	}
	return maxProfit2Solved(prices, dp, 0, 0, true)
}

func maxProfit2Solved(prices []int, dp [][]int, i, state int, buy bool) int {
	if i >= len(prices) {
		return 0
	}

	if dp[i][state] != math.MinInt {
		return dp[i][state]
	}

	if buy {
		dp[i][state] = max(maxProfit2Solved(prices, dp, i+1, 1, false)-prices[i], maxProfit2Solved(prices, dp, i+1, state, true))
	} else {
		dp[i][state] = max(maxProfit2Solved(prices, dp, i+2, 0, true)+prices[i], maxProfit2Solved(prices, dp, i+1, state, false))
	}
	return dp[i][state]
}
