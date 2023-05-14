package hard

import (
	"fmt"
	"math/bits"
)

// Reference: hhttps://leetcode.com/problems/maximize-score-after-n-operations/
func Leetcode_Max_Score() {
	fmt.Println("Input: nums = [1,2]")
	fmt.Println("Output:", maxScore([]int{1, 2}))
	fmt.Println("Input: nums = [3,4,6,8]")
	fmt.Println("Output:", maxScore([]int{3, 4, 6, 8}))
	fmt.Println("Input: nums = [1,2,3,4,5,6]")
	fmt.Println("Output:", maxScore([]int{1, 2, 3, 4, 5, 6}))
}

func maxScore(nums []int) int {
	length := len(nums)
	maxStates := 1 << length
	dp, finalMask := make([]int, maxStates), maxStates-1
	dp[finalMask] = 0
	for state := finalMask - 1; state >= 0; state-- {
		numbersTaken := bits.OnesCount(uint(state))
		if numbersTaken%2 != 0 {
			continue
		}

		pairsFormed := numbersTaken / 2
		for firstIndex := 0; firstIndex < length; firstIndex++ {
			for secondIndex := firstIndex + 1; secondIndex < length; secondIndex++ {
				if ((state>>firstIndex)&1) == 1 || ((state>>secondIndex)&1) == 1 {
					continue
				}

				currentScore := (pairsFormed + 1) * gcd(nums[firstIndex], nums[secondIndex])
				stateAfterPickingCurrPair := state | (1 << firstIndex) | (1 << secondIndex)
				dp[state] = max(dp[state], currentScore+dp[stateAfterPickingCurrPair])
			}
		}
	}
	return dp[0]
}
