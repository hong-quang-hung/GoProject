package medium

import "fmt"

// Reference: https://leetcode.com/problems/rotate-array/
func init() {
	Solutions[189] = func() {
		var nums []int
		fmt.Println(`Input: nums = [1,2,3,4,5,6,7], k = 3`)
		nums = []int{1, 2, 3, 4, 5, 6, 7}
		rotateII(nums, 3)
		fmt.Println(`Output:`, nums)
		fmt.Println(`Input: nums = [-1,-100,3,99], k = 2`)
		nums = []int{-1, -100, 3, 99}
		rotateII(nums, 2)
		fmt.Println(`Output:`, nums)
	}
}

func rotateII(nums []int, k int) {
	n := len(nums)
	k %= n
	reverseSlice(nums, 0, n-1)
	reverseSlice(nums, 0, k-1)
	reverseSlice(nums, k, n-1)
}

func reverseSlice(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
