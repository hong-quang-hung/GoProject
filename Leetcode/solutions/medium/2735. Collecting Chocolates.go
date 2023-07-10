package medium

import "fmt"

func init() {
	Solutions[2735] = Leetcode_Min_Cost
}

// Reference: https://leetcode.com/problems/collecting-chocolates/
func Leetcode_Min_Cost() {
	fmt.Println("Input: nums = [20,1,15], x = 5")
	fmt.Println("Output:", minCost([]int{20, 1, 15}, 5))
	fmt.Println("Input: nums = [1,2,3], x = 4")
	fmt.Println("Output:", minCost([]int{1, 2, 3}, 4))
}

func minCost(nums []int, x int) int64 {
	n := len(nums)
	dp := make([]int64, n)
	res := int64(0)
	for i := 0; i < n; i++ {
		dp[i] = int64(nums[i])
		res += int64(nums[i])
	}

	for i := 1; i < n; i++ {
		sum := int64(i) * int64(x)
		for j := 0; j < n; j++ {
			cur := j + i - n
			if i+j < n {
				cur = j + i
			}

			dp[j] = min64(dp[j], int64(nums[cur]))
			sum += int64(dp[j])
		}
		res = min64(res, sum)
	}
	return res
}
