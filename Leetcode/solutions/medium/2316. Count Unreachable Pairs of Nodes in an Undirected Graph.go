package medium

import "fmt"

func init() {
	Solutions[2316] = Leetcode_Count_Pairs
}

// Reference: https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
func Leetcode_Count_Pairs() {
	fmt.Println("Input: n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]")
	fmt.Println("Output:", countPairs(7, S2SoSliceInt("[[0,2],[0,5],[2,4],[1,6],[5,4]]")))
}

func countPairs(n int, edges [][]int) int64 {
	union := NewUnionFind(n)
	for _, edge := range edges {
		union.UnionSet(edge[0], edge[1])
	}

	group := make(map[int]int)
	for i := 0; i < n; i++ {
		root := union.Find(i)
		group[root]++
	}

	remainingNodes, numberOfPairs := int64(n), int64(0)
	for _, v := range group {
		numberOfPairs += int64(v) * (remainingNodes - int64(v))
		remainingNodes -= int64(v)
	}
	return numberOfPairs
}
