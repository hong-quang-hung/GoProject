package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/4sum/
func Leetcode_Four_Sum() {
	fmt.Println("Input: nums = [1,0,-1,0,-2,2], target = 0")
	fmt.Println("Output:", fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println("Input: nums = [2,2,2,2,2], target = 8")
	fmt.Println("Output:", fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSum(nums, int64(target), 0, 4)
}

func kSum(nums []int, target int64, start int, k int) [][]int {
	res := make([][]int, 0)
	if start == len(nums) {
		return res
	}

	average_value := target / int64(k)
	if int64(nums[start]) > average_value || average_value > int64(nums[len(nums)-1]) {
		return res
	}

	if k == 2 {
		return twoSum(nums, target, start)
	}

	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSum(nums, target-int64(nums[i]), i+1, k-1) {
				res = append(res, []int{nums[i]})
				res[len(res)-1] = append(res[len(res)-1], subset...)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int64, start int) [][]int {
	res := make([][]int, 0)
	s := make(map[int64]bool)
	for i := start; i < len(nums); i++ {
		if len(res) == 0 || res[len(res)-1][1] != nums[i] {
			if s[target-int64(nums[i])] {
				res = append(res, []int{int(target) - nums[i], nums[i]})
			}
		}
		s[int64(nums[i])] = true
	}
	return res
}
