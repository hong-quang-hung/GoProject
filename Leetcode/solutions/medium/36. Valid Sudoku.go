package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/valid-sudoku/
func Leetcode_Is_Valid_Sudoku() {
	fmt.Println("Input: board = ")
	fmt.Println("[['5','3','.','.','7','.','.','.','.']")
	fmt.Println(",['6','.','.','1','9','5','.','.','.']")
	fmt.Println(",['.','9','8','.','.','.','.','6','.']")
	fmt.Println(",['8','.','.','.','6','.','.','.','3']")
	fmt.Println(",['4','.','.','8','.','3','.','.','1']")
	fmt.Println(",['7','.','.','.','2','.','.','.','6']")
	fmt.Println(",['.','6','.','.','.','.','2','8','.']")
	fmt.Println(",['.','.','.','4','1','9','.','.','5']")
	fmt.Println(",['.','.','.','.','8','.','.','7','9']]")

	board := make([][]byte, 9)
	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	fmt.Println("Output:", isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if row[board[i][j]] {
				return false
			}
			row[board[i][j]] = true
		}

		col := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if col[board[j][i]] {
				return false
			}
			col[board[j][i]] = true
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			group := make(map[byte]bool)
			for g := 0; g < 9; g++ {
				r, c := i+g%3, j+g/3
				if board[r][c] == '.' {
					continue
				}
				if group[board[r][c]] {
					return false
				}
				group[board[r][c]] = true
			}
		}
	}
	return true
}
