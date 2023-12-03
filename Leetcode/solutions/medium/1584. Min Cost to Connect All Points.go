package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/min-cost-to-connect-all-points/
func init() {
	Solutions[1584] = func() {
		fmt.Println("Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]")
		fmt.Println("Output:", minCostConnectPoints(S2SoSliceInt("[[0,0],[2,2],[3,10],[5,2],[7,0]]")))
		fmt.Println("Input: points = [[3,12],[-2,5],[-4,1]]")
		fmt.Println("Output:", minCostConnectPoints(S2SoSliceInt("[[3,12],[-2,5],[-4,1]]")))
	}
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	m := [][]int{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			m = append(m, []int{i, j, manhattan(points[i], points[j])})
		}
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i][2] < m[j][2]
	})

	p, i, union, distance := 0, 0, NewUnionFind(n), make([]int, n)
	for p < n-1 {
		nextPoint := m[i]
		i++
		x := union.Find(nextPoint[0])
		y := union.Find(nextPoint[1])
		if x != y {
			distance[p] = nextPoint[2]
			union.UnionSet(x, y)
			p++
		}
	}

	minimumCost := 0
	for j := 0; j < p; j++ {
		minimumCost += distance[j]
	}
	return minimumCost
}
