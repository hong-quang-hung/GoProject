package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-number-of-coins-you-can-get/
func init() {
	Solutions[1561] = func() {
		fmt.Println("Input: piles = [2,4,1,2,7,8]")
		fmt.Println("Output:", maxCoins([]int{2, 4, 1, 2, 7, 8}))
		fmt.Println("Input: piles = [2,4,5")
		fmt.Println("Output:", maxCoins([]int{2, 4, 5}))
		fmt.Println("Input: piles = [9,8,7,6,5,1,2,3,4]")
		fmt.Println("Output:", maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
	}
}

func maxCoins(piles []int) int {
	sort.Slice(piles, func(i, j int) bool { return piles[i] > piles[j] })
	n := len(piles) / 3
	res := 0
	for i := 1; i < len(piles)-n; i += 2 {
		res += piles[i]
	}
	return res
}
