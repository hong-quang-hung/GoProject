package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/sum-in-a-matrix/
func init() {
	Solutions[2679] = func() {
		fmt.Println("Input: nums = [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]")
		fmt.Println("Output:", matrixSum(S2SoSliceInt("[[7,2,1],[6,4,2],[6,5,3],[3,2,1]]")))
	}
}

func matrixSum(nums [][]int) int {
	res, n, m := 0, len(nums), len(nums[0])
	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
	}

	for j := 0; j < m; j++ {
		maxSum := 0
		for i := 0; i < n; i++ {
			maxSum = max(maxSum, nums[i][j])
		}
		res += maxSum
	}
	return res
}
