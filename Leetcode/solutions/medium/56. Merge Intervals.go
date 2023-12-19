package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/merge-intervals/
func init() {
	Solutions[56] = func() {
		fmt.Println(`Input: intervals = [[1,3],[2,6],[8,10],[15,18]]`)
		fmt.Println(`Output:`, merge(S2SoSliceInt(`[[1,3],[2,6],[8,10],[15,18]]`)))
		fmt.Println(`Input: intervals = [[1,4],[4,5]]`)
		fmt.Println(`Output:`, merge(S2SoSliceInt(`[[1,4],[4,5]]`)))
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	n := len(intervals)
	res := [][]int{}
	i := 0
	for i < n {
		j := i + 1
		end := intervals[i][1]
		for j < n && end >= intervals[j][0] {
			if end < intervals[j][1] {
				end = intervals[j][1]
			}
			j++
		}

		res = append(res, []int{intervals[i][0], end})
		i = j
	}
	return res
}
