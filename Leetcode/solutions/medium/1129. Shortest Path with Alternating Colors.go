package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/shortest-path-with-alternating-colors/
func init() {
	Solutions[1129] = func() {
		fmt.Println(`Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]`)
		fmt.Println(`Output:`, shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
	}
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	graph := make(map[int][][]int)
	for _, e := range redEdges {
		if graph[e[0]] == nil {
			graph[e[0]] = [][]int{}
		}
		graph[e[0]] = append(graph[e[0]], []int{0, e[1]})
	}

	for _, e := range blueEdges {
		if graph[e[0]] == nil {
			graph[e[0]] = [][]int{}
		}
		graph[e[0]] = append(graph[e[0]], []int{1, e[1]})
	}

	shortestPath := make([][]int, 2)
	shortestPath[0] = make([]int, n)
	shortestPath[1] = make([]int, n)
	for i := 0; i < n; i++ {
		shortestPath[0][i] = math.MaxInt
		shortestPath[1][i] = math.MaxInt
	}
	shortestPath[0][0] = 0
	shortestPath[1][0] = 0

	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[cur] {
			color, node := neighbor[0], neighbor[1]
			tillThisNode := shortestPath[1-color][cur]
			if tillThisNode != math.MaxInt {
				if shortestPath[color][node] > tillThisNode+1 {
					shortestPath[color][node] = tillThisNode + 1
					queue = append(queue, node)
				}
			}
		}
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = min(shortestPath[0][i], shortestPath[1][i])
		if res[i] == math.MaxInt {
			res[i] = -1
		}
	}
	return res
}
