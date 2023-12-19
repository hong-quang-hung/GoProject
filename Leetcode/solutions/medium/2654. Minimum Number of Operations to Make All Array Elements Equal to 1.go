package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
func init() {
	Solutions[2654] = func() {
		fmt.Println(`Input: nums = [2,6,3,4]`)
		fmt.Println(`Output:`, minOperations_ii([]int{2, 6, 3, 4}))
		fmt.Println(`Input: nums = [2,10,6,14]`)
		fmt.Println(`Output:`, minOperations_ii([]int{2, 10, 6, 14}))
		fmt.Println(`Input: nums = [6,10,15]`)
		fmt.Println(`Output:`, minOperations_ii([]int{6, 10, 15}))
	}
}

func minOperations_ii(nums []int) int {
	ones, n := 0, len(nums)
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}

	if ones > 0 {
		return n - ones
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				res = min(res, j-i)
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}
	return res + n - 1
}
