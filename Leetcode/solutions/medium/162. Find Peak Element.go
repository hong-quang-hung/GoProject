package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-peak-element/
func init() {
	Solutions[162] = func() {
		fmt.Println("Input: nums = [1,2,3,1]")
		fmt.Println("Output:", findPeakElement([]int{1, 2, 3, 1}))
		fmt.Println("Input: nums = [1,2,1,3,5,6,4]")
		fmt.Println("Output:", findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
		fmt.Println("Input: nums = [3,1,2]")
		fmt.Println("Output:", findPeakElement([]int{3, 1, 2}))
	}
}

func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if mid < n-1 && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			if mid > 0 && nums[mid] > nums[mid-1] {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return 0
}
