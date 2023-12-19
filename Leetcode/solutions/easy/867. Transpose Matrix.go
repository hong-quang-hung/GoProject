package easy

import "fmt"

// Reference: https://leetcode.com/problems/transpose-matrix/
func init() {
	Solutions[867] = func() {
		fmt.Println(`Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]`)
		fmt.Println(`Output:`, transpose(S2SoSliceInt(`[[1,2,3],[4,5,6],[7,8,9]]`)))
		fmt.Println(`Input: matrix = [[1,2,3],[4,5,6]]`)
		fmt.Println(`Output:`, transpose(S2SoSliceInt(`[[1,2,3],[4,5,6]]`)))
	}
}

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
