package medium

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2576] = Leetcode_Max_Num_Of_Marked_Indices
}

// Reference: https://leetcode.com/problems/find-the-maximum-number-of-marked-indices/
func Leetcode_Max_Num_Of_Marked_Indices() {
	fmt.Println("Input: nums = [3,5,2,4]")
	fmt.Println("Output:", maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
}

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i, n := 0, len(nums)
	for j := n / 2; i < n/2 && j < n; j++ {
		if 2*nums[i] <= nums[j] {
			i++
		}
	}
	return i * 2
}
