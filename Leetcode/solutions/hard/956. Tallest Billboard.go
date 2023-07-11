package hard

import "fmt"

// Reference: https://leetcode.com/problems/tallest-billboard/
func init() {
	Solutions[956] = func() {
		fmt.Println("Input: nums = [1, 2, 3, 6]")
		fmt.Println("Output:", tallestBillboard([]int{1, 2, 3, 6}))
		fmt.Println("Input: nums = [1,2,3,4,5,6]")
		fmt.Println("Output:", tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	}
}

func tallestBillboard(rods []int) int {
	sum := 0
	for _, rod := range rods {
		sum += rod
	}

	dp := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		dp[i] = -1
	}

	for _, rod := range rods {
		temp := make([]int, sum+1)
		copy(temp, dp)
		for diff := 0; diff <= sum-rod; diff++ {
			if temp[diff] >= 0 {
				dp[diff+rod] = max(dp[diff+rod], temp[diff])
				idx := abs(diff - rod)
				dp[idx] = max(dp[idx], temp[diff]+min(diff, rod))
			}
		}
	}
	return dp[0]
}
