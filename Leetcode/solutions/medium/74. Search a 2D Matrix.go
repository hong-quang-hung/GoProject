package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/search-a-2d-matrix/
func init() {
	Solutions[74] = func() {
		fmt.Println(`Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3`)
		fmt.Println(`Output:`, searchMatrix(S2SoSliceInt(`[[1,3,5,7],[10,11,16,20],[23,30,34,60]]`), 3))
		fmt.Println(`Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13`)
		fmt.Println(`Output:`, searchMatrix(S2SoSliceInt(`[[1,3,5,7],[10,11,16,20],[23,30,34,60]]`), 13))
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	rl, rr := 0, m-1
	for rl <= rr {
		rm := (rl + rr) / 2
		if matrix[rm][0] > target {
			rr = rm - 1
		} else if matrix[rm][n-1] < target {
			rl = rm + 1
		} else {
			cm := sort.SearchInts(matrix[rm], target)
			return matrix[rm][cm] == target
		}
	}
	return false
}
