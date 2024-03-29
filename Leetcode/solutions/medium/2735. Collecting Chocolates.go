package medium

import "fmt"

// Reference: https://leetcode.com/problems/collecting-chocolates/
func init() {
	Solutions[2735] = func() {
		fmt.Println(`Input: nums = [20,1,15], x = 5`)
		fmt.Println(`Output:`, minCost2([]int{20, 1, 15}, 5))
		fmt.Println(`Input: nums = [1,2,3], x = 4`)
		fmt.Println(`Output:`, minCost2([]int{1, 2, 3}, 4))
	}
}

func minCost2(nums []int, x int) int64 {
	n := len(nums)
	dp := make([]int64, n)
	res := int64(0)
	for i := 0; i < n; i++ {
		dp[i] = int64(nums[i])
		res += int64(nums[i])
	}

	for i := 1; i < n; i++ {
		sum := int64(i) * int64(x)
		for j := 0; j < n; j++ {
			cur := j + i - n
			if i+j < n {
				cur = j + i
			}

			dp[j] = min(dp[j], int64(nums[cur]))
			sum += int64(dp[j])
		}
		res = min(res, sum)
	}
	return res
}
