package easy

import "fmt"

// Reference: https://leetcode.com/problems/single-number/
func init() {
	Solutions[136] = func() {
		fmt.Println("Input: nums = [2,2,1]")
		fmt.Println("Output:", singleNumber([]int{2, 2, 1}))
		fmt.Println("Input: nums = [4,1,2,1,2]")
		fmt.Println("Output:", singleNumber([]int{4, 1, 2, 1, 2}))
	}
}

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
