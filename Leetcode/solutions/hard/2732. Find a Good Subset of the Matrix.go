package hard

import "fmt"

// Reference: https://leetcode.com/problems/find-a-good-subset-of-the-matrix/
func init() {
	Solutions[2732] = func() {
		fmt.Println("Input: grid = [[0,1,1,0],[0,0,0,1],[1,1,1,1]]")
		fmt.Println("Output:", goodSubsetofBinaryMatrix(S2SoSliceInt("[[0,1,1,0],[0,0,0,1],[1,1,1,1]]")))
		fmt.Println("Input: grid = [[0]]")
		fmt.Println("Output:", goodSubsetofBinaryMatrix(S2SoSliceInt("[[0]]")))
		fmt.Println("Input: grid = [[1,1,1],[1,1,1]]")
		fmt.Println("Output:", goodSubsetofBinaryMatrix(S2SoSliceInt("[[1,1,1],[1,1,1]]")))
	}
}

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	hs := make(map[int]int)
	for i := 0; i < n; i++ {
		val := 0
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				val += (1 << j)
			}
		}

		if val == 0 {
			return []int{i}
		}

		for j := 1; j < 32; j++ {
			if val&j == 0 {
				if _, c := hs[j]; c {
					return []int{hs[j], i}
				}
			}
		}

		hs[val] = i
	}
	return []int{}
}
