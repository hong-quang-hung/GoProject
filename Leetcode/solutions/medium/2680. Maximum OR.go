package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-or/
func init() {
	Solutions[2680] = func() {
		fmt.Println(`Input: nums = [12,9], k = 1`)
		fmt.Println(`Output:`, maximumOr([]int{12, 9}, 1))
		fmt.Println(`Input: nums = [8,1,2], k = 2`)
		fmt.Println(`Output:`, maximumOr([]int{8, 1, 2}, 2))
		fmt.Println(`Input: nums = [10,8,4], k = 1`)
		fmt.Println(`Output:`, maximumOr([]int{10, 8, 4}, 1))
	}
}

func maximumOr(nums []int, k int) int64 {
	res, n := int64(0), len(nums)
	pre, suf := make([]int64, n+1), make([]int64, n+1)
	pre[0] = 0
	suf[n] = 0
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] | int64(nums[i])
	}

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | int64(nums[i])
	}

	p := int64(1) << k
	for i := 0; i < n; i++ {
		res = max(res, pre[i]|(int64(nums[i])*p)|suf[i+1])
	}
	return res
}
