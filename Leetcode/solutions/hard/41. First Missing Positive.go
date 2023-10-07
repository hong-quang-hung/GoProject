package hard

import "fmt"

// Reference: https://leetcode.com/problems/first-missing-positive/
func init() {
	Solutions[41] = func() {
		fmt.Println("Input: nums = [1,2,0]")
		fmt.Println("Output:", firstMissingPositive([]int{1, 2, 0}))
		fmt.Println("Input: nums = [3,4,-1,1]")
		fmt.Println("Output:", firstMissingPositive([]int{3, 4, -1, 1}))
		fmt.Println("Input: nums = [7,8,9,11,12]")
		fmt.Println("Output:", firstMissingPositive([]int{7, 8, 9, 11, 12}))
	}
}

func firstMissingPositive(nums []int) int {
	return 0
}
