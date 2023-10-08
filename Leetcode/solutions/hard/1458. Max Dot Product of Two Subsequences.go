package hard

import "fmt"

// Reference: https://leetcode.com/problems/max-dot-product-of-two-subsequences/
func init() {
	Solutions[1458] = func() {
		fmt.Println("Input: nums1 = [2,1,-2,5], nums2 = [3,0,-6]")
		fmt.Println("Output:", maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
		fmt.Println("Input: nums1 = [3,-2], nums2 = [2,-6,7]")
		fmt.Println("Output:", maxDotProduct([]int{3, -2}, []int{2, -6, 7}))
		fmt.Println("Input: nums1 = [-1,-1], nums2 = [1,1]")
		fmt.Println("Output:", maxDotProduct([]int{-1, -1}, []int{1, 1}))
	}
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	return 1
}
