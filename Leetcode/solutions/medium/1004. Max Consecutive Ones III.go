package medium

import "fmt"

// Reference: https://leetcode.com/problems/max-consecutive-ones-iii/
func init() {
	Solutions[1004] = func() {
		fmt.Println("Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2")
		fmt.Println("Output:", longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
		fmt.Println("Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3")
		fmt.Println("Output:", longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	}
}

func longestOnes(nums []int, k int) int {
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
	return i - j
}
