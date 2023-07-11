package medium

import "fmt"

// Reference: https://leetcode.com/problems/sum-of-matrix-after-queries/
func init() {
	Solutions[2718] = func() {
		fmt.Println("Input: n = 3, queries = [[0,0,1],[1,2,2],[0,2,3],[1,0,4]]")
		fmt.Println("Output:", matrixSumQueries(3, S2SoSliceInt("[[0,0,1],[1,2,2],[0,2,3],[1,0,4]]")))
		fmt.Println("Input: n = 3, queries = [[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]")
		fmt.Println("Output:", matrixSumQueries(3, S2SoSliceInt("[[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]")))
	}
}

func matrixSumQueries(n int, queries [][]int) int64 {
	m := len(queries)
	col, row := make(map[int]bool), make(map[int]bool)
	res := int64(0)
	for i := m - 1; i >= 0; i-- {
		t, index, val := queries[i][0], queries[i][1], queries[i][2]
		if t == 0 {
			if !row[index] {
				row[index] = true
				res += int64(val) * int64(n)
				res -= int64(val) * int64(len(col))
			}
		} else {
			if !col[index] {
				col[index] = true
				res += int64(val) * int64(n)
				res -= int64(val) * int64(len(row))
			}
		}
	}
	return res
}
