package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/
func Leetcode_Largest_Combination() {
	fmt.Println("Input: candidates = [16,17,71,62,12,24,14]")
	fmt.Println("Output:", largestCombination([]int{16, 17, 71, 62, 12, 24, 14}))
	fmt.Println("Input: candidates = [8,8]")
	fmt.Println("Output:", largestCombination([]int{8, 8}))
}

func largestCombination(candidates []int) int {
	res := 0
	arr := make([]int, 24)
	for _, cacandidate := range candidates {
		for i := 0; i < 24; i++ {
			if (cacandidate>>i)&1 == 1 {
				arr[i]++
			}
		}
	}

	for i := 0; i < 24; i++ {
		res = max(res, arr[i])
	}
	return res
}
