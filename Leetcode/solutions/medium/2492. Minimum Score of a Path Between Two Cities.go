package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
func init() {
	Solutions[2492] = func() {
		fmt.Println("Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]")
		fmt.Println("Output:", minScore(4, S2SoSliceInt("[[1,2,9],[2,3,6],[2,4,5],[1,4,7]]")))
	}
}

func minScore(n int, roads [][]int) int {
	union := NewUnionFind(n + 1)
	var answer int = math.MaxInt16

	for _, road := range roads {
		union.UnionSet(road[0], road[1])
	}

	for _, road := range roads {
		if union.Find(1) == union.Find(road[0]) {
			if answer > road[2] {
				answer = road[2]
			}
		}
	}
	return answer
}
