package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/spiral-matrix/
func Leetcode_Spiral_Order() {
	fmt.Println("Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]")
	fmt.Println("Output:", spiralOrder(utils.S2SoSliceInt("[[1,2,3],[4,5,6],[7,8,9]]")))
	fmt.Println("Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]")
	fmt.Println("Output:", spiralOrder(utils.S2SoSliceInt("[[1,2,3,4],[5,6,7,8],[9,10,11,12]]")))
}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	res := make([]int, n*m)

	r, c := -1, -1
	isNextR, isNextC := true, true
	limitR, LimitR := n-1, m-1
	for i := range res {
		if isNextR {
			r++
		} else {
			r--
		}

		if isNextC {
			c++
		} else {
			c--
		}

		res[i] = matrix[r][c]
	}
	return res
}
