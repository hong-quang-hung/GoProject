package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-falling-path-sum/
func init() {
	Solutions[931] = func() {
		fmt.Println(`Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]`)
		fmt.Println(`Output:`, minFallingPathSum(S2SoSliceInt("[[2,1,3],[6,5,4],[7,8,9]]")))
		fmt.Println(`Input: matrix = [[-19,57],[-40,-5]]`)
		fmt.Println(`Output:`, minFallingPathSum(S2SoSliceInt("[[-19,57],[-40,-5]]")))
	}
}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = matrix[0][i]
	}

	for i := 1; i < n; i++ {
		prevDp := make([]int, n)
		copy(prevDp, dp)
		dp[0] = min(prevDp[0], prevDp[1]) + matrix[i][0]
		for j := 1; j < n-1; j++ {
			dp[j] = min(prevDp[j-1], prevDp[j], prevDp[j+1]) + matrix[i][j]
		}
		dp[n-1] = min(prevDp[n-1], prevDp[n-2]) + matrix[i][n-1]
		prevDp = nil
	}

	res := 100 * n
	for i := 0; i < n; i++ {
		res = min(res, dp[i])
	}
	return res
}
