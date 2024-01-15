package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/find-players-with-zero-or-one-losses/
func init() {
	Solutions[2225] = func() {
		fmt.Println(`Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]`)
		fmt.Println(`Output:`, findWinners(S2SoSliceInt("[[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]")))
		fmt.Println(`Input: matches = [[2,3],[1,3],[5,4],[6,4]]`)
		fmt.Println(`Output:`, findWinners(S2SoSliceInt("[[2,3],[1,3],[5,4],[6,4]]")))
	}
}

func findWinners(matches [][]int) [][]int {
	m := make([]int, 100001)
	for i := 0; i < len(matches); i++ {
		if m[matches[i][0]] == 0 {
			m[matches[i][0]] = 1
		}

		if m[matches[i][1]] == 0 {
			m[matches[i][1]] = 1
		}

		m[matches[i][1]]++
	}

	res := make([][]int, 2)
	for i := 0; i < len(m); i++ {
		if m[i] == 1 {
			res[0] = append(res[0], i)
		} else if m[i] == 2 {
			res[1] = append(res[1], i)
		}
	}
	return res
}
