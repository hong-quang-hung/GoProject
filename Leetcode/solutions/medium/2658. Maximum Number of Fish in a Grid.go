package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/
func Leetcode_Find_Max_Fish() {
	fmt.Println("Input: grid = [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]")
	fmt.Println("Output:", findMaxFish(utils.S2SoSliceInt("[[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]")))
	fmt.Println("Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]")
	fmt.Println("Output:", findMaxFish(utils.S2SoSliceInt("[[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]")))
}

func findMaxFish(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 0 && !visited[i][j] {
				res = max(res, findMaxFishGet(grid, visited, n, m, i, j))
			}
		}
	}
	return res
}

func findMaxFishGet(grid [][]int, visited [][]bool, n, m, i, j int) int {
	if i < 0 || j < 0 || i == n || j == m || visited[i][j] || grid[i][j] == 0 {
		return 0
	}

	visited[i][j] = true
	sum := grid[i][j]
	sum += findMaxFishGet(grid, visited, n, m, i-1, j)
	sum += findMaxFishGet(grid, visited, n, m, i, j-1)
	sum += findMaxFishGet(grid, visited, n, m, i+1, j)
	sum += findMaxFishGet(grid, visited, n, m, i, j+1)
	return sum
}
