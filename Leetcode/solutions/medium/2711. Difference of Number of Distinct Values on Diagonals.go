package medium

import "fmt"

// Reference: https://leetcode.com/problems/difference-of-number-of-distinct-values-on-diagonals/
func init() {
	Solutions[2711] = func() {
		fmt.Println("Input: grid = [[1,2,3],[3,1,5],[3,2,1]]")
		fmt.Println("Output:", differenceOfDistinctValues(S2SoSliceInt("[[1,2,3],[3,1,5],[3,2,1]]")))
		fmt.Println("Input: grid = [[1]]")
		fmt.Println("Output:", differenceOfDistinctValues(S2SoSliceInt("[[1]]")))
	}
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0])
	res := make([][]int, n)
	for r := 0; r < n; r++ {
		res[r] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x, y, p, q := i+1, j+1, i-1, j-1
			BR, TL := make(map[int]int), make(map[int]int)
			for x < n && y < m {
				BR[grid[x][y]]++
				x++
				y++
			}

			for p >= 0 && q >= 0 {
				TL[grid[p][q]]++
				p--
				q--
			}

			res[i][j] = abs(len(BR) - len(TL))
		}
	}
	return res
}
