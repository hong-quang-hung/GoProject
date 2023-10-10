package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func init() {
	Solutions[122] = func() {
		fmt.Println("Input: prices = [7,1,5,3,6,4]")
		fmt.Println("Output:", maxProfit([]int{7, 1, 5, 3, 6, 4}))
		fmt.Println("Input: prices = [1,2,3,4,5]")
		fmt.Println("Output:", maxProfit([]int{1, 2, 3, 4, 5}))
		fmt.Println("Input: prices = [7,6,4,3,1]")
		fmt.Println("Output:", maxProfit([]int{7, 6, 4, 3, 1}))
	}
}

func maxProfit(prices []int) int {
	n := len(prices)
	res := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
