package medium

import "fmt"

func init() {
	Solutions[2661] = Leetcode_First_Complete_Index
}

// Reference: https://leetcode.com/problems/first-completely-painted-row-or-column/
func Leetcode_First_Complete_Index() {
	fmt.Println("Input: arr = [1,3,4,2], mat = [[1,4],[2,3]]")
	fmt.Println("Output:", firstCompleteIndex([]int{1, 3, 4, 2}, S2SoSliceInt("[[1,4],[2,3]]")))
	fmt.Println("Input: arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]")
	fmt.Println("Output:", firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, S2SoSliceInt("[[3,2,5],[1,4,6],[8,7,9]]")))
	fmt.Println("Input: arr = [1,4,5,2,6,3], mat = [[4,3,5],[1,2,6]]")
	fmt.Println("Output:", firstCompleteIndex([]int{1, 4, 5, 2, 6, 3}, S2SoSliceInt("[[4,3,5],[1,2,6]]")))
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	n, m := len(mat), len(mat[0])
	row := make([]int, n)
	col := make([]int, m)
	hm := make(map[int][2]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			hm[mat[i][j]] = [2]int{i, j}
		}
	}

	for i := 0; i < len(arr); i++ {
		k := hm[arr[i]]
		row[k[0]]++
		if row[k[0]] == m {
			return i
		}

		col[k[1]]++
		if col[k[1]] == n {
			return i
		}
	}
	return 0
}
