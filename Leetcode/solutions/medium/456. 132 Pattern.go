package medium

import "fmt"

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
	return false
}
