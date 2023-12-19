package medium

import "fmt"

// Reference: https://leetcode.com/problems/01-matrix/
func init() {
	Solutions[542] = func() {
		fmt.Println(`Input: mat = [[0,0,0],[0,1,0],[0,0,0]]`)
		fmt.Println(`Output:`, updateMatrix(S2SoSliceInt(`[[0,0,0],[0,1,0],[0,0,0]]`)))
		fmt.Println(`Input: mat = [[0,0,0],[0,1,0],[1,1,1]]`)
		fmt.Println(`Output:`, updateMatrix(S2SoSliceInt(`[[0,0,0],[0,1,0],[1,1,1]]`)))
	}
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	queue := [][]int{}
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				res[i][j] = 100000001
			}
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, dir := range BoardDirection {
			r, c := cell[0]+dir[0], cell[1]+dir[1]
			if r >= 0 && r < m && c >= 0 && c < n && res[r][c] > res[cell[0]][cell[1]]+1 {
				queue = append(queue, []int{r, c})
				res[r][c] = res[cell[0]][cell[1]] + 1
			}
		}
	}
	return res
}
