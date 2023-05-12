package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/uncrossed-lines/
func Leetcode_Max_Uncrossed_Lines() {
	fmt.Println("Input: nums1 = [1,4,2], nums2 = [1,2,4]")
	fmt.Println("Output:", maxUncrossedLines([]int{1, 4, 2}, []int{1, 2, 4}))
	fmt.Println("Input: nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]")
	fmt.Println("Output:", maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
	fmt.Println("Input: nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]")
	fmt.Println("Output:", maxUncrossedLines([]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n, m := len(nums1)+1, len(nums2)+1
	prev := make([]int, m)
	curr := make([]int, m)
	for i := 1; i < n; i++ {
		n1 := nums1[i-1]
		for j := 1; j < m; j++ {
			n2 := nums2[j-1]
			if n1 == n2 {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		curr, prev = prev, curr
	}
	return prev[m-1]
}
