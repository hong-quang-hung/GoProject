package medium

import "fmt"

func init() {
	Solutions[137] = Leetcode_Single_Number_II
}

// Reference: https://leetcode.com/problems/single-number-ii/
func Leetcode_Single_Number_II() {
	fmt.Println("Input: nums = [2,2,3,2]")
	fmt.Println("Output:", singleNumber([]int{2, 2, 3, 2}))
	fmt.Println("Input: nums = [0,1,0,1,0,1,99]")
	fmt.Println("Output:", singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println("Input: nums = [2,5,2,11,11,3,2,5,5,11]")
	fmt.Println("Output:", singleNumber([]int{2, 5, 2, 11, 11, 3, 2, 5, 5, 11}))
}

func singleNumber(nums []int) int {
	b1, b2 := 0, 0
	for _, nums := range nums {
		b1 = (b1 ^ nums) & (^b2)
		b2 = (b2 ^ nums) & (^b1)
	}
	return b1
}
