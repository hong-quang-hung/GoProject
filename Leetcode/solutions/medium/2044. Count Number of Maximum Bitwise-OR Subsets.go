package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
func Leetcode_Count_Max_Or_Subsets() {
	fmt.Println("Input: nums = [3,1]")
	fmt.Println("Output:", countMaxOrSubsets([]int{3, 1}))
	fmt.Println("Input: nums = [2,2,2]")
	fmt.Println("Output:", countMaxOrSubsets([]int{2, 2, 2}))
	fmt.Println("Input: nums = [10,8,4]")
	fmt.Println("Output:", countMaxOrSubsets([]int{10, 8, 4}))
}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	pre[0] = 0
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] | nums[i]
	}

	maxOr, res := pre[n], 0
	countMaxOrSubsetsBacktrack(nums, &res, 0, maxOr, 0)
	return res
}

func countMaxOrSubsetsBacktrack(nums []int, res *int, current int, maxOr int, i int) {
	if current == maxOr {
		*res++
	}

	temp := current
	for j := i; j < len(nums); j++ {
		current = current | nums[j]
		countMaxOrSubsetsBacktrack(nums, res, current, maxOr, j+1)
		current = temp
	}
}
