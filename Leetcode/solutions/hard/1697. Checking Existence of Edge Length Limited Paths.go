package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/
func Leetcode_Distance_Limited_Paths_Exist() {
	fmt.Println("Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]")
	fmt.Println("Output:", distanceLimitedPathsExist(3, S2SoSliceInt("[[0,1,2],[1,2,4],[2,0,8],[1,0,16]]"), S2SoSliceInt("[[0,1,2],[0,2,5]]")))
	fmt.Println("Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]")
	fmt.Println("Output:", distanceLimitedPathsExist(5, S2SoSliceInt("[[0,1,10],[1,2,5],[2,3,9],[3,4,13]]"), S2SoSliceInt("[[0,4,14],[1,4,13]]")))
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	union := NewUnionFind(n)
	queriesCount := len(queries)
	queriesIndex, res := make([][4]int, queriesCount), make([]bool, queriesCount)
	for i := 0; i < queriesCount; i++ {
		queriesIndex[i][0] = queries[i][0]
		queriesIndex[i][1] = queries[i][1]
		queriesIndex[i][2] = queries[i][2]
		queriesIndex[i][3] = i
	}

	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })
	sort.Slice(queriesIndex, func(i, j int) bool { return queriesIndex[i][2] < queriesIndex[j][2] })

	edgesIndex := 0
	for queryIndex := 0; queryIndex < queriesCount; queryIndex += 1 {
		p, q, limit, originalIndex := queriesIndex[queryIndex][0], queriesIndex[queryIndex][1], queriesIndex[queryIndex][2], queriesIndex[queryIndex][3]
		for edgesIndex < len(edgeList) && edgeList[edgesIndex][2] < limit {
			node1 := edgeList[edgesIndex][0]
			node2 := edgeList[edgesIndex][1]
			union.UnionSet(node1, node2)
			edgesIndex += 1
		}
		res[originalIndex] = union.Find(p) == union.Find(q)
	}
	return res
}
