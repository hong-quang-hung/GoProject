package medium

import "fmt"

// Reference: https://leetcode.com/problems/snakes-and-ladders/
func init() {
	Solutions[221] = func() {
		fmt.Println(`Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`)
		fmt.Println(`Output:`, maximalSquare(S2SoSliceByte(`[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`)))
		fmt.Println(`Input: matrix = [["0","1"],["1","0"]]`)
		fmt.Println(`Output:`, maximalSquare(S2SoSliceByte(`[["0","1"],["1","0"]]`)))
	}
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	res := 0
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i][j+1], min(dp[i+1][j], dp[i+1][j+1])) + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res * res
}
