package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/minimum-score-by-changing-two-elements/
func init() {
	Solutions[2567] = func() {
		fmt.Println("Input: nums = [1,4,3]")
		fmt.Println("Output:", minimizeSum([]int{1, 4, 3}))
		fmt.Println("Input: nums = [1,4,7,8,5]")
		fmt.Println("Output:", minimizeSum([]int{1, 4, 7, 8, 5}))
	}
}

func minimizeSum(nums []int) int {
	n := len(nums)
	if n == 3 {
		return 0
	}

	sort.Ints(nums)
	return min(nums[n-1]-nums[2], min(nums[n-3]-nums[0], nums[n-2]-nums[1]))
}
