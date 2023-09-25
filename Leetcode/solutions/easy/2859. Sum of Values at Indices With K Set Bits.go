package easy

import (
	"fmt"
	"math/bits"
)

// Reference: https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/
func init() {
	Solutions[2859] = func() {
		fmt.Println("Input: nums = [5,10,1,5,2], k = 1")
		fmt.Println("Output:", sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
		fmt.Println("Input: nums = [4,3,2,1], k = 2")
		fmt.Println("Output:", sumIndicesWithKSetBits([]int{4, 3, 2, 1}, 2))
	}
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for i, num := range nums {
		if bits.OnesCount(uint(i)) == k {
			res += num
		}
	}
	return res
}
