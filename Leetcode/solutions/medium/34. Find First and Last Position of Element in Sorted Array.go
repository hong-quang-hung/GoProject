package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func init() {
	Solutions[34] = func() {
		fmt.Println(`Input: nums = [5,7,7,8,8,10], target = 8`)
		fmt.Println(`Output:`, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
		fmt.Println(`Input: nums = [5,7,7,8,8,10], target = 6`)
		fmt.Println(`Output:`, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
		fmt.Println(`Input: nums = [], target = 0`)
		fmt.Println(`Output:`, searchRange([]int{}, 0))
	}
}

func searchRange(nums []int, target int) []int {
	start := sort.SearchInts(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	end := start
	for end < len(nums)-1 && nums[end+1] == target {
		end++
	}
	return []int{start, end}
}
