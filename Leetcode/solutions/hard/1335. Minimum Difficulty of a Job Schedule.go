package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/
func init() {
	Solutions[1335] = func() {
		fmt.Println(`Input: jobDifficulty = [6,5,4,3,2,1], d = 2`)
		fmt.Println(`Output:`, minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
		fmt.Println(`Input: jobDifficulty = [9,9,9], d = 4`)
		fmt.Println(`Output:`, minDifficulty([]int{9, 9, 9}, 4))
		fmt.Println(`Input: jobDifficulty = [1,1,1], d = 3`)
		fmt.Println(`Output:`, minDifficulty([]int{1, 1, 1}, 3))
	}
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, d+1)
		for j := 0; j <= d; j++ {
			dp[i][j] = -1
		}
	}

	maxJob := make([]int, n)
	maxJob[n-1] = jobDifficulty[n-1]
	for i := n - 2; i >= 0; i-- {
		maxJob[i] = max(maxJob[i+1], jobDifficulty[i])
	}

	var f func(i, j int) int
	f = func(i, j int) int {
		if j == 1 {
			return maxJob[i]
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		res := math.MaxInt
		maxDifficulty := 0
		for k := i; k < n+1-j; k++ {
			maxDifficulty = max(maxDifficulty, jobDifficulty[k])
			res = min(res, maxDifficulty+f(k+1, j-1))
		}

		dp[i][j] = res
		return res
	}
	return f(0, d)
}
