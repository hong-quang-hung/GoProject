package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
func init() {
	Solutions[123] = func() {
		fmt.Println(`Input: prices = [3,3,5,0,0,3,1,4]`)
		fmt.Println(`Output:`, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
		fmt.Println(`Input: prices = [1,2,3,4,5]`)
		fmt.Println(`Output:`, maxProfit([]int{1, 2, 3, 4, 5}))
		fmt.Println(`Input: prices = [7,6,4,3,1]`)
		fmt.Println(`Output:`, maxProfit([]int{7, 6, 4, 3, 1}))
	}
}

func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 0; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
