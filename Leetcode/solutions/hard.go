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
	n, m := len(nums1), len(nums2)
	res := make([]int, (n+m)/2+1)
	even := (n+m)&1 == 0

	l, r := 0, 0
	for i := 0; i < len(res); i++ {

		if l == len(nums1) && r < len(nums2) {
			res[i] = nums2[r]
			r++
			continue
		} else if r == len(nums2) && l < len(nums1) {
			res[i] = nums1[l]
			l++
			continue
		}

		if nums1[l] <= nums2[r] {
			res[i] = nums1[l]
			l++
		} else if r < len(nums2) {
			res[i] = nums2[r]
			r++
		}

	}

	if even && n+m > 1 {
		return float64(res[len(res)-1]+res[len(res)-2]) / 2
	} else {
		return float64(res[len(res)-1])
	}
}
