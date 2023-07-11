package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/
func init() {
	Solutions[2684] = func() {
		fmt.Println("Input: grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]")
		fmt.Println("Output:", maxMoves(S2SoSliceInt("[[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]")))
		fmt.Println("Input: grid = [[3,2,4],[2,1,9],[1,1,7]]")
		fmt.Println("Output:", maxMoves(S2SoSliceInt("[[3,2,4],[2,1,9],[1,1,7]]")))
	}
}

func maxMoves(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			dp[i][j] = -1
		}
	}

	res := 0
	for i := 0; i < row; i++ {
		res = max(res, maxMovesDFS(grid, dp, i, 0, row, col, 0))
	}
	return res - 1
}

func maxMovesDFS(grid [][]int, dp [][]int, i int, j int, n int, m int, prev int) int {
	if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] <= prev {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	moves := [][]int{{-1, 1}, {0, 1}, {1, 1}}
	maxMove := 0
	for _, move := range moves {
		maxMove = max(maxMove, 1+maxMovesDFS(grid, dp, i+move[0], j+move[1], n, m, grid[i][j]))
	}

	dp[i][j] = maxMove
	return dp[i][j]
}
