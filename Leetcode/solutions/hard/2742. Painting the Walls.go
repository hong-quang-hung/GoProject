package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/painting-the-walls/
func Leetcode_Special_Perm() {
	fmt.Println("Input: cost = [1,2,3,2], time = [1,2,3,2]")
	fmt.Println("Output:", paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
	fmt.Println("Input: cost = [2,3,4,2], time = [1,1,1,1]")
	fmt.Println("Output:", paintWalls([]int{2, 3, 4, 2}, []int{1, 1, 1, 1}))
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 500_000_000
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			dp[j] = min(dp[j], dp[max(j-time[i]-1, 0)]+cost[i])
		}
	}
	return dp[n]
}
