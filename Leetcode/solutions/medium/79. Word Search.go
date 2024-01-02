package medium

import "fmt"

// Reference: https://leetcode.com/problems/word-search/
func init() {
	Solutions[79] = func() {
		fmt.Println(`Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"`)
		fmt.Println(`Output:`, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, `ABCCED`))
		fmt.Println(`Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"`)
		fmt.Println(`Output:`, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, `SEE`))
		fmt.Println(`Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"`)
		fmt.Println(`Output:`, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, `ABCB`))
	}
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var f func(r, c, i int) bool
	f = func(r, c, i int) bool {
		if i >= len(word) {
			return true
		}

		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[i] || board[r][c] == '*' {
			return false
		}

		visited := board[r][c]
		board[r][c] = '*'
		res := f(r-1, c, i+1) || f(r+1, c, i+1) || f(r, c-1, i+1) || f(r, c+1, i+1)
		board[r][c] = visited
		return res
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if f(r, c, 0) {
				return true
			}
		}
	}
	return false
}
