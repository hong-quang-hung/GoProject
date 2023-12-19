package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
func init() {
	Solutions[1489] = func() {
		fmt.Println(`Input: n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]`)
		fmt.Println(`Output:`, findCriticalAndPseudoCriticalEdges(5, S2SoSliceInt(`[[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]`)))
		fmt.Println(`Input: n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]`)
		fmt.Println(`Output:`, findCriticalAndPseudoCriticalEdges(4, S2SoSliceInt(`[[0,1,1],[1,2,1],[2,3,1],[0,3,1]]`)))
	}
}

type DSU struct {
	parent  []int
	size    []int
	maxSize int
}

func newDSU(n int) DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return DSU{parent, size, 1}
}

func (u *DSU) find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *DSU) set(x int, y int) bool {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return false
	}

	if u.size[rootX] < u.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	u.parent[rootY] = rootX
	u.size[rootX] += u.size[rootY]
	u.maxSize = max(u.maxSize, u.size[rootX])
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	ufStd, stdWeight := newDSU(n), 0
	for _, edge := range edges {
		if ufStd.set(edge[0], edge[1]) {
			stdWeight += edge[2]
		}
	}

	res := [][]int{{}, {}}
	for i := 0; i < m; i++ {
		ufIgnore, ignoreWeight := newDSU(n), 0
		for j := 0; j < m; j++ {
			if i != j && ufIgnore.set(edges[j][0], edges[j][1]) {
				ignoreWeight += edges[j][2]
			}
		}

		if ufIgnore.maxSize < n || ignoreWeight > stdWeight {
			res[0] = append(res[0], edges[i][3])
		} else {
			ufForce := newDSU(n)
			ufForce.set(edges[i][0], edges[i][1])
			forceWeight := edges[i][2]
			for j := 0; j < m; j++ {
				if i != j && ufForce.set(edges[j][0], edges[j][1]) {
					forceWeight += edges[j][2]
				}
			}

			if forceWeight == stdWeight {
				res[1] = append(res[1], edges[i][3])
			}
		}
	}
	return res
}
