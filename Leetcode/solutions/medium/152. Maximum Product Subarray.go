package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/maximum-product-subarray/
func init() {
	Solutions[152] = func() {
		fmt.Println("Input: nums = [2,3,-2,4]")
		fmt.Println("Output:", maxProduct([]int{2, 3, -2, 4}))
		fmt.Println("Input: nums = [-2,0,-1]")
		fmt.Println("Output:", maxProduct([]int{-2, 0, -1}))
	}
}

func maxProduct(nums []int) int {
	maxVal, minVal, res := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		temp := maxVal
		maxVal = max(num, max(maxVal*num, minVal*num))
		minVal = min(num, min(temp*num, minVal*num))
		res = max(res, maxVal)
	}
	return res
}
