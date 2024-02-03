package medium

import "fmt"

// Reference: https://leetcode.com/problems/partition-array-for-maximum-sum/
func init() {
	Solutions[1043] = func() {
		fmt.Println(`Input: arr = [1,15,7,9,2,5,10], k = 3`)
		fmt.Println(`Output:`, maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
		fmt.Println(`Input: arr = [1,4,1,5,7,3,6,1,9,9,3], k = 4`)
		fmt.Println(`Output:`, maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
	}
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	m := k + 1
	dp := make([]int, m)
	for start := n - 1; start >= 0; start-- {
		curr := 0
		end := min(n, start+k)
		for i := start; i < end; i++ {
			curr = max(curr, arr[i])
			dp[start%m] = max(dp[start%m], dp[(i+1)%m]+curr*(i-start+1))
		}
	}
	return dp[0]
}
