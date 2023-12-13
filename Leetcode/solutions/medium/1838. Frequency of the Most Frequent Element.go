package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/frequency-of-the-most-frequent-element/
func init() {
	Solutions[1838] = func() {
		fmt.Println("Input: nums = [1,2,4], k = 5")
		fmt.Println("Output:", maxFrequency([]int{1, 2, 4}, 5))
		fmt.Println("Input: nums = [1,4,8,13], k = 5")
		fmt.Println("Output:", maxFrequency([]int{1, 4, 8, 13}, 5))
		fmt.Println("Input: nums = [3,9,6], k = 2")
		fmt.Println("Output:", maxFrequency([]int{3, 9, 6}, 2))
	}
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	res := 1
	for i := 1; i < len(nums); i++ {
		s := nums[i] - nums[i-1]
		for (i-left)*s > k {
			k += nums[i-1] - nums[left]
			left++
		}
		k -= (i - left) * s
		res = max(i-left+1, res)
	}
	return res
}
