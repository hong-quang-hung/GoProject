package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-zero-filled-subarrays/
func init() {
	Solutions[2348] = func() {
		fmt.Println("Input: nums = [0,0,0,2,0,0]")
		fmt.Println("Output:", zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	}
}

func zeroFilledSubarray(nums []int) int64 {
	res, size := int64(0), int64(0)
	for _, num := range nums {
		if num == 0 {
			size++
		} else {
			res += (size * (size + 1) / 2)
			size = 0
		}
	}

	if size == 0 {
		return res
	} else {
		return res + (size * (size + 1) / 2)
	}
}
