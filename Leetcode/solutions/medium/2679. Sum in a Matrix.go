package medium

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2679] = Leetcode_Matrix_Sum
}

// Reference: https://leetcode.com/problems/sum-in-a-matrix/
func Leetcode_Matrix_Sum() {
	fmt.Println("Input: nums = [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]")
	fmt.Println("Output:", matrixSum(S2SoSliceInt("[[7,2,1],[6,4,2],[6,5,3],[3,2,1]]")))
	fmt.Println("Input: nums =[[15,18,5,6,1,5,5,10,16,8,3,19],[14,5,0,15,13,18,16,9,20,11,12,8],[4,17,0,14,2,10,1,17,8,4,7,15],[11,19,9,11,18,20,19,4,9,12,13,20],[2,0,19,6,10,5,7,7,20,14,12,18],[13,1,12,18,13,1,5,14,2,8,5,14],[16,15,17,19,0,1,15,10,9,14,1,13],[6,17,20,2,4,0,12,13,10,0,6,0],[0,16,19,3,6,3,19,20,6,9,8,5]]")
	fmt.Println("Output:", matrixSum(S2SoSliceInt("[[15,18,5,6,1,5,5,10,16,8,3,19],[14,5,0,15,13,18,16,9,20,11,12,8],[4,17,0,14,2,10,1,17,8,4,7,15],[11,19,9,11,18,20,19,4,9,12,13,20],[2,0,19,6,10,5,7,7,20,14,12,18],[13,1,12,18,13,1,5,14,2,8,5,14],[16,15,17,19,0,1,15,10,9,14,1,13],[6,17,20,2,4,0,12,13,10,0,6,0],[0,16,19,3,6,3,19,20,6,9,8,5]]")))
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
