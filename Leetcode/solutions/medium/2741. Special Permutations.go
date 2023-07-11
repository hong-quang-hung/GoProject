package medium

import "fmt"

// Reference: https://leetcode.com/problems/special-permutations/
func init() {
	Solutions[2741] = func() {
		fmt.Println("Input: nums = [2,3,6]")
		fmt.Println("Output:", specialPerm([]int{2, 3, 6}))
		fmt.Println("Input: nums = [20,100,50,5,10,70,7]")
		fmt.Println("Output:", specialPerm([]int{20, 100, 50, 5, 10, 70, 7}))
	}
}

func specialPerm(nums []int) int {
	n := len(nums)
	target := (1 << n) - 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target)
		for j := 0; j < target; j++ {
			dp[i][j] = -1
		}
	}
	return specialPermSolved(nums, dp, 0, 0, target)
}

func specialPermSolved(nums []int, dp [][]int, mask int, i int, target int) int {
	if mask == target {
		return 1
	}

	if dp[i][mask] == -1 {
		dp[i][mask] = 0
		for j := 0; j < len(nums); j++ {
			newMask := 1 << j
			if (mask&newMask) == 0 && (mask == 0 || nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0) {
				dp[i][mask] = (dp[i][mask] + specialPermSolved(nums, dp, mask+newMask, j, target)) % mod
			}
		}
	}
	return dp[i][mask]
}
