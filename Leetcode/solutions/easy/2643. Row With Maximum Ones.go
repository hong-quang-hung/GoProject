package easy

import "fmt"

// Reference: https://leetcode.com/problems/row-with-maximum-ones/
func Leetcode_Row_And_Maximum_Ones() {
	fmt.Println("Input: mat = [[0,1],[1,0]]")
	fmt.Println("Output:", rowAndMaximumOnes(S2SoSliceInt("[[0,1],[1,0]]")))
	fmt.Println("Input: mat = [[0,0,0],[0,1,1]]")
	fmt.Println("Output:", rowAndMaximumOnes(S2SoSliceInt("[[0,0,0],[0,1,1]]")))
	fmt.Println("Input: mat = [[0,0],[1,1],[0,0]]")
	fmt.Println("Output:", rowAndMaximumOnes(S2SoSliceInt("[[0,0],[1,1],[0,0]]")))
}

func rowAndMaximumOnes(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	res := make([]int, 2)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				count++
			}
		}
		if count > res[1] {
			res[0] = i
			res[1] = count
		}
	}
	return res
}
