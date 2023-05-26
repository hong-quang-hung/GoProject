package medium

import "fmt"

// Reference: https://leetcode.com/problems/solving-questions-with-brainpower/
func Leetcode_Most_Points() {
	fmt.Println("Input: questions = [[3,2],[4,3],[4,4],[2,5]]")
	fmt.Println("Output:", mostPoints(S2SoSliceInt("[[3,2],[4,3],[4,4],[2,5]]")))
	fmt.Println("Input: questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]")
	fmt.Println("Output:", mostPoints(S2SoSliceInt("[[1,1],[2,2],[3,3],[4,4],[5,5]]")))
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n)
	dp[n-1] = int64(questions[n-1][0])
	for i := 0; i < n-1; i++ {
		dp[i] = -1
	}
	return mostPointsSolve(questions, dp, 0, n)
}

func mostPointsSolve(questions [][]int, dp []int64, i int, n int) int64 {
	if i >= n {
		return 0
	}

	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = max64(int64(questions[i][0])+mostPointsSolve(questions, dp, i+questions[i][1]+1, n), mostPointsSolve(questions, dp, i+1, n))
	return dp[i]
}
