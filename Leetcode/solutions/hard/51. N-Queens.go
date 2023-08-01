package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/n-queens/
func init() {
	Solutions[51] = func() {
		fmt.Println("Input: n = 4")
		fmt.Println("Output:", solveNQueens(4))
		fmt.Println("Input: n = 1")
		fmt.Println("Output:", solveNQueens(1))
	}
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	queens := make(map[int]int)
	solveNQueensBackTracking(n, &res, 0, &queens)
	return res
}

func solveNQueensBackTracking(n int, res *[][]string, row int, queens *map[int]int) {
	if row == n {
		rowString := []string{}
		for i := 0; i < n; i++ {
			s := make([]byte, n)
			for r := 0; r < n; r++ {
				if r == (*queens)[i] {
					s[r] = 'Q'
				} else {
					s[r] = '.'
				}
			}
			rowString = append(rowString, string(s))
		}
		*res = append(*res, rowString)
	}

	for col := 0; col < n; col++ {
		isValid := true
		for r, c := range *queens {
			if c == col || row-r == col-c || row-r == c-col {
				isValid = false
				break
			}
		}

		if isValid {
			(*queens)[row] = col
			solveNQueensBackTracking(n, res, row+1, queens)
			delete(*queens, row)
		}
	}
}
