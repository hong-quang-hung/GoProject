package medium

import "fmt"

// Reference: https://leetcode.com/problems/check-knight-tour-configuration/
func Leetcode_Check_Valid_Grid() {
	fmt.Println("Input: grid = [[8,3,6],[5,0,1],[2,7,4]]")
	fmt.Println("Output:", checkValidGrid(S2SoSliceInt("[[8,3,6],[5,0,1],[2,7,4]]")))
}

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}

	m, n, i, j := 1, len(grid), 0, 0
	max := n*n - 1
	for m <= max {
		if i+2 < n && j+1 < n && checkValidGridMove(grid, &m, &i, &j, i+2, j+1) {
			continue
		}
		if i+2 < n && j-1 > -1 && checkValidGridMove(grid, &m, &i, &j, i+2, j-1) {
			continue
		}
		if i-2 > -1 && j+1 < n && checkValidGridMove(grid, &m, &i, &j, i-2, j+1) {
			continue
		}
		if i-2 > -1 && j-1 > -1 && checkValidGridMove(grid, &m, &i, &j, i-2, j-1) {
			continue
		}
		if i+1 < n && j-2 > -1 && checkValidGridMove(grid, &m, &i, &j, i+1, j-2) {
			continue
		}
		if i+1 < n && j+2 < n && checkValidGridMove(grid, &m, &i, &j, i+1, j+2) {
			continue
		}
		if i-1 > -1 && j-2 > -1 && checkValidGridMove(grid, &m, &i, &j, i-1, j-2) {
			continue
		}
		if i-1 > -1 && j+2 < n && checkValidGridMove(grid, &m, &i, &j, i-1, j+2) {
			continue
		}
		return false
	}
	return true
}

func checkValidGridMove(grid [][]int, m *int, i *int, j *int, x int, y int) bool {
	if grid[x][y] != *m {
		return false
	}

	*i = x
	*j = y
	*m++
	return true
}
