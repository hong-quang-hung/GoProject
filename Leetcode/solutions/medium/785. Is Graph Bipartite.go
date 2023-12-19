package medium

import "fmt"

// Reference: https://leetcode.com/problems/is-graph-bipartite/
func init() {
	Solutions[785] = func() {
		fmt.Println(`Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]`)
		fmt.Println(`Output:`, isBipartite(S2SoSliceInt(`[[1,2,3],[0,2],[0,1,3],[0,2]]`)))
	}
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited := make([]int, n)
	res := true
	for i := 0; i < n && res; i++ {
		if visited[i] == 0 {
			res = isBipartiteDFS(i, 1, graph, visited)
		}
	}
	return res
}

func isBipartiteDFS(curr int, currGrp int, graph [][]int, visited []int) bool {
	if visited[curr] == -currGrp {
		return false
	}

	if visited[curr] == 0 {
		visited[curr] = currGrp
		for _, node := range graph[curr] {
			if !isBipartiteDFS(node, -currGrp, graph, visited) {
				return false
			}
		}
		return true
	}
	return true
}
