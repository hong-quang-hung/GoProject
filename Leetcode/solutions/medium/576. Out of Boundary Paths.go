package medium

import "fmt"

// Reference: https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
func init() {
	Solutions[576] = func() {
		fmt.Println(`Input: m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0`)
		fmt.Println(`Output:`, findPaths(2, 2, 2, 0, 0))
		fmt.Println(`Input: m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1`)
		fmt.Println(`Output:`, findPaths(1, 3, 3, 0, 1))
	}
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxMove)
			for z := 0; z < maxMove; z++ {
				dp[i][j][z] = -1
			}
		}
	}

	var f func(row, col, remain int) int
	f = func(row, col, remain int) int {
		if remain < 0 {
			return 0
		}

		if row >= m || row < 0 || col >= n || col < 0 {
			return 1
		}

		if dp[row][col][remain] != -1 {
			return dp[row][col][remain]
		}

		res := 0
		for _, dir := range BoardDirection {
			res += f(row+dir[0], col+dir[1], remain-1)
		}

		dp[row][col][remain] = res
		return res
	}
	return f(startRow, startColumn, maxMove-1)
}
