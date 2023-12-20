package medium

import "fmt"

// Reference: https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/
func init() {
	Solutions[2369] = func() {
		fmt.Println(`Input: nums = [4,4,4,5,6]`)
		fmt.Println(`Output:`, validPartition([]int{4, 4, 4, 5, 6}))
		fmt.Println(`Input: nums = [1,1,1,2]`)
		fmt.Println(`Output:`, validPartition([]int{1, 1, 1, 2}))
	}
}

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		idx := i + 1
		if i > 0 && nums[i] == nums[i-1] {
			dp[idx] = dp[idx] || dp[idx-2]
		}

		if i > 1 && nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			dp[idx] = dp[idx] || dp[idx-3]
		}

		if i > 1 && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			dp[idx] = dp[idx] || dp[idx-3]
		}
	}
	return dp[n]
}
