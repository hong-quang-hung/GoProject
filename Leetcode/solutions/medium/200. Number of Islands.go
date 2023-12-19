package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-islands/
func init() {
	Solutions[200] = func() {
		fmt.Println(`Input: grid = [['1','1','1','1','0'],['1','1','0','1','0'],['1','1','0','0','0'], ['0','0','0','0','0']]`)
		fmt.Println(`Output:`, numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
		fmt.Println(`Input: grid = [['1','1','1','1','0'],['1','1','0','1','0'],['1','1','0','0','0'], ['0','0','0','0','0']]`)
		fmt.Println(`Output:`, numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
	}
}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if numIslandsCheck(grid, visited, m, n, i, j) {
				res++
			}
		}
	}
	return res
}

func numIslandsCheck(grid [][]byte, visited [][]bool, m int, n int, i int, j int) bool {
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == '0' {
		return false
	}

	visited[i][j] = true
	for _, dir := range BoardDirection {
		numIslandsCheck(grid, visited, m, n, i+dir[0], j+dir[1])
	}
	return true
}
