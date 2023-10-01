package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/132-pattern/
func init() {
	Solutions[456] = func() {
		fmt.Println("Input: nums = [1,2,3,4]")
		fmt.Println("Output:", find132pattern([]int{1, 2, 3, 4}))
		fmt.Println("Input: nums = [3,1,4,2]")
		fmt.Println("Output:", find132pattern([]int{3, 1, 4, 2}))
		fmt.Println("Input: nums = [-1,3,2,0]")
		fmt.Println("Output:", find132pattern([]int{-1, 3, 2, 0}))
	}
}

func find132pattern(nums []int) bool {
	num, stack := -1_000_000_001, []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i]
		if val < num {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < val {
			num, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}
	return false
}
