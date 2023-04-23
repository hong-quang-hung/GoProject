package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
func Leetcode_Min_Operations_Equal_1() {
	fmt.Println("Input: nums = [2,6,3,4]")
	fmt.Println("Output:", minOperations_ii([]int{2, 6, 3, 4}))
	fmt.Println("Input: nums = [2,10,6,14]")
	fmt.Println("Output:", minOperations_ii([]int{2, 10, 6, 14}))
	fmt.Println("Input: nums = [6,10,15]")
	fmt.Println("Output:", minOperations_ii([]int{6, 10, 15}))
}

func minOperations_ii(nums []int) int {
	res, n := 0, len(nums)
	return n - res
}
