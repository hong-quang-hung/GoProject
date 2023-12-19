package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/movement-of-robots/
func init() {
	Solutions[2731] = func() {
		fmt.Println(`Input: nums = [-2,0,2], s = "RLL", d = 3`)
		fmt.Println(`Output:`, sumDistance([]int{-2, 0, 2}, `RLL`, 3))
		fmt.Println(`Input: nums = [1,0], s = "RL", d = 2`)
		fmt.Println(`Output:`, sumDistance([]int{1, 0}, `RL`, 2))
	}
}

func sumDistance(nums []int, s string, d int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			nums[i] += d
		} else {
			nums[i] -= d
		}
	}

	sort.Ints(nums)

	res, prefix := 0, 0
	for i := 0; i < n; i++ {
		res += i*nums[i] - prefix
		res %= mod
		prefix += nums[i]
	}
	return res
}
