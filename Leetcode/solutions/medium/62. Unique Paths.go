package medium

import "fmt"

// Reference: https://leetcode.com/problems/unique-paths/
func Leetcode_Unique_Paths() {
	fmt.Println("Input: m = 3, n = 7")
	fmt.Println("Output:", uniquePaths(3, 7))
	fmt.Println("Input: m = 3, n = 2")
	fmt.Println("Output:", uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	return uniquePathsSolved(dp, m, n, 0, 0)
}

func uniquePathsSolved(dp [][]int, m, n, i, j int) int {
	if i >= m || j >= n {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	moves := [][2]int{{1, 0}, {0, 1}}
	count := 0
	for _, move := range moves {
		count += uniquePathsSolved(dp, m, n, i+move[0], j+move[1])
	}

	dp[i][j] = count
	return count
}
