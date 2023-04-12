package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/number-of-enclaves/
func Leetcode_Num_Enclaves() {
	fmt.Println("Input: grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]")
	fmt.Println("Output:", numEnclaves(utils.S2SoSliceInt("[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]")))
	fmt.Println("Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]")
	fmt.Println("Output:", numEnclaves(utils.S2SoSliceInt("[[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]")))
}

func numEnclaves(grid [][]int) int {
	res, n, m := 0, len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && !visited[i][0] {
			numEnclavesDFS(grid, visited, i, 0, n, m)
		}
		if grid[i][n-1] == 1 && !visited[i][m-1] {
			numEnclavesDFS(grid, visited, i, m-1, n, m)
		}
	}

	for i := 0; i < m; i++ {
		if grid[0][i] == 1 && !visited[0][i] {
			numEnclavesDFS(grid, visited, 0, i, n, m)
		}
		if grid[n-1][i] == 1 && !visited[n-1][i] {
			numEnclavesDFS(grid, visited, n-1, i, n, m)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				res++
			}
		}
	}
	return res
}

func numEnclavesDFS(grid [][]int, visited [][]bool, i int, j int, n int, m int) {
	if i < 0 || i >= n || j >= m || j < 0 || grid[i][j] == 0 || visited[i][j] {
		return
	}

	visited[i][j] = true
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		numEnclavesDFS(grid, visited, i+move[0], j+move[1], n, m)
	}
}
