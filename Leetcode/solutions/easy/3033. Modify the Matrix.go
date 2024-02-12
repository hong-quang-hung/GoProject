package easy

import "fmt"

// Reference: https://leetcode.com/problems/modify-the-matrix/
func init() {
	Solutions[3033] = func() {
		fmt.Println(`Input: matrix = [[1,2,-1],[4,-1,6],[7,8,9]]`)
		fmt.Println(`Output:`, modifiedMatrix(S2SoSliceInt("[[1,2,-1],[4,-1,6],[7,8,9]]")))
		fmt.Println(`Input: matrix = [[3,-1],[5,2]]`)
		fmt.Println(`Output:`, modifiedMatrix(S2SoSliceInt("[[3,-1],[5,2]]")))
	}
}

func modifiedMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	maxCols := make([]int, n)
	for i := 0; i < n; i++ {
		maxCols[i] = -2
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == -1 {
				if maxCols[j] == -2 {
					maxVal := -2
					for c := 0; c < m; c++ {
						maxVal = max(maxVal, matrix[c][j])
					}
					maxCols[j] = maxVal
				}
				matrix[i][j] = maxCols[j]
			}
		}
	}
	return matrix
}
