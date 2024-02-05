package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/type-of-triangle-ii/
func init() {
	Solutions[3024] = func() {
		fmt.Println(`Input: nums = [3,3,3]`)
		fmt.Println(`Output:`, triangleType([]int{3, 3, 3}))
		fmt.Println(`Input: nums = [3,4,5]`)
		fmt.Println(`Output:`, triangleType([]int{3, 4, 5}))
	}
}

func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[2] >= nums[1]+nums[0] {
		return "none"
	}

	if nums[2] == nums[0] {
		return "equilateral"
	}

	if nums[2] == nums[1] || nums[1] == nums[0] {
		return "isosceles"
	}
	return "scalene"
}
