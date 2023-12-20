package hard

import "fmt"

// Reference: https://leetcode.com/problems/max-points-on-a-line/
func init() {
	Solutions[149] = func() {
		fmt.Println(`Input: points = [[1,1],[2,2],[3,3]]`)
		fmt.Println(`Output:`, maxPoints(S2SoSliceInt(`[[1,1],[2,2],[3,3]]`)))
		fmt.Println(`Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]`)
		fmt.Println(`Output:`, maxPoints(S2SoSliceInt(`[[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]`)))
	}
}

func maxPoints(points [][]int) int {
	n := len(points)
	res := 1
	for i := 0; i < n; i++ {
		p1 := points[i]
		m := make(map[float64]int)
		vertical := 0
		for j := i + 1; j < n; j++ {
			p2 := points[j]
			if p1[0] == p2[0] {
				vertical += 1
				continue
			}
			slope := float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
			m[slope]++
			res = max(m[slope]+1, res)
		}
		res = max(vertical+1, res)
	}
	return res
}
