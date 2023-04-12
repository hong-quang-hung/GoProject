package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/number-of-closed-islands/
func Leetcode_Closed_Island() {
	fmt.Println("Input: grid = [[1,1,1,1,1,1,1],[1,0,0,0,0,0,1],[1,0,1,1,1,0,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,0,0,0,0,0,1],[1,1,1,1,1,1,1]]")
	fmt.Println("Output:", closedIsland(utils.S2SoSliceInt("[[1,1,1,1,1,1,1],[1,0,0,0,0,0,1],[1,0,1,1,1,0,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,0,0,0,0,0,1],[1,1,1,1,1,1,1]]")))
	fmt.Println("Input: grid = [[0,0,1,1,0,1,0,0,1,0],[1,1,0,1,1,0,1,1,1,0],[1,0,1,1,1,0,0,1,1,0],[0,1,1,0,0,0,0,1,0,1],[0,0,0,0,0,0,1,1,1,0],[0,1,0,1,0,1,0,1,1,1],[1,0,1,0,1,1,0,0,0,1],[1,1,1,1,1,1,0,0,0,0],[1,1,1,0,0,1,0,1,0,1],[1,1,1,0,1,1,0,1,1,0]]")
	fmt.Println("Output:", closedIsland(utils.S2SoSliceInt("[[0,0,1,1,0,1,0,0,1,0],[1,1,0,1,1,0,1,1,1,0],[1,0,1,1,1,0,0,1,1,0],[0,1,1,0,0,0,0,1,0,1],[0,0,0,0,0,0,1,1,1,0],[0,1,0,1,0,1,0,1,1,1],[1,0,1,0,1,1,0,0,0,1],[1,1,1,1,1,1,0,0,0,0],[1,1,1,0,0,1,0,1,0,1],[1,1,1,0,1,1,0,1,1,0]]")))
}

func closedIsland(grid [][]int) int {
	res, n, m := 0, len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 0 && !visited[i][j] && closedIslandDFS(grid, visited, i, j, n, m) {
				res++
			}
		}
	}
	return res
}

func closedIslandDFS(grid [][]int, visited [][]bool, i int, j int, n int, m int) bool {
	if i < 0 || i >= n || j >= m || j < 0 {
		return false
	}

	if grid[i][j] == 1 || visited[i][j] {
		return true
	}

	visited[i][j] = true
	up := closedIslandDFS(grid, visited, i-1, j, n, m)
	down := closedIslandDFS(grid, visited, i+1, j, n, m)
	left := closedIslandDFS(grid, visited, i, j-1, n, m)
	right := closedIslandDFS(grid, visited, i, j+1, n, m)
	return up && down && left && right
}
