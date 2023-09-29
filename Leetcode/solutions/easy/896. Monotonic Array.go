package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/monotonic-array/
func init() {
	Solutions[896] = func() {
		fmt.Println("Input: nums = [1,2,2,3]")
		fmt.Println("Output:", isMonotonic([]int{1, 2, 2, 3}))
		fmt.Println("Input: nums = [1,3,2]")
		fmt.Println("Output:", isMonotonic([]int{1, 3, 2}))
		fmt.Println("Input: nums = [3,4,2,3]")
		fmt.Println("Output:", isMonotonic([]int{3, 4, 2, 3}))
	}
}

func isMonotonic(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return true
	}

	m := map[int]bool{}
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			if nums[i] > nums[i+1] {
				m[0] = true
			} else {
				m[1] = true
			}

			if len(m) == 2 {
				return false
			}
		}
	}
	return true
}
