package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/clone-graph/
func Leetcode_Clone_Graph() {
	fmt.Println("Input: adjList = [[2,4],[1,3],[2,4],[1,3]]")
	adjList1 := new(types.Node)
	adjList2 := new(types.Node)
	adjList3 := new(types.Node)
	adjList4 := new(types.Node)

	adjList1.Val = 1
	adjList1.Neighbors = []*types.Node{adjList2, adjList4}
	adjList2.Val = 2
	adjList2.Neighbors = []*types.Node{adjList1, adjList3}
	adjList3.Val = 3
	adjList3.Neighbors = []*types.Node{adjList2, adjList4}
	adjList4.Val = 4
	adjList3.Neighbors = []*types.Node{adjList1, adjList3}
	fmt.Println("Output:", cloneGraph(adjList1))
}

func cloneGraph(node *types.Node) *types.Node {
	if node == nil {
		return nil
	}

	visited := make([]*types.Node, 101)
	return cloneGraphTraversal(node, visited)
}

func cloneGraphTraversal(node *types.Node, visited []*types.Node) *types.Node {
	if visited[node.Val] != nil {
		return visited[node.Val]
	}

	res := new(types.Node)
	res.Val = node.Val
	visited[node.Val] = res
	for _, neighbor := range node.Neighbors {
		res.Neighbors = append(res.Neighbors, cloneGraphTraversal(neighbor, visited))
	}
	return res
}
