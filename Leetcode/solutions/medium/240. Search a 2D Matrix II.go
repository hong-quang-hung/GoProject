package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/search-a-2d-matrix-ii/
func init() {
	Solutions[240] = func() {
		fmt.Println(`Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5`)
		fmt.Println(`Output:`, searchMatrixII(S2SoSliceInt(`[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]`), 5))
		fmt.Println(`Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20`)
		fmt.Println(`Output:`, searchMatrixII(S2SoSliceInt(`[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]`), 20))
	}
}

func searchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	r, c := 0, n-1
	for r < m && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			r++
		} else {
			c--
		}
	}
	return false
}
