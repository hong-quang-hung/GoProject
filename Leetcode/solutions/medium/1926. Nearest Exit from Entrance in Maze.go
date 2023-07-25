package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-eventual-safe-states/
func init() {
	Solutions[1926] = func() {
		fmt.Println("Input: maze = [['+','+','.','+'],['.','.','.','+'],['+','+','+','.']], entrance = [1,2]")
		fmt.Println("Output:", nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
		fmt.Println("Input: maze = [['+','+','+'],['.','.','.'],['+','+','+']], entrance = [1,0]")
		fmt.Println("Output:", nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))
		fmt.Println("Input: maze = [['.','+']], entrance = [0,0]")
		fmt.Println("Output:", nearestExit([][]byte{{'.', '+'}}, []int{0, 0}))
	}
}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	res := 0
	queue := [][]int{{entrance[0], entrance[1]}}
	for len(queue) > 0 {
		temp := [][]int{}
		for _, next := range queue {
			for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				row, col := next[0]+dir[0], next[1]+dir[1]
				if row >= 0 && row < n && col >= 0 && col < m && maze[row][col] == '.' {
					if row == 0 || row == n || col == 0 || col == m {
						return res
					}

					res++
					temp = append(temp, []int{row, col})
				}
			}
		}
		queue = temp
	}
	return -1
}
