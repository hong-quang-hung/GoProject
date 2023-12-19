package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximal-network-rank/
func init() {
	Solutions[1615] = func() {
		fmt.Println(`Input: n = 4, roads = [[0,1],[0,3],[1,2],[1,3]]`)
		fmt.Println(`Output:`, maximalNetworkRank(4, S2SoSliceInt(`[[0,1],[0,3],[1,2],[1,3]]`)))
		fmt.Println(`Input: n = 5, roads = [[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]`)
		fmt.Println(`Output:`, maximalNetworkRank(5, S2SoSliceInt(`[[0,1],[0,3],[1,2],[1,3],[2,3],[2,4]]`)))
		fmt.Println(`Input: n = 8, roads = [[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]`)
		fmt.Println(`Output:`, maximalNetworkRank(8, S2SoSliceInt(`[[0,1],[1,2],[2,3],[2,4],[5,6],[5,7]]`)))
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	m := make([][]bool, n)
	for i := 0; i < n; i++ {
		m[i] = make([]bool, n)
	}

	degree := make([]int, n)
	for _, r := range roads {
		m[r[0]][r[1]] = true
		m[r[1]][r[0]] = true

		degree[r[0]]++
		degree[r[1]]++
	}

	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[i][j] {
				res = max(res, degree[i]+degree[j]-1)
			} else {
				res = max(res, degree[i]+degree[j])
			}
		}
	}
	return res
}
