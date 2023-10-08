package hard

import (
	"fmt"
	"math"
)

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
	max1, max2, min1, min2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
	for _, num := range nums1 {
		max1 = max(max1, num)
		min1 = min(min1, num)
	}

	for _, num := range nums2 {
		max2 = max(max2, num)
		min2 = min(min2, num)
	}

	if max1 < 0 && min2 > 0 {
		return max1 * min2
	}

	if min1 > 0 && max2 < 0 {
		return min1 * max2
	}

	n, m := len(nums1), len(nums2)
	dp := make([]int, m+1)
	prev := make([]int, m+1)
	for i := n - 1; i >= 0; i-- {
		dp = make([]int, m+1)
		for j := m - 1; j >= 0; j-- {
			use := nums1[i]*nums2[j] + prev[j+1]
			dp[j] = max(use, max(prev[j], dp[j+1]))
		}
		prev = dp
	}
	return dp[0]
}
