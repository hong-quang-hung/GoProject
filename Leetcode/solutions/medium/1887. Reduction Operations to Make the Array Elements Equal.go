package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
func init() {
	Solutions[1887] = func() {
		fmt.Println("Input: nums = [5,1,3]")
		fmt.Println("Output:", reductionOperations([]int{5, 1, 3}))
		fmt.Println("Input: nums = [1,1,1]")
		fmt.Println("Output:", reductionOperations([]int{1, 1, 1}))
		fmt.Println("Input: nums = [1,1,2,2,3]")
		fmt.Println("Output:", reductionOperations([]int{1, 1, 2, 2, 3}))
	}
}

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	peak := nums[len(nums)-1]
	res := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < peak {
			peak = nums[i]
			res += len(nums) - i - 1
		}
	}
	return res
}
