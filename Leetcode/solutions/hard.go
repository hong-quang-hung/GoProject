package solutions

// Reference: https://leetcode.com/problems/minimize-deviation-in-array/
func minimumDeviation(nums []int) int {
	return 0
}

// Reference: https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
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

// Reference: https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	return float64(n1) / float64(n2)
}
