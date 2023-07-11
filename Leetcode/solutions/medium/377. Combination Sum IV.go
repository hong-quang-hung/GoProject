package medium

import "fmt"

// Reference: https://leetcode.com/problems/combination-sum-iii/
func init() {
	Solutions[377] = func() {
		fmt.Println("Input: nums = [1,2,3], target = 4")
		fmt.Println("Output:", combinationSum4([]int{1, 2, 3}, 4))
		fmt.Println("Input: nums = [9], target = 3")
		fmt.Println("Output:", combinationSum4([]int{9}, 3))
		fmt.Println("Input: nums = [2,1,3], target = 35")
		fmt.Println("Output:", combinationSum4([]int{2, 1, 3}, 35))
	}
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = -1
	}
	return combinationSum4Solved(nums, dp, target)
}

func combinationSum4Solved(nums []int, dp []int, target int) int {
	if target == 0 {
		return 1
	}

	if dp[target] != -1 {
		return dp[target]
	}

	sum := 0
	for _, num := range nums {
		if target-num >= 0 {
			sum += combinationSum4Solved(nums, dp, target-num)
		}
	}

	dp[target] = sum
	return dp[target]
}
