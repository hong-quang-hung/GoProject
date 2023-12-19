package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/maximum-subarray/
func init() {
	Solutions[53] = func() {
		fmt.Println(`Input: nums = [-2,1,-3,4,-1,2,1,-5,4]`)
		fmt.Println(`Output:`, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
		fmt.Println(`Input: nums = [5,4,-1,7,8]`)
		fmt.Println(`Output:`, maxSubArray([]int{5, 4, -1, 7, 8}))
	}
}

func maxSubArray(nums []int) int {
	res, n, sum := nums[0], len(nums), nums[0]
	for i := 1; i < n; i++ {
		sum = max(nums[i], sum+nums[i])
		res = max(res, sum)
	}
	return res
}
