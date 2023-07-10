package medium

import "fmt"

func init() {
	Solutions[912] = Leetcode_Sort_Array
}

// Reference: https://leetcode.com/problems/sort-an-array/
func Leetcode_Sort_Array() {
	fmt.Println("Input: nums = [5,1,1,2,0,0]")
	fmt.Println("Output:", fmt.Sprint(sortArray([]int{5, 1, 1, 2, 0, 0})))
}

func sortArray(nums []int) []int {
	countingSort(nums)
	return nums
}
