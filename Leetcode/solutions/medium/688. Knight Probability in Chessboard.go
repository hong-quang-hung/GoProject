package medium

import "fmt"

// Reference: https://leetcode.com/problems/knight-probability-in-chessboard/
func init() {
	Solutions[688] = func() {
		fmt.Println("Input: n = 3, k = 2, row = 0, column = 0")
		fmt.Println("Output:", knightProbability(3, 2, 0, 0))
		fmt.Println("Input: n = 8, k = 30, row = 6, column = 4")
		fmt.Println("Output:", knightProbability(8, 30, 6, 4))
		fmt.Println("Input: n = 3, k = 3, row = 0, column = 0")
		fmt.Println("Output:", knightProbability(3, 3, 0, 0))
	}
}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, k)
	for i := 0; i < k; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, n)
		}
	}

	return knightProbabilityCount(n, k-1, row, column, dp)
}

func knightProbabilityCount(n int, k int, row int, column int, dp [][][]float64) float64 {
	if row < 0 || row >= n || column < 0 || column >= n {
		return 0
	}

	if k < 0 {
		return 1
	}

	if dp[k][row][column] != 0 {
		return dp[k][row][column]
	}

	res := float64(0)
	for _, dir := range KnightDirection {
		res += knightProbabilityCount(n, k-1, row+dir[0], column+dir[1], dp) / 8
	}

	dp[k][row][column] = res
	return res
}
