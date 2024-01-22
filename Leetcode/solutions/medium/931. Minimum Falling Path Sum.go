package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-falling-path-sum/
func init() {
	Solutions[931] = func() {
		fmt.Println(`Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]`)
		fmt.Println(`Output:`, minFallingPathSum(S2SoSliceInt("[[2,1,3],[6,5,4],[7,8,9]]")))
		fmt.Println(`Input: matrix = [[-19,57],[-40,-5]]`)
		fmt.Println(`Output:`, minFallingPathSum(S2SoSliceInt("[[-19,57],[-40,-5]]")))
	}
}

func minFallingPathSum(matrix [][]int) int {
	return 0
}
