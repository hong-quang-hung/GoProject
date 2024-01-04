package hard

import "fmt"

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
func init() {
	Solutions[188] = func() {
		fmt.Println(`Input: k = 2, prices = [2,4,1]`)
		fmt.Println(`Output:`, maxProfit4(2, []int{2, 4, 1}))
		fmt.Println(`Input: k = 2, prices = [3,2,6,5,0,3]`)
		fmt.Println(`Output:`, maxProfit4(2, []int{3, 2, 6, 5, 0, 3}))
	}
}

type Transaction struct {
	buy  int
	sell int
}

func maxProfit4(k int, prices []int) int {
	dp := make([]Transaction, k)
	for i := 0; i < k; i++ {
		dp[i] = Transaction{buy: -prices[0], sell: 0}
	}

	for i := 1; i < len(prices); i++ {
		dp[0].buy = max(dp[0].buy, -prices[i])
		dp[0].sell = max(dp[0].sell, dp[0].buy+prices[i])
		for j := 1; j < k; j++ {
			dp[j].buy = max(dp[j].buy, dp[j-1].sell-prices[i])
			dp[j].sell = max(dp[j].sell, dp[j].buy+prices[i])
		}
	}
	return dp[k-1].sell
}
