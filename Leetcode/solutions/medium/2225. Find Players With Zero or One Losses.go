package medium

import (
	"fmt"
	"sort"
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
	m := make(map[int]int)
	for _, match := range matches {
		m[match[1]]++
		if _, ok := m[match[0]]; ok {
			continue
		} else {
			m[match[0]] = 0
		}
	}

	res := make([][]int, 2)
	for k, v := range m {
		if v == 0 {
			res[0] = append(res[0], k)
		} else if v == 1 {
			res[1] = append(res[1], k)
		}
	}

	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}
