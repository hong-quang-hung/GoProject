package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/triangle/
func init() {
	Solutions[120] = func() {
		fmt.Println(`Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]`)
		fmt.Println(`Output:`, minimumTotal(S2SoSliceInt(`[[2],[3,4],[6,5,7],[4,1,8,3]]`)))
		fmt.Println(`Input: triangle = [[-10]]`)
		fmt.Println(`Output:`, minimumTotal(S2SoSliceInt(` [[-10]]`)))
	}
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	var f func(i, j int) int
	f = func(i, j int) int {
		if i == n-1 {
			return triangle[i][j]
		}
		if dp[i][j] != math.MaxInt {
			return dp[i][j]
		}

		res := triangle[i][j] + min(f(i+1, j), f(i+1, j+1))
		dp[i][j] = res
		return res
	}
	return f(0, 0)
}
