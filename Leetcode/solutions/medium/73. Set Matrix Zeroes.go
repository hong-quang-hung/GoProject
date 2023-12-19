package medium

import "fmt"

// Reference: https://leetcode.com/problems/set-matrix-zeroes/
func init() {
	Solutions[73] = func() {
		var matrix [][]int
		fmt.Println(`Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]`)
		matrix = S2SoSliceInt(`[[1,1,1],[1,0,1],[1,1,1]]`)
		setZeroes(matrix)

		fmt.Println(`Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]`)
		matrix = S2SoSliceInt(`[[0,1,2,0],[3,4,5,2],[1,3,1,5]]`)
		setZeroes(matrix)
	}
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
