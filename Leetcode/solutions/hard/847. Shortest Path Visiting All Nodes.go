package hard

import "fmt"

// Reference: https://leetcode.com/problems/shortest-path-visiting-all-nodes/
func init() {
	Solutions[847] = func() {
		fmt.Println(`Input: graph = [[1,2,3],[0],[0],[0]]`)
		fmt.Println(`Output:`, shortestPathLength(S2SoSliceInt(`[[1,2,3],[0],[0],[0]]`)))
		fmt.Println(`Input: graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]`)
		fmt.Println(`Output:`, shortestPathLength(S2SoSliceInt(`[[1],[0,2,4],[1,3,4],[2],[1,2]]`)))
	}
}

func shortestPathLength(graph [][]int) int {
	if len(graph) <= 1 {
		return 0
	}

	des := (1 << len(graph)) - 1
	queue := make([][2]int, 0, len(graph))
	for i := 0; i < len(graph); i++ {
		queue = append(queue, [2]int{i, 1 << i})
	}

	visited := make(map[[2]int]struct{})
	var shortest int
	for len(queue) > 0 {
		shortest++
		var level [][2]int
		for i := 0; i < len(queue); i++ {
			u := queue[i]
			for j := 0; j < len(graph[u[0]]); j++ {
				v := graph[u[0]][j]
				vmask := u[1] | (1 << v)
				if _, ok := visited[[2]int{v, vmask}]; ok {
					continue
				}
				if vmask == des {
					return shortest
				}
				visited[[2]int{v, vmask}] = struct{}{}
				level = append(level, [2]int{v, vmask})
			}
		}
		queue = queue[:0]
		queue = append(queue, level...)
	}
	return 0
}
