package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/delete-and-earn/
func init() {
	Solutions[740] = func() {
		fmt.Println(`Input: nums = [3,4,2]`)
		fmt.Println(`Output:`, deleteAndEarn([]int{3, 4, 2}))
		fmt.Println(`Input: nums = [2,2,3,3,3,4]`)
		fmt.Println(`Output:`, deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
	}
}

func deleteAndEarn(nums []int) int {
	m := make([]int, 10001)
	for _, num := range nums {
		m[num]++
	}

	dp := make([]int, 10003)
	for i := 3; i < 10003; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+(i-2)*m[i-2])
	}
	return dp[10002]
}
