package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/search-a-2d-matrix-ii/
func init() {
	Solutions[240] = func() {
		fmt.Println("Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5")
		fmt.Println("Output:", searchMatrixII(S2SoSliceInt("[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"), 5))
		fmt.Println("Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20")
		fmt.Println("Output:", searchMatrixII(S2SoSliceInt("[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]"), 20))
	}
}

func searchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	rl, rr, cl, cr := 0, m-1, 0, n-1
	for rl <= rr && cl <= cr {
		rm := (rl + rr) / 2
		cm := (cl + cr) / 2
		if matrix[rm][cm] == target {
			return true
		} else if matrix[rm][cm] > target {
			for matrix[rm][cl] < target {
				cl++
			}

			for matrix[rl][cm] < target {
				rl++
			}
		} else {
			for matrix[rm][cr] > target {
				cr++
			}

			for matrix[rr][cm] > target {
				rr--
			}
		}
	}
	return false
}
