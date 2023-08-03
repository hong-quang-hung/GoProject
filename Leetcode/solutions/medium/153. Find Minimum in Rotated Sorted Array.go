package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func init() {
	Solutions[153] = func() {
		// fmt.Println("Input: nums = [3,4,5,1,2]")
		// fmt.Println("Output:", findMin([]int{3, 4, 5, 1, 2}))
		// fmt.Println("Input: nums = [4,5,6,7,0,1,2]")
		// fmt.Println("Output:", findMin([]int{4, 5, 6, 7, 0, 1, 2}))
		// fmt.Println("Input: nums = [11,13,15,17]")
		// fmt.Println("Output:", findMin([]int{11, 13, 15, 17}))

		fmt.Println("Input: nums = [5,1,2,3,4]")
		fmt.Println("Output:", findMin([]int{5, 1, 2, 3, 4}))
	}
}

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return min(nums[0], nums[1])
	}

	for left, right := 0, n-1; left <= right; {
		mid := (left + right) / 2
		if (mid == 0 && nums[mid] < nums[mid+1]) || (mid == n-1 && nums[mid] < nums[mid-1]) || (mid > 0 && mid < n-1 && nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1]) {
			return nums[mid]
		} else {
			if nums[mid] > nums[left] && nums[left] < nums[right] || nums[mid] < nums[left] && nums[left] > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
