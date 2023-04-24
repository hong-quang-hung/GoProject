package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func Leetcode_Max_Profit() {
	fmt.Println("Input: prices = [1,2,3,0,2]")
	fmt.Println("Output:", maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println("Input: prices = [1,2,4]")
	fmt.Println("Output:", maxProfit([]int{1, 2, 4}))
	fmt.Println("Input: prices = [2,1,4]")
	fmt.Println("Output:", maxProfit([]int{2, 1, 4}))
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{math.MinInt, math.MinInt}
	}
	return maxProfitSolved(prices, dp, 0, 0, true)
}

func maxProfitSolved(prices []int, dp [][]int, i, state int, buy bool) int {
	if i >= len(prices) {
		return 0
	}

	if dp[i][state] != math.MinInt {
		return dp[i][state]
	}

	if buy {
		dp[i][state] = max(maxProfitSolved(prices, dp, i+1, 1, false)-prices[i], maxProfitSolved(prices, dp, i+1, state, true))
	} else {
		dp[i][state] = max(maxProfitSolved(prices, dp, i+2, 0, true)+prices[i], maxProfitSolved(prices, dp, i+1, state, false))
	}
	return dp[i][state]
}
