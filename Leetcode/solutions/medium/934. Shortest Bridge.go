package medium

import "fmt"

// Reference: https://leetcode.com/problems/shortest-bridge/
func init() {
	Solutions[934] = func() {
		fmt.Println("Input: grid = [[0,1],[1,0]]")
		fmt.Println("Output:", shortestBridge(S2SoSliceInt("[[0,1],[1,0]]")))
		fmt.Println("Input: grid = [[0,1,0],[0,0,0],[0,0,1]]")
		fmt.Println("Output:", shortestBridge(S2SoSliceInt("[[0,1,0],[0,0,0],[0,0,1]]")))
		fmt.Println("Input: grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]")
	}
}

func shortestBridge(grid [][]int) int {
	n := len(grid)
	firstX, firstY := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				firstX = i
				firstY = j
				break
			}
		}
	}

	bfsQueue := make([][]int, 0)
	shortestBridgeDFS(&grid, &bfsQueue, firstX, firstY, n)

	distance := 0
	for len(bfsQueue) != 0 {
		newBfs := make([][]int, 0)
		for _, pair := range bfsQueue {
			x, y := pair[0], pair[1]
			for _, nextPair := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
				curX, curY := nextPair[0], nextPair[1]
				if curX >= 0 && curX < n && curY >= 0 && curY < n {
					if grid[curX][curY] == 1 {
						return distance
					} else if grid[curX][curY] == 0 {
						newBfs = append(newBfs, nextPair)
						grid[curX][curY] = -1
					}
				}
			}
		}

		bfsQueue = newBfs
		distance++
	}
	return distance
}

func shortestBridgeDFS(grid *[][]int, bfsQueue *[][]int, x, y, n int) {
	(*grid)[x][y] = 2
	*bfsQueue = append(*bfsQueue, []int{x, y})
	for _, pair := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
		curX, curY := pair[0], pair[1]
		if 0 <= curX && curX < n && 0 <= curY && curY < n && (*grid)[curX][curY] == 1 {
			shortestBridgeDFS(grid, bfsQueue, curX, curY, n)
		}
	}
}
