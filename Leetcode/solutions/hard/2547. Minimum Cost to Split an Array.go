package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-cost-to-split-an-array/
func init() {
	Solutions[2547] = func() {
		fmt.Println("Input: nums = [1,2,1,2,1], k = 2")
		fmt.Println("Output:", minCost([]int{1, 2, 1, 2, 1}, 2))
		fmt.Println("Input: nums = [1,2,1,2,1], k = 5")
		fmt.Println("Output:", minCost([]int{1, 2, 1, 2, 1}, 5))
	}
}

func minCost(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	norms := minCostSolved(nums)
	for i := 1; i <= n; i++ {
		ones := 0
		count := make([]int, n)
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0; j-- {
			count[norms[j]]++
			if count[norms[j]] == 1 {
				ones++
			} else if count[norms[j]] == 2 {
				ones--
			}
			dp[i] = min(dp[i], dp[j]+k+i-j-ones)
		}
	}
	return dp[n]
}

func minCostSolved(nums []int) []int {
	norms := make([]int, len(nums))
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, c := m[nums[i]]; c {
			norms[i] = j
		} else {
			norms[i] = len(m)
			m[nums[i]] = norms[i]
		}
	}
	return norms
}
