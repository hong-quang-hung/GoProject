package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/coin-change/
func init() {
	Solutions[322] = func() {
		fmt.Println("Input: coins = [1,2,5], amount = 11")
		fmt.Println("Output:", coinChange([]int{1, 2, 5}, 11))
		fmt.Println("Input: coins = [2], amount = 3")
		fmt.Println("Output:", coinChange([]int{2}, 3))
		fmt.Println("Input: coins = [1], amount = 0")
		fmt.Println("Output:", coinChange([]int{1}, 0))
	}
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		leastCoins := math.MaxInt
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != -1 {
				leastCoins = min(leastCoins, dp[i-coin]+1)
			}
		}

		if leastCoins != math.MaxInt {
			dp[i] = leastCoins
		}
	}
	return dp[amount]
}
