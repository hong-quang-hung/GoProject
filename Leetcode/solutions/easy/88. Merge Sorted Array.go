package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/merge-sorted-array/
func init() {
	Solutions[88] = func() {
		fmt.Println(`Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3`)
		nums1 := []int{1, 2, 3, 0, 0, 0}
		merge(nums1, 3, []int{2, 5, 6}, 3)
		fmt.Println(`Output:`, nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n && i+m < len(nums1); i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}
