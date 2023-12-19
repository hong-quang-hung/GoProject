package hard

import "fmt"

// Reference: https://leetcode.com/problems/last-day-where-you-can-still-cross/
func init() {
	Solutions[1970] = func() {
		fmt.Println(`Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]`)
		fmt.Println(`Output:`, latestDayToCross(2, 2, S2SoSliceInt(`[[1,1],[2,1],[1,2],[2,2]]`)))
		fmt.Println(`Input: row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]`)
		fmt.Println(`Output:`, latestDayToCross(2, 2, S2SoSliceInt(`[[1,1],[1,2],[2,1],[2,2]]`)))
		fmt.Println(`Input: row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]`)
		fmt.Println(`Output:`, latestDayToCross(3, 3, S2SoSliceInt(`[[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]`)))
	}
}

func latestDayToCross(row int, col int, cells [][]int) int {
	left, right := 0, row*col
	res := 0
	for left < right-1 {
		mid := left + (right-left)/2
		if latestDayToCrossCheck(row, col, mid, cells) {
			left = mid
			res = mid
		} else {
			right = mid
		}
	}
	return res
}

func latestDayToCrossCheck(m int, n int, t int, cells [][]int) bool {
	grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < t; i++ {
		grid[cells[i][0]][cells[i][1]] = 1
	}

	queue := [][]int{}
	for i := 1; i <= n; i++ {
		if grid[1][i] == 0 {
			queue = append(queue, []int{1, i})
			grid[1][i] = 1
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]
		for _, move := range moves {
			nr, nc := r+move[0], c+move[1]
			if nr > 0 && nc > 0 && nr <= m && nc <= n && grid[nr][nc] == 0 {
				grid[nr][nc] = 1
				if nr == m {
					return true
				}
				queue = append(queue, []int{nr, nc})
			}
		}
	}
	return false
}
