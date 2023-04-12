package hard

import "fmt"

// Reference: https://leetcode.com/problems/longest-cycle-in-a-graph/
func Leetcode_Longest_Cycle() {
	fmt.Println("Input: edges = [3,3,4,2,3]")
	fmt.Println("Output:", longestCycle([]int{3, 3, 4, 2, 3}))
	fmt.Println("Input: edges = [2,-1,3,1]")
	fmt.Println("Output:", longestCycle([]int{2, -1, 3, 1}))
}

func longestCycle(edges []int) int {
	n := len(edges)
	visit := make([]bool, n)
	indegree := make([]int, n)

	for _, edge := range edges {
		if edge != -1 {
			indegree[edge]++
		}
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		var node int = q[0]
		q = q[1:]
		visit[node] = true
		var neighbor int = edges[node]
		if neighbor != -1 {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	var answer int = -1
	for i := 0; i < n; i++ {
		if !visit[i] {
			var neighbor int = edges[i]
			var count int = 1
			visit[i] = true
			for neighbor != i {
				visit[neighbor] = true
				count++
				neighbor = edges[neighbor]
			}
			if answer < count {
				answer = count
			}
		}
	}
	return answer
}
