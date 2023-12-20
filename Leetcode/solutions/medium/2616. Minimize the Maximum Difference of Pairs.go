package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/
func init() {
	Solutions[2616] = func() {
		fmt.Println(`Input: nums = [10,1,2,7,1,3], p = 2`)
		fmt.Println(`Output:`, minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
		fmt.Println(`Input: nums = [4,2,1,2], p = 1`)
		fmt.Println(`Output:`, minimizeMax([]int{4, 2, 1, 2}, 1))
	}
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := (left + right) / 2
		if minimizeMaxCountPairs(nums, mid, n) >= p {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func minimizeMaxCountPairs(nums []int, mid int, n int) int {
	index, count := 0, 0
	for index < n-1 {
		if nums[index+1]-nums[index] <= mid {
			count++
			index++
		}
		index++
	}
	return count
}
