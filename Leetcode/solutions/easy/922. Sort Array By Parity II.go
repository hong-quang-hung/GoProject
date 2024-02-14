package easy

import "fmt"

// Reference: https://leetcode.com/problems/sort-array-by-parity-ii/
func init() {
	Solutions[922] = func() {
		fmt.Println(`Input: nums = [4,2,5,7]`)
		fmt.Println(`Output:`, sortArrayByParityII([]int{4, 2, 5, 7}))
		fmt.Println(`Input: nums = [2,3]`)
		fmt.Println(`Output:`, sortArrayByParityII([]int{2, 3}))
	}
}

func sortArrayByParityII(nums []int) []int {
	return nums
}
