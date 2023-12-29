package hard

import "fmt"

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
	return 0
}
