package medium

import "fmt"

// Reference: https://leetcode.com/problems/search-in-rotated-sorted-array/
func init() {
	Solutions[33] = func() {
		fmt.Println("Input: nums = [4,5,6,7,0,1,2], target = 0")
		fmt.Println("Output:", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
		fmt.Println("Input: nums = [4,5,6,7,0,1,2], target = 3")
		fmt.Println("Output:", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
		fmt.Println("Input: nums = [1], target = 0")
		fmt.Println("Output:", search([]int{1}, 0))
	}
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
