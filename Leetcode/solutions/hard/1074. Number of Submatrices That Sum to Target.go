package hard

import "fmt"

// Reference: https://leetcode.com/problems/k-inverse-pairs-array/
func init() {
	Solutions[1074] = func() {
		fmt.Println(`Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0`)
		fmt.Println(`Output:`, numSubmatrixSumTarget(S2SoSliceInt("[[0,1,0],[1,1,1],[0,1,0]]"), 0))
		fmt.Println(`Input: matrix = [[1,-1],[-1,1]], target = 0`)
		fmt.Println(`Output:`, numSubmatrixSumTarget(S2SoSliceInt("[[1,-1],[-1,1]]"), 0))
	}
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	suffix := make([][]int, m)
	for i := 0; i < m; i++ {
		suffix[i] = make([]int, n)
	}

	suffix[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		suffix[i][0] = suffix[i-1][0] + matrix[i][0]
	}

	for i := 1; i < n; i++ {
		suffix[0][i] = suffix[0][i-1] + matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			suffix[i][j] = matrix[i][j] + suffix[i-1][j] + suffix[i][j-1] - suffix[i-1][j-1]
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			count := make(map[int]int)
			count[0] = 1
			for k := 0; k < n; k++ {
				cur_suffix := suffix[j][k]
				if i > 0 {
					cur_suffix -= suffix[i-1][k]
				}

				diff := cur_suffix - target
				res += count[diff]
				count[cur_suffix]++
			}
		}
	}
	return res
}
