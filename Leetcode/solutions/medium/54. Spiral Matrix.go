package medium

import "fmt"

// Reference: https://leetcode.com/problems/spiral-matrix/
func init() {
	Solutions[54] = func() {
		fmt.Println(`Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]`)
		fmt.Println(`Output:`, spiralOrder(S2SoSliceInt(`[[1,2,3],[4,5,6],[7,8,9]]`)))
		fmt.Println(`Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]`)
		fmt.Println(`Output:`, spiralOrder(S2SoSliceInt(`[[1,2,3,4],[5,6,7,8],[9,10,11,12]]`)))
	}
}

func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1
	res := []int{}
	for len(res) != rows*cols {
		for i := left; i < right+1; i++ {
			res = append(res, matrix[top][i])
		}

		top++
		for i := top; i < bottom+1; i++ {
			res = append(res, matrix[i][right])
		}

		right--
		if top <= bottom {
			for i := right; i > left-1; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i > top-1; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
