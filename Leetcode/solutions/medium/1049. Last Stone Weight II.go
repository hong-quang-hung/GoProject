package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/last-stone-weight-ii/
func Leetcode_Last_Stone_Weight_II() {
	fmt.Println("Input: stones = [2,7,4,1,8,1]")
	fmt.Println("Output:", lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("Input: stones = [31,26,33,21,40]")
	fmt.Println("Output:", lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func lastStoneWeightII(stones []int) int {
	dp := make([]int, len(stones))
	return lastStoneWeightIISolve(stones, dp, 0)
}

func lastStoneWeightIISolve(stones []int, dp []int, i int) int {
	return 0
}
