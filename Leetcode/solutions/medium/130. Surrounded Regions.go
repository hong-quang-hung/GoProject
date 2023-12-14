package medium

import "fmt"

// Reference: https://leetcode.com/problems/surrounded-regions/
func init() {
	Solutions[130] = func() {
		var grid [][]byte

		fmt.Println("Input: board = [[\"X\",\"X\",\"X\",\"X\"],[\"X\",\"O\",\"O\",\"X\"],[\"X\",\"X\",\"O\",\"X\"],[\"X\",\"O\",\"X\",\"X\"]]")
		grid = [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
		solve(grid)
		fmt.Println("Output:", grid)
	}
}

func solve(board [][]byte) {
	lnY, lnX := len(board), len(board[0])
	if lnY == 1 && lnX == 1 {
		return
	}

	visited := make([][]bool, lnY)
	for i := 0; i < lnY; i++ {
		visited[i] = make([]bool, lnX)
	}

	dirs := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := 0; i < lnY; i++ {
		for j := 0; j < lnX; j++ {
			if visited[i][j] {
				if i > 0 && i < lnY-1 {
					j += lnX - 2
				}
				continue
			}

			visited[i][j] = true
			if board[i][j] == 'O' {
				queue := [][2]int{{i, j}}
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]

					for _, dir := range dirs {
						y, x := cur[0]+dir[0], cur[1]+dir[1]
						if y >= 0 && y < lnY && x >= 0 && x < lnX && !visited[y][x] && board[y][x] == 'O' {
							queue = append(queue, [2]int{y, x})
							visited[y][x] = true
						}
					}
				}
			}
			if i > 0 && i < lnY-1 {
				j += lnX - 2
			}
		}
	}

	for i := 1; i < lnY-1; i++ {
		for j := 1; j < lnX-1; j++ {
			if !(visited[i][j]) {
				board[i][j] = 'X'
			}
		}
	}
}
