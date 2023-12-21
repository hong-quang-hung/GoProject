package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
func init() {
	Solutions[1637] = func() {
		fmt.Println(`Input: points = [[8,7],[9,9],[7,4],[9,7]]`)
		fmt.Println(`Output:`, maxWidthOfVerticalArea(S2SoSliceInt(`[[8,7],[9,9],[7,4],[9,7]]`)))
		fmt.Println(`Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]`)
		fmt.Println(`Output:`, maxWidthOfVerticalArea(S2SoSliceInt(`[[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]`)))
	}
}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	res := 0
	for i := 1; i < len(points); i++ {
		res = max(res, points[i][0]-points[i-1][0])
	}
	return res
}
