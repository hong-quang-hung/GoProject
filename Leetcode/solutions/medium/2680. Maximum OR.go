package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/maximum-or/
func Leetcode_Maximum_Or() {
	fmt.Println("Input: nums = [12,9], k = 1")
	fmt.Println("Output:", maximumOr([]int{12, 9}, 1))
	fmt.Println("Input: nums = [8,1,2], k = 2")
	fmt.Println("Output:", maximumOr([]int{8, 1, 2}, 2))
}

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	return maximumOrSolved(nums, dp, n-1, k)
}

func maximumOrSolved(nums []int, dp [][]int64, i, k int) int64 {
	if i < 0 || k < 0 {
		return 0
	}

	if dp[i][k] != -1 {
		return dp[i][k]
	}

	maxOr := int64(math.MinInt64)
	for j := 0; j <= k; j++ {
		sum := int64(nums[i]) * int64(math.Pow(2, float64(j)))
		maxOr = max64(maxOr, sum|maximumOrSolved(nums, dp, i-1, k-j))
	}

	dp[i][k] = maxOr
	return dp[i][k]
}
