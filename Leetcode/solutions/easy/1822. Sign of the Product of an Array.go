package easy

import "fmt"

// Reference: https://leetcode.com/problems/sign-of-the-product-of-an-array/
func Leetcode_Array_Sign() {
	fmt.Println("Input: nums = [-1,-2,-3,-4,3,2,1]")
	fmt.Println("Output:", arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println("Input: nums = [1,5,0,2,-3]")
	fmt.Println("Output:", arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println("Input: nums = [-1,1,-1,1,-1]")
	fmt.Println("Output:", arraySign([]int{-1, 1, -1, 1, -1}))
}

func arraySign(nums []int) int {
	product := 1
	for _, num := range nums {
		product *= signFunc(num)
		if product == 0 {
			return 0
		}
	}
	return product
}

func signFunc(product int) int {
	if product > 0 {
		return 1
	}
	if product < 0 {
		return -1
	}
	return 0
}
