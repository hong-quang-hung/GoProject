package easy

import "fmt"

// Reference: https://leetcode.com/problems/special-positions-in-a-binary-matrix/
func init() {
	Solutions[1582] = func() {
		fmt.Println(`Input: mat = [[1,0,0],[0,0,1],[1,0,0]]`)
		fmt.Println(`Output:`, numSpecial(S2SoSliceInt(`[[1,0,0],[0,0,1],[1,0,0]]`)))
		fmt.Println(`Input: mat = [[1,0,0],[0,1,0],[0,0,1]]`)
		fmt.Println(`Output:`, numSpecial(S2SoSliceInt(`[[1,0,0],[0,1,0],[0,0,1]]`)))
		fmt.Println(`Input: mat = [[0,0,0,0,0,1,0,0],[0,0,0,0,1,0,0,1],[0,0,0,0,1,0,0,0],[1,0,0,0,1,0,0,0],[0,0,1,1,0,0,0,0]]`)
		fmt.Println(`Output:`, numSpecial(S2SoSliceInt(`[[0,0,0,0,0,1,0,0],[0,0,0,0,1,0,0,1],[0,0,0,0,1,0,0,0],[1,0,0,0,1,0,0,0],[0,0,1,1,0,0,0,0]]`)))
	}
}

func numSpecial(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	r := make([]int, n)
	c := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				r[i]++
				c[j]++
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 && r[i] == 1 && c[j] == 1 {
				res++
			}
		}
	}
	return res
}
