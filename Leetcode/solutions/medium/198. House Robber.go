package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/house-robber/
func Leetcode_House_Robber() {
	fmt.Println("Input: nums = [1,2,3,1]")
	fmt.Println("Output:", rob([]int{1, 2, 3, 1}))
	fmt.Println("Input: nums = [2,7,9,3,1]")
	fmt.Println("Output:", rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	return robSolve(nums, dp, n-1)
}

func robSolve(nums []int, dp []int, i int) int {
	if i < 0 {
		return 0
	}

	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max(nums[i]+robSolve(nums, dp, i-2), robSolve(nums, dp, i-1))
	return dp[i]
}
