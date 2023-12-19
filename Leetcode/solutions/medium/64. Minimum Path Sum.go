package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-path-sum/
func init() {
	Solutions[64] = func() {
		fmt.Println(`Input: grid = [[1,3,1],[1,5,1],[4,2,1]]`)
		fmt.Println(`Output:`, minPathSum(S2SoSliceInt(`[[1,3,1],[1,5,1],[4,2,1]]`)))
		fmt.Println(`Input: grid = [[1,2,3],[4,5,6]]`)
		fmt.Println(`Output:`, minPathSum(S2SoSliceInt(`[[1,2,3],[4,5,6]]`)))
	}
}

func minPathSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r)

	for i := range dp {
		dp[i] = make([]int, c)
	}

	dp[0][0] += grid[0][0]
	for i := 1; i < c; i++ {
		dp[0][i] += dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < r; i++ {
		dp[i][0] += dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[r-1][c-1]
}
