package easy

import "fmt"

// Reference: https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
func init() {
	Solutions[1464] = func() {
		fmt.Println("Input: nums = [3,4,5,2]")
		fmt.Println("Output:", maxProduct([]int{3, 4, 5, 2}))
		fmt.Println("Input: nums = [1,5,4,5]")
		fmt.Println("Output:", maxProduct([]int{1, 5, 4, 5}))
		fmt.Println("Input: nums = [3,7]")
		fmt.Println("Output:", maxProduct([]int{3, 7}))
	}
}

func maxProduct(nums []int) int {
	first, second := 0, 0
	for _, num := range nums {
		if num > first {
			second = first
			first = num
		} else {
			second = max(second, num)
		}
	}
	return (first - 1) * (second - 1)
}
