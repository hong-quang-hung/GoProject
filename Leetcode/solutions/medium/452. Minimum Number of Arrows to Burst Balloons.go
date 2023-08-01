package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
func init() {
	Solutions[452] = func() {
		fmt.Println("Input: points = [[10,16],[2,8],[1,6],[7,12]]")
		fmt.Println("Output:", findMinArrowShots(S2SoSliceInt("[[10,16],[2,8],[1,6],[7,12]]")))
		fmt.Println("Input: points = [[1,2],[3,4],[5,6],[7,8]]")
		fmt.Println("Output:", findMinArrowShots(S2SoSliceInt("[[1,2],[3,4],[5,6],[7,8]]")))
		fmt.Println("Input: points = [[1,2],[2,3],[3,4],[4,5]]")
		fmt.Println("Output:", findMinArrowShots(S2SoSliceInt("[[1,2],[2,3],[3,4],[4,5]]")))
	}
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	n := len(points)
	res := 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && points[i][1] >= points[j][0] {
			j++
		}

		res++
		i = j
	}
	return res
}
