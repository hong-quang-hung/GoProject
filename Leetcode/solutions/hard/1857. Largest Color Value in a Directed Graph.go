package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
func init() {
	Solutions[1857] = func() {
		fmt.Println("Input: colors = 'abaca', edges = [[0,1],[0,2],[2,3],[3,4]]")
		fmt.Println("Output:", largestPathValue("abaca", S2SoSliceInt("[[0,1],[0,2],[2,3],[3,4]]")))
	}
}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adj := make(map[int][]int)
	indegree := make([]int, n)
	for _, edge := range edges {
		if adj[edge[0]] == nil {
			adj[edge[0]] = []int{}
		}
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	count, visit, inStack, res := make([][]int, n), make([]bool, n), make([]bool, n), 0
	for i := 0; i < n; i++ {
		count[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		res = max(res, largestPathValueDFS(i, colors, adj, count, visit, inStack))
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func largestPathValueDFS(node int, colors string, adj map[int][]int, count [][]int, visit []bool, inStack []bool) int {
	if inStack[node] {
		return math.MaxInt
	}

	if visit[node] {
		return count[node][colors[node]-'a']
	}

	visit[node] = true
	inStack[node] = true

	if adj[node] != nil {
		for _, neighbor := range adj[node] {
			if largestPathValueDFS(neighbor, colors, adj, count, visit, inStack) == math.MaxInt {
				return math.MaxInt
			}
			for i := 0; i < 26; i++ {
				count[node][i] = max(count[node][i], count[neighbor][i])
			}
		}
	}

	count[node][colors[node]-'a']++
	inStack[node] = false
	return count[node][colors[node]-'a']
}
