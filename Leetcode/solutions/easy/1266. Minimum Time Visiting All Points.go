package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-time-visiting-all-points/
func init() {
	Solutions[1266] = func() {
		fmt.Println(`Input: points = [[1,1],[3,4],[-1,0]]`)
		fmt.Println(`Output:`, minTimeToVisitAllPoints(S2SoSliceInt(`[[1,1],[3,4],[-1,0]]`)))
		fmt.Println(`Input: points = [[3,2],[-2,2]]`)
		fmt.Println(`Output:`, minTimeToVisitAllPoints(S2SoSliceInt(`[[3,2],[-2,2]]`)))
	}
}

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		res += chebyshev(points[i], points[i-1])
	}
	return res
}
