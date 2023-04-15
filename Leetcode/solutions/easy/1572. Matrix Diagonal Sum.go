package easy

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/matrix-diagonal-sum/
func Leetcode_Diagonal_Sum() {
	fmt.Println("Input: mat = [[1,2,3],[4,5,6],[7,8,9]]")
	fmt.Println("Output:", diagonalSum(utils.S2SoSliceInt("[[1,2,3],[4,5,6],[7,8,9]]")))
	fmt.Println("Input: mat = [[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]")
	fmt.Println("Output:", diagonalSum(utils.S2SoSliceInt("[[1,1,1,1],[1,1,1,1],[1,1,1,1],[1,1,1,1]]")))
}

func diagonalSum(mat [][]int) int {
	n := len(mat)
	move := [][]int{{0, 0}, {0, n - 1}}
	for _, m := range move {
		i, j := m[0], m[1]
	}
	return 1
}
