package medium

import "fmt"

func init() {
	Solutions[1466] = Leetcode_Min_Reorder
}

// Reference: https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
func Leetcode_Min_Reorder() {
	fmt.Println("Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]")
	fmt.Println("Output:", minReorder(6, S2SoSliceInt("[[0,1],[1,3],[2,3],[4,0],[4,5]]")))
}

func minReorder(n int, connections [][]int) int {
	count, adj := 0, make(map[int][][]int)
	for _, connection := range connections {
		if adj[connection[0]] == nil {
			adj[connection[0]] = make([][]int, 0)
		}
		adj[connection[0]] = append(adj[connection[0]], []int{connection[1], 1})

		if adj[connection[1]] == nil {
			adj[connection[1]] = make([][]int, 0)
		}
		adj[connection[1]] = append(adj[connection[1]], []int{connection[0], 0})
	}

	minReorderDFS(adj, 0, n, &count)
	return count
}

func minReorderDFS(adj map[int][][]int, node int, parent int, count *int) {
	if adj[node] == nil {
		return
	}

	for _, neib := range adj[node] {
		child, sign := neib[0], neib[1]
		if child != parent {
			*count += sign
			minReorderDFS(adj, child, node, count)
		}
	}
}
