package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/product-of-array-except-self/
func Leetcode_Product_ExceptSelf() {
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println("Input: nums = [-1, 1, 0, -3, 3]")
	fmt.Println("Output:", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		prefix[i] *= suffix
		suffix *= nums[i]
	}
	return prefix
}
