package hard

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/sudoku-solver/
func init() {
	Solutions[37] = func() {
		fmt.Println(`Input: board = [['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]`)
		board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
		solveSudoku(board)
		fmt.Println(`Output:`, board)
	}
}

func solveSudoku(board [][]byte) {
	rows := [9][10]bool{}
	cols := [9][10]bool{}
	grid := [9][10]bool{}
	dots := make([][]int, 0)
	b := [9][9]int{}

	convert := func(b byte) int {
		if b == '.' {
			return 0
		}
		v, _ := strconv.Atoi(string(b))
		return v
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[i][j] = convert(board[i][j])
			if b[i][j] == 0 {
				dots = append(dots, []int{i, j})
				continue
			}
			rows[i][b[i][j]] = true
			cols[j][b[i][j]] = true
			grid[i/3*3+j/3][b[i][j]] = true
		}
	}

	m := len(dots)
	var f func(int) bool
	f = func(step int) bool {
		if step == m {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					board[i][j] = fmt.Sprint(b[i][j])[0]
				}
			}
			return true
		}

		x, y := dots[step][0], dots[step][1]
		for i := 1; i < 10; i++ {
			if !rows[x][i] && !cols[y][i] && !grid[x/3*3+y/3][i] {
				b[x][y] = i
				rows[x][i], cols[y][i], grid[x/3*3+y/3][i] = true, true, true
				if f(step + 1) {
					return true
				}
				rows[x][i], cols[y][i], grid[x/3*3+y/3][i] = false, false, false
			}
		}
		return false
	}
	f(0)
}
