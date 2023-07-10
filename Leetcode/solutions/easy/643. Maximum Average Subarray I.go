package easy

import "fmt"

func init() {
	Solutions[643] = Leetcode_Find_Max_Average
}

// Reference: https://leetcode.com/problems/maximum-average-subarray-i/
func Leetcode_Find_Max_Average() {
	fmt.Println("Input: nums = [1,12,-5,-6,50,3], k = 4")
	fmt.Println("Output:", findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println("Input: nums = [5], k = 1")
	fmt.Println("Output:", findMaxAverage([]int{5}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}
