package hard

import "fmt"

// Reference: https://leetcode.com/problems/n-queens-ii/
func init() {
	Solutions[52] = func() {
		fmt.Println(`Input: n = 4`)
		fmt.Println(`Output:`, totalNQueens(4))
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, totalNQueens(1))
	}
}

func totalNQueens(n int) int {
	res := 0
	var f func(queen *map[int]int, row int)
	f = func(queen *map[int]int, row int) {
		if row == n {
			res++
		}

		for col := 0; col < n; col++ {
			isValid := true
			for r, c := range *queen {
				if c == col || row-r == col-c || row-r == c-col {
					isValid = false
					break
				}
			}

			if isValid {
				(*queen)[row] = col
				f(queen, row+1)
				delete(*queen, row)
			}
		}
	}

	queen := make(map[int]int)
	f(&queen, 0)
	return res
}
