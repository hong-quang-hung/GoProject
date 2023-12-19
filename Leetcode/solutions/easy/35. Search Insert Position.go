package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/search-insert-position/
func init() {
	Solutions[35] = func() {
		fmt.Println(`Input: nums = [1,3,5,6], target = 7`)
		fmt.Println(`Output:`, searchInsert([]int{1, 3, 5, 6}, 7))
	}
}

func searchInsert(nums []int, target int) int {
	mid, left, right := 0, 0, len(nums)-1
	for left <= right {
		mid = int(math.Floor(float64(left+right) / float64(2)))
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
