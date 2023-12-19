package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/non-overlapping-intervals/
func init() {
	Solutions[435] = func() {
		fmt.Println(`Input: intervals = [[1,2],[2,3],[3,4],[1,3]]`)
		fmt.Println(`Output:`, eraseOverlapIntervals(S2SoSliceInt(`[[1,2],[2,3],[3,4],[1,3]]`)))
		fmt.Println(`Input: intervals = [[1,2],[1,2],[1,2]]`)
		fmt.Println(`Output:`, eraseOverlapIntervals(S2SoSliceInt(`[[1,2],[1,2],[1,2]]`)))
		fmt.Println(`Input: intervals = [[1,2],[2,3]]`)
		fmt.Println(`Output:`, eraseOverlapIntervals(S2SoSliceInt(`[[1,2],[2,3]]`)))
	}
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	n := len(intervals)
	res := 0
	i := 0
	for i < n {
		j := i + 1
		for j < n && intervals[i][1] > intervals[j][0] {
			j++
			res++
		}
		i = j
	}
	return res
}
