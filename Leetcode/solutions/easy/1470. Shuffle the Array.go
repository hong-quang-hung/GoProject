package easy

import "fmt"

// Reference: https://leetcode.com/problems/shuffle-the-array/
func init() {
	Solutions[1470] = func() {
		fmt.Println("Input: nums = [1,2,3,4,4,3,2,1], n = 4")
		fmt.Println("Output:", shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	}
}

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+len(nums)/2])
	}
	return res
}
