package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-all-duplicates-in-an-array/
func init() {
	Solutions[442] = func() {
		fmt.Println(`Input: nums = [4,3,2,7,8,2,3,1]`)
		fmt.Println(`Output:`, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
		fmt.Println(`Input: nums = [1,1,2]`)
		fmt.Println(`Output:`, findDuplicates([]int{1, 1, 2}))
	}
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := range nums {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			res = append(res, abs(nums[i]))
		}
		nums[index] = -nums[index]
	}
	return res
}
