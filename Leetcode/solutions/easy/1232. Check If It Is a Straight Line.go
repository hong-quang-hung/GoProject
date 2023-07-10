package easy

import "fmt"

func init() {
	Solutions[1232] = Leetcode_Check_Straight_Line
}

// Reference: https://leetcode.com/problems/check-if-it-is-a-straight-line/
func Leetcode_Check_Straight_Line() {
	fmt.Println("Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]")
	fmt.Println("Output:", checkStraightLine(S2SoSliceInt("[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]")))
	fmt.Println("Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]")
	fmt.Println("Output:", checkStraightLine(S2SoSliceInt("[[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]")))
	fmt.Println("Input: coordinates = [[0,0],[0,1],[0,-1]]")
	fmt.Println("Output:", checkStraightLine(S2SoSliceInt("[[0,0],[0,1],[0,-1]]")))
}

func checkStraightLine(coordinates [][]int) bool {
	v := []int{coordinates[0][0] - coordinates[1][0], coordinates[0][1] - coordinates[1][1]}
	fx := func(x int) float64 { return float64(x-coordinates[0][0]) / float64(v[0]) }
	fy := func(y int) float64 { return float64(y-coordinates[0][1]) / float64(v[1]) }
	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (v[0] != 0 && v[1] != 0 && fx(x) != fy(y)) || (v[0] == 0 && x != coordinates[0][0]) || (v[1] == 0 && y != coordinates[0][1]) {
			return false
		}
	}
	return true
}
