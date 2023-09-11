package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/find-the-duplicate-number/
func init() {
	Solutions[287] = func() {
		fmt.Println("Input: nums = [1,3,4,2,2]")
		fmt.Println("Output:", findDuplicate([]int{3, 1, 3, 4, 2}))
		fmt.Println("Input: nums = [1,3,4,2,2]")
		fmt.Println("Output:", findDuplicate([]int{1, 3, 4, 2, 2}))
	}
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}
