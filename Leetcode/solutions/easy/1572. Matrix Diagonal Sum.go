package easy

import "fmt"

func init() {
	Solutions[1572] = Leetcode_Diagonal_Sum
}

// Reference: https://leetcode.com/problems/matrix-diagonal-sum/
func Leetcode_Diagonal_Sum() {
	fmt.Println("Input: mat = [[1,2,3],[4,5,6],[7,8,9]]")
	fmt.Println("Output:", diagonalSum(S2SoSliceInt("[[1,2,3],[4,5,6],[7,8,9]]")))
	fmt.Println("Input: mat = [[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]")
	fmt.Println("Output:", diagonalSum(S2SoSliceInt("[[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]")))
	fmt.Println("Input: mat = [[5]]")
	fmt.Println("Output:", diagonalSum(S2SoSliceInt("[[5]]")))
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
