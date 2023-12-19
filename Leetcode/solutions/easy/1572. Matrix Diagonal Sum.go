package easy

import "fmt"

// Reference: https://leetcode.com/problems/matrix-diagonal-sum/
func init() {
	Solutions[1572] = func() {
		fmt.Println(`Input: mat = [[1,2,3],[4,5,6],[7,8,9]]`)
		fmt.Println(`Output:`, diagonalSum(S2SoSliceInt(`[[1,2,3],[4,5,6],[7,8,9]]`)))
		fmt.Println(`Input: mat = [[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]`)
		fmt.Println(`Output:`, diagonalSum(S2SoSliceInt(`[[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]`)))
	}
}

func diagonalSum(mat [][]int) int {
	n, res := len(mat), 0
	for i := 0; i < n; i++ {
		res += mat[i][i] + mat[n-1-i][i]
	}

	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}
