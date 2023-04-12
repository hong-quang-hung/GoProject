package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/
func Leetcode_Min_Number() {
	fmt.Println("Input: nums1 = [4,1,3], nums2 = [5,7]")
	fmt.Println("Output:", minNumber([]int{1, 3}, []int{5, 7}))
}

func minNumber(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				return nums1[i]
			}
		}
	}
	if nums1[0] < nums2[0] {
		return nums1[0]*10 + nums2[0]
	}
	return nums2[0]*10 + nums1[0]
}
