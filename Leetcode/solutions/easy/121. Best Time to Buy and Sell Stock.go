package easy

import "fmt"

// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func LeetCode_Max_Profit() {
	fmt.Println("Input: x = [2,1,2,1,0,1,2]")
	fmt.Println("Output:", maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func maxProfit(prices []int) int {
	maxPrices, minPrices := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrices > prices[i] {
			minPrices = prices[i]
		} else if maxPrices < prices[i]-minPrices {
			maxPrices = prices[i] - minPrices
		}
	}
	return maxPrices
}
