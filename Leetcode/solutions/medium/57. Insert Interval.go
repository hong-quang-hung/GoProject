package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/insert-interval/
func init() {
	Solutions[57] = func() {
		fmt.Println("Input: intervals = [[1,1],[6,9]], newInterval = [2,5]")
		fmt.Println("Output:", insert(S2SoSliceInt("[[1,1],[6,9]]"), []int{2, 5}))
		fmt.Println("Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]")
		fmt.Println("Output:", insert(S2SoSliceInt("[[1,2],[3,5],[6,7],[8,10],[12,16]]"), []int{4, 8}))
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	i := sort.Search(n, func(i int) bool { return intervals[i][0] > newInterval[0] })
	j := sort.Search(n, func(j int) bool { return intervals[j][1] > newInterval[1] })
	if i >= 1 && newInterval[0] <= intervals[i-1][1] {
		newInterval[0] = intervals[i-1][0]
		i--
	}
	if j < n && newInterval[1] >= intervals[j][0] {
		newInterval[1] = intervals[j][1]
		j++
	}
	return append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
}
