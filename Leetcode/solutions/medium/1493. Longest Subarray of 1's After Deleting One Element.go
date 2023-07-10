package medium

import "fmt"

func init() {
	Solutions[1493] = Leetcode_Longest_Subarray
}

// Reference: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
func Leetcode_Longest_Subarray() {
	fmt.Println("Input: nums = [1,1,0,1]")
	fmt.Println("Output:", longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println("Input: nums = [0,1,1,1,0,1,1,0,1]")
	fmt.Println("Output:", longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println("Input: nums = [1,1,1]")
	fmt.Println("Output:", longestSubarray([]int{1, 1, 1}))
}

func longestSubarray(nums []int) int {
	k := 1
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			k--
		}

		if k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
		i++
	}
	return i - j - 1
}
