package hard

import "fmt"

// Reference: https://leetcode.com/problems/shortest-path-visiting-all-nodes/
func init() {
	Solutions[847] = func() {
		fmt.Println("Input: graph = [[1,2,3],[0],[0],[0]]")
		fmt.Println("Output:", shortestPathLength(S2SoSliceInt("[[1,2,3],[0],[0],[0]]")))
		fmt.Println("Input: graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]")
		fmt.Println("Output:", shortestPathLength(S2SoSliceInt("[[1],[0,2,4],[1,3,4],[2],[1,2]]")))
	}
}

func shortestPathLength(graph [][]int) int {
	return 0
}
