package medium

import "fmt"

// Reference: https://leetcode.com/problems/house-robber-ii/
func Leetcode_House_Robber_II() {
	fmt.Println("Input: nums = [2,3,2]")
	fmt.Println("Output:", rob_ii([]int{2, 3, 2}))
	fmt.Println("Input: nums = [1,2,3,1]")
	fmt.Println("Output:", rob_ii([]int{1, 2, 3, 1}))
	fmt.Println("Input: nums = [1,2,3]")
	fmt.Println("Output:", rob_ii([]int{1, 2, 3}))
}

func rob_ii(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(rob_ii_solved(nums[1:]), rob_ii_solved(nums[0:n-1]))
}

func rob_ii_solved(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := 2; i < n+2; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[n+1]
}
