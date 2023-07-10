package medium

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2592] = Leetcode_Maximize_Greatness
}

// Reference: https://leetcode.com/problems/maximize-greatness-of-an-array/
func Leetcode_Maximize_Greatness() {
	fmt.Println("Input: nums = [42,8,75,28,35,21,13,21]")
	fmt.Println("Output:", maximizeGreatness([]int{42, 8, 75, 28, 35, 21, 13, 21}))
}

func maximizeGreatness(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res++
		}
	}
	return res
}
