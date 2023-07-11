package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-operations-to-make-network-connected/
func init() {
	Solutions[1319] = func() {
		fmt.Println("Input: n = 4, connections = [[0,1],[0,2],[1,2]]")
		fmt.Println("Output:", makeConnected(4, S2SoSliceInt("[[0,1],[0,2],[1,2]]")))
	}
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	union := NewUnionFind(n)
	var numberOfConnectedComponents int = n
	for _, connection := range connections {
		if union.Find(connection[0]) != union.Find(connection[1]) {
			numberOfConnectedComponents--
			union.UnionSet(connection[0], connection[1])
		}
	}
	return numberOfConnectedComponents - 1
}
