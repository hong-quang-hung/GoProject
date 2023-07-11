package easy

import "fmt"

// Reference: https://leetcode.com/problems/binary-search/
func init() {
	Solutions[704] = func() {
		fmt.Println("Input: nums = [-1,0,3,5,9,12], target = 9")
		fmt.Println("Output:", search([]int{-1, 0, 3, 5, 9, 12}, 9))
	}
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (r + l) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
