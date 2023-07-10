package medium

import "fmt"

func init() {
	Solutions[2348] = Leetcode_Zero_Filled_Subarray
}

// Reference: https://leetcode.com/problems/number-of-zero-filled-subarrays/
func Leetcode_Zero_Filled_Subarray() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output:", zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
}

func zeroFilledSubarray(nums []int) int64 {
	var res int64 = 0
	var size int64 = 0
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
