package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/the-number-of-beautiful-subsets/
func init() {
	Solutions[2597] = func() {
		fmt.Println("Input: nums = [10,4,5,7,2,1], k = 3")
		fmt.Println("Output:", beautifulSubsets([]int{10, 4, 5, 7, 2, 1}, 3))
	}
}

func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	return counter(0, map[int]int{}, nums, k) - 1
}

func counter(index int, m map[int]int, nums []int, k int) int {
	if index == len(nums) {
		return 1
	}
	count := counter(index+1, m, nums, k)
	if m[nums[index]-k] == 0 {
		m[nums[index]]++
		count += counter(index+1, m, nums, k)
		m[nums[index]]--
	}
	return count
}
