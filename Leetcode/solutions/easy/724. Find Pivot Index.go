package easy

import "fmt"

func init() {
	Solutions[724] = Leetcode_Pivot_Index
}

// Reference: https://leetcode.com/problems/find-pivot-index/
func Leetcode_Pivot_Index() {
	fmt.Println("Input: nums = [1,7,3,6,5,6]")
	fmt.Println("Output:", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println("Input: nums = [1,2,3]")
	fmt.Println("Output:", pivotIndex([]int{1, 2, 3}))
	fmt.Println("Input: nums = [2,1,-1]")
	fmt.Println("Output:", pivotIndex([]int{2, 1, -1}))
}

func pivotIndex(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	left := 0
	for i := range nums {
		if left == nums[n-1]-nums[i] {
			return i
		}
		left = nums[i]
	}
	return -1
}
