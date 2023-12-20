package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-eventual-safe-states/
func init() {
	Solutions[1926] = func() {
		fmt.Println(`Input: maze = [['+','+','.','+'],['.','.','.','+'],['+','+','+','.']], entrance = [1,2]`)
		fmt.Println(`Output:`, nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
		fmt.Println(`Input: maze = [['+','+','+'],['.','.','.'],['+','+','+']], entrance = [1,0]`)
		fmt.Println(`Output:`, nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))
		fmt.Println(`Input: maze = [{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '+', '.', '+'}, {'+', '+', '+', '+', '+', '.', '+'}], entrance = [3,2]`)
		fmt.Println(`Output:`, nearestExit([][]byte{{'+', '.', '+', '+', '+', '+', '+'}, {'+', '.', '+', '.', '.', '.', '+'}, {'+', '.', '+', '.', '+', '.', '+'}, {'+', '.', '.', '.', '+', '.', '+'}, {'+', '+', '+', '+', '+', '.', '+'}}, []int{3, 2}))
	}
}

func nearestExit(maze [][]byte, entrance []int) int {
	n, m := len(maze), len(maze[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	res := 0
	queue := [][]int{{entrance[0], entrance[1]}}
	directs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) > 0 {
		temp := [][]int{}
		res++
		for _, next := range queue {
			if visited[next[0]][next[1]] {
				continue
			}

			visited[next[0]][next[1]] = true
			for _, dir := range directs {
				row, col := next[0]+dir[0], next[1]+dir[1]
				if row >= 0 && row < n && col >= 0 && col < m && !visited[row][col] && maze[row][col] == '.' {
					if row == 0 || row == n-1 || col == 0 || col == m-1 {
						return res
					}

					temp = append(temp, []int{row, col})
				}
			}
		}
		queue = temp
	}
	return -1
}
