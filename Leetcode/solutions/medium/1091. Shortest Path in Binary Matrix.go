package medium

import "fmt"

// Reference: https://leetcode.com/problems/shortest-path-in-binary-matrix/
func Leetcode_Shortest_Path_Binary_Matrix() {
	fmt.Println("Input: grid = [[0,1],[1,0]]")
	fmt.Println("Output:", shortestPathBinaryMatrix(S2SoSliceInt("[[0,1],[1,0]]")))
	fmt.Println("Input: grid = [[0,0,0],[1,1,0],[1,1,0]]")
	fmt.Println("Output:", shortestPathBinaryMatrix(S2SoSliceInt("[[0,0,0],[1,1,0],[1,1,0]]")))
	fmt.Println("Input: grid = [[1,0,0],[1,1,0],[1,1,0]]")
	fmt.Println("Output:", shortestPathBinaryMatrix(S2SoSliceInt("[[1,0,0],[1,1,0],[1,1,0]]")))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	depth, directions := 1, [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	points := [][2]int{{0, 0}}
	for len(points) > 0 {
		var tmp [][2]int
		for _, point := range points {
			if point[0] == n-1 && point[1] == n-1 {
				return depth
			}

			for _, d := range directions {
				y, x := point[0]+d[0], point[1]+d[1]
				if y < 0 || y >= n || x < 0 || x >= n || grid[y][x] != 0 {
					continue
				}
				grid[y][x] = 2
				tmp = append(tmp, [2]int{y, x})
			}
			grid[point[0]][point[1]] = 2
		}
		depth++
		points = tmp
	}
	return -1
}
