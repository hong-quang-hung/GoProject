package medium

import "fmt"

// Reference: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func init() {
	Solutions[81] = func() {
		fmt.Println("Input: nums = [2,5,6,0,0,1,2], target = 0")
		fmt.Println("Output:", search_ii([]int{2, 5, 6, 0, 0, 1, 2}, 0))
		fmt.Println("Input: nums = [2,5,6,0,0,1,2], target = 3")
		fmt.Println("Output:", search_ii([]int{2, 5, 6, 0, 0, 1, 2}, 3))
		fmt.Println("Input: nums = [1,0,1,1,1], target = 0")
		fmt.Println("Output:", search_ii([]int{1, 0, 1, 1, 1}, 0))
	}
}

func search_ii(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[left] {
			left++
			continue
		}

		if nums[mid] >= nums[left] {
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
	return false
}
