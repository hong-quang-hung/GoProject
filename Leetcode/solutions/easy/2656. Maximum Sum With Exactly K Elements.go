package easy

import "fmt"

// Reference: https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/
func init() {
	Solutions[2656] = func() {
		fmt.Println("Input: nums = [1,2,3,4,5], k = 3")
		fmt.Println("Output:", maximizeSum([]int{1, 2, 3, 4, 5}, 3))
		fmt.Println("Input: nums = [5,5,5], k = 2")
		fmt.Println("Output:", maximizeSum([]int{5, 5, 5}, 2))
	}
}

func maximizeSum(nums []int, k int) int {
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}

	res := 0
	for i := 0; i < k; i++ {
		res += maxNum + i
	}
	return res
}
