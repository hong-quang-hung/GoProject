package medium

import "fmt"

// Reference: https://leetcode.com/problems/rotting-oranges/
func init() {
	Solutions[944] = func() {
		fmt.Println("Input: grid = [[2,1,1],[1,1,0],[0,1,1]]")
		fmt.Println("Output:", orangesRotting(S2SoSliceInt("[[2,1,1],[1,1,0],[0,1,1]]")))
		fmt.Println("Input: grid = [[2,1,1],[0,1,1],[1,0,1]]")
		fmt.Println("Output:", orangesRotting(S2SoSliceInt("[[2,1,1],[0,1,1],[1,0,1]]")))
		fmt.Println("Input: grid = [[0,2]]")
		fmt.Println("Output:", orangesRotting(S2SoSliceInt("[[0,2]]")))
		fmt.Println("Input: grid = [[2,2],[1,1],[0,0],[2,0]]")
		fmt.Println("Output:", orangesRotting(S2SoSliceInt("[[2,2],[1,1],[0,0],[2,0]]")))
	}
}

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	n, m := len(grid), len(grid[0])
	rot := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				rot++
			}
		}
	}

	if rot == 0 {
		return 0
	}

	res := -1
	rot += len(queue)
	direction := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	for len(queue) > 0 {
		temp := [][]int{}
		res++
		for _, next := range queue {
			rot--
			for _, dir := range direction {
				row, col := next[0]+dir[0], next[1]+dir[1]
				if row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == 1 {
					grid[row][col] = 2
					temp = append(temp, []int{row, col})
				}

			}
		}
		queue = temp
	}

	if rot > 0 {
		return -1
	}
	return res
}
