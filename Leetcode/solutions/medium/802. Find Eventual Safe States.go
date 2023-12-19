package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-eventual-safe-states/
func init() {
	Solutions[802] = func() {
		fmt.Println(`Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]`)
		fmt.Println(`Output:`, eventualSafeNodes(S2SoSliceInt(`[[1,2],[2,3],[5],[0],[5],[],[]]`)))
		fmt.Println(`Input: graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]`)
		fmt.Println(`Output:`, eventualSafeNodes(S2SoSliceInt(`[[1,2,3,4],[1,2],[3,4],[0,4],[]]`)))
	}
}

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	dp, visited := make([]int, n), make([]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	for i := 0; i < n; i++ {
		if eventualSafeNodesDFS(graph, dp, visited, i) {
			ans = append(ans, i)
		}
	}
	return
}

func eventualSafeNodesDFS(graph [][]int, isSafe []int, visited []bool, i int) bool {
	if isSafe[i] != -1 {
		return isSafe[i] == 1
	}

	if visited[i] {
		isSafe[i] = 0
		return false
	}

	visited[i] = true
	for _, u := range graph[i] {
		if !eventualSafeNodesDFS(graph, isSafe, visited, u) {
			isSafe[i] = 0
			return false
		}
	}

	isSafe[i] = 1
	return true
}
