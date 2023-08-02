package medium

import "fmt"

// Reference: https://leetcode.com/problems/word-search/
func init() {
	Solutions[79] = func() {
		fmt.Println("Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCCED'")
		fmt.Println("Output:", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
		fmt.Println("Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'SEE'")
		fmt.Println("Output:", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
		fmt.Println("Input: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCB'")
		fmt.Println("Output:", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
	}
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existCheck(board, word, visited, m, n, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func existCheck(board [][]byte, word string, visited [][]bool, m, n, row, col, idx int) bool {
	if idx == len(word) {
		return true
	}

	if word[idx] != board[row][col] {
		return false
	} else if idx == len(word)-1 {
		return true
	}

	visited[row][col] = true
	for _, dir := range BoardDirection {
		nextRow, nextCol := row+dir[0], col+dir[1]
		if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n {
			if !visited[nextRow][nextCol] && existCheck(board, word, visited, m, n, nextRow, nextCol, idx+1) {
				return true
			}
			visited[nextRow][nextCol] = false
		}
	}

	visited[row][col] = false
	return false
}
