package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
func init() {
	Solutions[2419] = func() {
		fmt.Println(`Input: nums = [1,2,3,3,2,2]`)
		fmt.Println(`Output:`, longestSubarray_ii([]int{1, 2, 3, 3, 2, 2}))
		fmt.Println(`Input: nums = [1,2,3,3,2,2,2,4]`)
		fmt.Println(`Output:`, longestSubarray_ii([]int{1, 2, 3, 3, 2, 2, 2, 4}))
	}
}

func longestSubarray_ii(nums []int) int {
	res, curVal, curCount := 1, nums[0], 1
	for i := 1; i < len(nums); i++ {
		if curVal > nums[i] {
			curCount = 0
			continue
		}

		if curVal == nums[i] {
			curCount++
			res = max(res, curCount)
		} else {
			curVal = nums[i]
			curCount = 1
			res = 1
		}
	}
	return res
}
