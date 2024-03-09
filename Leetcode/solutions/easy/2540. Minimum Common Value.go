package easy

import "fmt"

// Reference: https://leetcode.com/problems/minimum-common-value/
func init() {
	Solutions[2540] = func() {
		fmt.Println(`Input: nums1 = [1,2,3], nums2 = [2,4]`)
		fmt.Println(`Output:`, getCommon([]int{1, 2, 3}, []int{2, 4}))
		fmt.Println(`Input: nums1 = [1,2,3,6], nums2 = [2,3,4,5]`)
		fmt.Println(`Output:`, getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))
	}
}

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			break
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	if i >= len(nums1) || j >= len(nums2) {
		return -1
	}
	return nums1[i]
}
