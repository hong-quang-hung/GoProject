package medium

import "fmt"

// Reference: https://leetcode.com/problems/rotate-image/
func init() {
	Solutions[48] = func() {
		fmt.Println("Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]")
		matrix := S2SoSliceInt("[[1,2,3],[4,5,6],[7,8,9]]")
		rotate(matrix)
		fmt.Println("Output:", matrix)
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
