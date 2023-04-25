package easy

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/special-positions-in-a-binary-matrix/
func Leetcode_Num_Special() {
	// fmt.Println("Input: mat = [[1,0,0],[0,0,1],[1,0,0]]")
	// fmt.Println("Output:", numSpecial(utils.S2SoSliceInt("[[1,0,0],[0,0,1],[1,0,0]]")))
	// fmt.Println("Input: mat = [[1,0,0],[0,1,0],[0,0,1]]")
	// fmt.Println("Output:", numSpecial(utils.S2SoSliceInt("[[1,0,0],[0,1,0],[0,0,1]]")))
	// fmt.Println("Input: mat = [[0,0,0,0,0,1,0,0],[0,0,0,0,1,0,0,1],[0,0,0,0,1,0,0,0],[1,0,0,0,1,0,0,0],[0,0,1,1,0,0,0,0]]")
	// fmt.Println("Output:", numSpecial(utils.S2SoSliceInt("[[0,0,0,0,0,1,0,0],[0,0,0,0,1,0,0,1],[0,0,0,0,1,0,0,0],[1,0,0,0,1,0,0,0],[0,0,1,1,0,0,0,0]]")))
	fmt.Println("Input: mat = [[0,0,0,0,0,0,0,0],[0,0,0,1,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,1,0],[0,1,0,0,0,0,1,0],[0,1,0,0,0,0,0,0]]")
	fmt.Println("Output:", numSpecial(utils.S2SoSliceInt("[[0,0,0,0,0,0,0,0],[0,0,0,1,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,1,0],[0,1,0,0,0,0,1,0],[0,1,0,0,0,0,0,0]]")))
}

func numSpecial(mat [][]int) int {
	res, n, m := 0, len(mat), len(mat[0])
	r := make([]int, n)
	c := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				r[i]++
				c[j]++
				res++

				if r[i] >= 2 {
					res--
				}

				if c[j] >= 2 {
					res--
				}
			}
		}
	}
	return res
}
