package medium

import "fmt"

// Reference: https://leetcode.com/problems/game-of-life/
func init() {
	Solutions[289] = func() {
		var board [][]int

		fmt.Println(`Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]`)
		board = S2SoSliceInt(`[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]`)
		gameOfLife(board)
		fmt.Println(`Output:`, board)

		fmt.Println(`Input: board = [[1,1],[1,0]]`)
		board = S2SoSliceInt(`[[1,1],[1,0]]`)
		gameOfLife(board)
		fmt.Println(`Output:`, board)
	}
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			neighbor := gameOfLifeCount(board, i, j)
			if board[i][j] == 1 {
				if neighbor == 2 || neighbor == 3 {
					board[i][j] += 2
				}
			} else if board[i][j] == 0 {
				if neighbor == 3 {
					board[i][j] += 2
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func gameOfLifeCount(board [][]int, r, c int) int {
	count := 0
	for i := max(0, r-1); i <= min(r+1, len(board)-1); i++ {
		for j := max(0, c-1); j <= min(c+1, len(board[0])-1); j++ {
			if i == r && j == c {
				continue
			}
			if board[i][j]&1 == 1 {
				count++
			}
		}
	}
	return count
}
