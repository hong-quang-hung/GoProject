package medium

import "fmt"

// Reference: https://leetcode.com/problems/unique-paths-ii/
func init() {
	Solutions[63] = func() {
		fmt.Println("Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]")
		fmt.Println("Output:", uniquePathsWithObstacles(S2SoSliceInt("[[0,0,0],[0,1,0],[0,0,0]]")))
		fmt.Println("Input: obstacleGrid = [[0,1],[0,0]]")
		fmt.Println("Output:", uniquePathsWithObstacles(S2SoSliceInt("[[0,1],[0,0]]")))
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	return uniquePathsWithObstaclesSolved(obstacleGrid, dp, m, n, 0, 0)
}

func uniquePathsWithObstaclesSolved(obstacleGrid [][]int, dp [][]int, m, n, i, j int) int {
	if i >= m || j >= n || obstacleGrid[i][j] != 0 {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	res := 0
	for _, dir := range [][2]int{{1, 0}, {0, 1}} {
		res += uniquePathsWithObstaclesSolved(obstacleGrid, dp, m, n, i+dir[0], j+dir[1])
	}

	dp[i][j] = res
	return res
}
