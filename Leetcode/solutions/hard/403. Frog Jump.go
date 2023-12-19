package hard

import "fmt"

// Reference: https://leetcode.com/problems/frog-jump/
func init() {
	Solutions[403] = func() {
		fmt.Println(`Input: stones = [0,1,3,5,6,8,12,17]`)
		fmt.Println(`Output:`, canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
		fmt.Println(`Input: stones = [0,1,2,3,4,8,9,11]`)
		fmt.Println(`Output:`, canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	}
}

func canCross(stones []int) bool {
	n := len(stones)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[stones[i]] = i
	}

	dp := [2001][2001]bool{}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for prevJump := 0; prevJump <= n; prevJump++ {
			if dp[i][prevJump] {
				if v, ok := m[stones[i]+prevJump]; ok {
					dp[v][prevJump] = true
				}

				if v, ok := m[stones[i]+prevJump+1]; ok {
					dp[v][prevJump+1] = true
				}

				if v, ok := m[stones[i]+prevJump-1]; ok {
					dp[v][prevJump-1] = true
				}
			}
		}
	}

	for i := 0; i <= n; i++ {
		if dp[n-1][i] {
			return true
		}
	}
	return false
}
