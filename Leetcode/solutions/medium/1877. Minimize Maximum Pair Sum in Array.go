package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
func init() {
	Solutions[1877] = func() {
		fmt.Println("Input: nums = [3,5,2,3]")
		fmt.Println("Output:", minPairSum([]int{3, 5, 2, 3}))
		fmt.Println("Input: nums = [3,5,4,2,4,6]")
		fmt.Println("Output:", minPairSum([]int{3, 5, 4, 2, 4, 6}))
	}
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for i := 0; i < n/2; i++ {
		res = max(res, nums[i]+nums[n-1-i])
	}
	return res
}
