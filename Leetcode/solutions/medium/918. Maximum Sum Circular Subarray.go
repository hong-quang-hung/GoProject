package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-sum-circular-subarray/
func init() {
	Solutions[918] = func() {
		fmt.Println(`Input: nums = [1,-2,3,-2]`)
		fmt.Println(`Output:`, maxSubarraySumCircular([]int{1, -2, 3, -2}))
		fmt.Println(`Input: nums = [5,-3,5]`)
		fmt.Println(`Output:`, maxSubarraySumCircular([]int{5, -3, 5}))
		fmt.Println(`Input: nums = [-3,-2,-3]`)
		fmt.Println(`Output:`, maxSubarraySumCircular([]int{-3, -2, -3}))
	}
}

func maxSubarraySumCircular(nums []int) int {
	totalSum := 0
	maxSum, curMax := nums[0], 0
	minSum, curMin := nums[0], 0
	for _, num := range nums {
		curMax = max(num, curMax+num)
		maxSum = max(maxSum, curMax)

		curMin = min(num, curMin+num)
		minSum = min(minSum, curMin)

		totalSum += num
	}

	if totalSum == minSum {
		return maxSum
	}
	return max(maxSum, totalSum-minSum)
}
