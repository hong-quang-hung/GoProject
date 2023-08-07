package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/partition-equal-subset-sum/
func init() {
	Solutions[416] = func() {
		fmt.Println("Input: nums = [1,5,11,5]")
		fmt.Println("Output:", canPartition([]int{1, 5, 11, 5}))
		fmt.Println("Input: nums = [1,2,3,5]")
		fmt.Println("Output:", canPartition([]int{1, 2, 3, 5}))
	}
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	sum = sum >> 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j-nums[i]] || dp[j]
			}
		}
	}
	return dp[sum]
}
