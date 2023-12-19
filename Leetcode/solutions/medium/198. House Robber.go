package medium

import "fmt"

// Reference: https://leetcode.com/problems/house-robber/
func init() {
	Solutions[198] = func() {
		fmt.Println(`Input: nums = [1,2,3,1]`)
		fmt.Println(`Output:`, rob([]int{1, 2, 3, 1}))
		fmt.Println(`Input: nums = [2,7,9,3,1]`)
		fmt.Println(`Output:`, rob([]int{2, 7, 9, 3, 1}))
	}
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := 2; i < n+2; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[n+1]
}
