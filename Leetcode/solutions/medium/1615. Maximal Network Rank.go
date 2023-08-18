package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximal-network-rank/
func init() {
	Solutions[1615] = func() {
		fmt.Println("Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]")
		fmt.Println("Output:", maximalNetworkRank(4, S2SoSliceInt("[[0,1],[0,3],[1,2],[1,3]]")))
		fmt.Println("Input: n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]")
		fmt.Println("Output:", maximalNetworkRank(5, S2SoSliceInt("[[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]")))
		fmt.Println("Input: n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]")
		fmt.Println("Output:", maximalNetworkRank(8, S2SoSliceInt("[[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]")))
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	m := make([]int, n)
	maxA, maxB := 0, 0
	for _, r := range roads {
		m[r[0]]++
		m[r[1]]++
		maxA = max(maxA, m[r[0]])
		maxA = max(maxA, m[r[1]])
	}
	return maxA + maxB
}
