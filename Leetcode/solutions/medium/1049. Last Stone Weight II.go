package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/last-stone-weight-ii/
func init() {
	Solutions[1049] = func() {
		fmt.Println(`Input: stones = [2,7,4,1,8,1]`)
		fmt.Println(`Output:`, lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
		fmt.Println(`Input: stones = [31,26,33,21,40]`)
		fmt.Println(`Output:`, lastStoneWeightII([]int{31, 26, 33, 21, 40}))
	}
}

func lastStoneWeightII(stones []int) int {
	sum, n := 0, len(stones)
	for _, stone := range stones {
		sum += stone
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, sum)
		for j := 0; j < sum; j++ {
			dp[i][j] = -1
		}
	}
	return lastStoneWeightIISolved(stones, dp, 0, 0, 0)
}

func lastStoneWeightIISolved(stones []int, dp [][]int, i int, sumL int, sumR int) int {
	if i == len(stones) {
		return int(math.Abs(float64(sumL) - float64(sumR)))
	}

	if dp[i][sumL] != -1 {
		return dp[i][sumL]
	}

	if dp[i][sumR] != -1 {
		return dp[i][sumR]
	}

	dp[i][sumL] = min(lastStoneWeightIISolved(stones, dp, i+1, sumL+stones[i], sumR), lastStoneWeightIISolved(stones, dp, i+1, sumL, sumR+stones[i]))
	return dp[i][sumL]
}
