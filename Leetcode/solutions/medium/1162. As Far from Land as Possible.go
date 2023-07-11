package medium

import "fmt"

// Reference: https://leetcode.com/problems/as-far-from-land-as-possible/
func init() {
	Solutions[1162] = func() {
		fmt.Println("Input: grid = [[1,0,0],[0,0,0],[0,0,0]]")
		fmt.Println("Output:", maxVDistance(S2SoSliceInt("[[1,0,0],[0,0,0],[0,0,0]]")))
	}
}

func maxVDistance(grid [][]int) int {
	maxV := -1
	n, maxV := len(grid)-1, -1
	queue := [][2]int{}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		land := queue[0]
		queue = queue[1:]
		point := grid[land[0]][land[1]]
		col, row := land[0], land[1]

		if col > 0 && grid[col-1][row] == 0 {
			grid[col-1][row] = point + 1
			if maxV < point {
				maxV = point
			}
			queue = append(queue, [2]int{col - 1, row})
		}

		if col < n && grid[col+1][row] == 0 {
			grid[col+1][row] = point + 1
			if maxV < point {
				maxV = point
			}
			queue = append(queue, [2]int{col + 1, row})
		}

		if row > 0 && grid[col][row-1] == 0 {
			grid[col][row-1] = point + 1
			if maxV < point {
				maxV = point
			}
			queue = append(queue, [2]int{col, row - 1})
		}

		if row < n && grid[col][row+1] == 0 {
			grid[col][row+1] = point + 1
			if maxV < point {
				maxV = point
			}
			queue = append(queue, [2]int{col, row + 1})
		}
	}
	return maxV
}
