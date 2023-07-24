package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-longest-increasing-subsequence/
func init() {
	Solutions[673] = func() {
		fmt.Println("Input: nums = [1,3,5,4,7]")
		fmt.Println("Output:", findNumberOfLIS([]int{1, 3, 5, 4, 7}))

		fmt.Println("Input: nums = [2,2,2,2,2]")
		fmt.Println("Output:", findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	}
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n)
	res, maxLength := 0, 0
	for i := 0; i < n; i++ {
		findNumberOfLISSolved(nums, dp, i)
		maxLength = max(maxLength, dp[i][0])
	}

	for i := 0; i < n; i++ {
		if dp[i][0] == maxLength {
			res += dp[i][1]
		}
	}
	return res
}

func findNumberOfLISSolved(nums []int, dp [][2]int, i int) {
	if dp[i][0] != 0 {
		return
	}

	dp[i][0] = 1
	dp[i][1] = 1

	for j := 0; j < i; j++ {
		if nums[j] < nums[i] {
			findNumberOfLISSolved(nums, dp, j)

			if dp[j][0]+1 > dp[i][0] {
				dp[i][0] = dp[j][0] + 1
				dp[i][1] = 0
			}

			if dp[j][0]+1 == dp[i][0] {
				dp[i][1] += dp[j][1]
			}
		}
	}
}
