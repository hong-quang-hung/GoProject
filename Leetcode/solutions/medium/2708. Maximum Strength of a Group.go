package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-strength-of-a-group/
func Leetcode_Max_Strength() {
	fmt.Println("Input: nums = [3,-1,-5,2,5,-9]")
	fmt.Println("Output:", maxStrength([]int{3, -1, -5, 2, 5, -9}))
	fmt.Println("Input: nums = [-4,-5,-4]")
	fmt.Println("Output:", maxStrength([]int{-4, -5, -4}))
	fmt.Println("Input: nums = [-9]")
	fmt.Println("Output:", maxStrength([]int{-9}))
}

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	res, i := int64(0), 0
	for i < len(nums) && nums[i] > 0 {
		if res == 0 {
			res = 1
		}
		res *= int64(nums[i])
		i++
	}

	for i < len(nums) && nums[i] == 0 {
		i++
	}

	neg := nums[i:]
	n := len(neg)
	for j := n % 2; j < n; j++ {
		if res == 0 {
			res = int64(neg[n%2])
			continue
		}
		res *= int64(neg[j])
	}
	return res
}
