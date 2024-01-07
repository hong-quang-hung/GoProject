package hard

import "fmt"

// Reference: https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
func init() {
	Solutions[446] = func() {
		fmt.Println(`Input: nums = [2,4,6,8,10]`)
		fmt.Println(`Output:`, numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
		fmt.Println(`Input: nums = [7,7,7,7,7]`)
		fmt.Println(`Output:`, numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
	}
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += dp[j][diff] + 1
			res += dp[j][diff]
		}
	}
	return res
}
