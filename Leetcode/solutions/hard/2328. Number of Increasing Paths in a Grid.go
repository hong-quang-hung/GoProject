package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/
func init() {
	Solutions[2328] = func() {
		fmt.Println(`Input: grid = [[1,1],[3,4]]`)
		fmt.Println(`Output:`, countPaths(S2SoSliceInt(`[[1,1],[3,4]]`)))
		fmt.Println(`Input: grid = [[1],[2]]`)
		fmt.Println(`Output:`, countPaths(S2SoSliceInt(`[[1],[2]]`)))
	}
}

func countPaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = (res + countPathsDFS(grid, dp, i, j, m, n)) % mod
		}
	}
	return res
}

func countPathsDFS(grid [][]int, dp [][]int, i, j, m, n int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	count := 1
	moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, move := range moves {
		prevI := i + move[0]
		prevJ := j + move[1]
		if 0 <= prevI && prevI < m && 0 <= prevJ && prevJ < n && grid[prevI][prevJ] < grid[i][j] {
			count += countPathsDFS(grid, dp, prevI, prevJ, m, n)
			count %= mod
		}
	}

	dp[i][j] = count
	return count
}
