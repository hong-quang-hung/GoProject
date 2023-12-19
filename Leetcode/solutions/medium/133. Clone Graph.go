package medium

import "fmt"

// Reference: https://leetcode.com/problems/clone-graph/
func init() {
	Solutions[133] = func() {
		fmt.Println(`Input: adjList = [[2,4],[1,3],[2,4],[1,3]]`)
		adjList1 := new(Node)
		adjList2 := new(Node)
		adjList3 := new(Node)
		adjList4 := new(Node)

		adjList1.Val = 1
		adjList1.Neighbors = []*Node{adjList2, adjList4}
		adjList2.Val = 2
		adjList2.Neighbors = []*Node{adjList1, adjList3}
		adjList3.Val = 3
		adjList3.Neighbors = []*Node{adjList2, adjList4}
		adjList4.Val = 4
		adjList3.Neighbors = []*Node{adjList1, adjList3}
		fmt.Println(`Output:`, cloneGraph(adjList1))
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make([]*Node, 101)
	return cloneGraphTraversal(node, visited)
}

func cloneGraphTraversal(node *Node, visited []*Node) *Node {
	if visited[node.Val] != nil {
		return visited[node.Val]
	}

	res := new(Node)
	res.Val = node.Val
	visited[node.Val] = res
	for _, neighbor := range node.Neighbors {
		res.Neighbors = append(res.Neighbors, cloneGraphTraversal(neighbor, visited))
	}
	return res
}
