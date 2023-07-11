package medium

import "fmt"

// Reference: https://leetcode.com/problems/sort-an-array/
func init() {
	Solutions[912] = func() {
		fmt.Println("Input: nums = [5,1,1,2,0,0]")
		fmt.Println("Output:", fmt.Sprint(sortArray([]int{5, 1, 1, 2, 0, 0})))
	}
}

func sortArray(nums []int) []int {
	countingSort(nums)
	return nums
}
