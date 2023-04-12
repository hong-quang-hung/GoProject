package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
func Leetcode_Count_Subarrays() {
	fmt.Println("Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5")
	fmt.Println("Output:", countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	fmt.Println("Input: nums = [1,1,1,1], minK = 1, maxK = 1")
	fmt.Println("Output:", countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	min, max, left := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			left = i
		}
		if nums[i] == maxK {
			max = i
		}
		if nums[i] == minK {
			min = i
		}
		if min == -1 || max == -1 {
			continue
		}
		var minLeft int
		if min > max {
			minLeft = max
		} else {
			minLeft = min
		}
		if minLeft > left {
			res += int64(minLeft - left)
		}
	}
	return res
}
