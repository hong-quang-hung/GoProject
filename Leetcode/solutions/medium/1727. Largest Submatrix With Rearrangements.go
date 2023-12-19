package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
func init() {
	Solutions[1727] = func() {
		fmt.Println(`Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]`)
		fmt.Println(`Output:`, largestSubmatrix(S2SoSliceInt(`[[0,0,1],[1,1,1],[1,0,1]]`)))
		fmt.Println(`Input: matrix = [[1,1,0],[1,0,1]]`)
		fmt.Println(`Output:`, largestSubmatrix(S2SoSliceInt(`[[1,1,0],[1,0,1]]`)))
	}
}

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] > 0 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}

	res := 0
	for _, row := range matrix {
		sort.Ints(row)
		for j := n - 1; j >= 0; j-- {
			width, height := n-j, row[j]
			area := width * height
			res = max(res, area)
		}
	}
	return res
}
