package medium

import "fmt"

// Reference: https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/
func init() {
	Solutions[2482] = func() {
		fmt.Println(`Input: grid = [[0,1,1],[1,0,1],[0,0,1]]`)
		fmt.Println(`Output:`, onesMinusZeros(S2SoSliceInt(`[[0,1,1],[1,0,1],[0,0,1]]`)))
		fmt.Println(`Input: grid = [[1,1,1],[1,1,1]]`)
		fmt.Println(`Output:`, onesMinusZeros(S2SoSliceInt(`[[1,1,1],[1,1,1]]`)))
	}
}

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				row[i]--
				col[j]--
			} else {
				row[i]++
				col[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = row[i] + col[j]
		}
	}
	return res
}
