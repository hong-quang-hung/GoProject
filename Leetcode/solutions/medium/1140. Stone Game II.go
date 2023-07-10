package medium

import "fmt"

func init() {
	Solutions[1140] = Leetcode_Stone_Game_II
}

// Reference: https://leetcode.com/problems/stone-game-ii/
func Leetcode_Stone_Game_II() {
	fmt.Println("Input: piles = [2,7,9,4,4]")
	fmt.Println("Output:", stoneGameII([]int{2, 7, 9, 4, 4}))
	fmt.Println("Input: piles = [1,2,3,4,5,100]")
	fmt.Println("Output:", stoneGameII([]int{1, 2, 3, 4, 5, 100}))
}

func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][][]int, 2)
	for p := 0; p < 2; p++ {
		dp[p] = make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[p][i] = make([]int, n+1)
			for m := 0; m <= n; m++ {
				dp[p][i][m] = -1
			}
		}
	}
	return stoneGameIISolved(piles, dp, 0, 0, 1)
}

func stoneGameIISolved(piles []int, dp [][][]int, p, i, m int) int {
	if i == len(piles) {
		return 0
	}

	if dp[p][i][m] != -1 {
		return dp[p][i][m]
	}

	var res int
	if p == 1 {
		res = 1000000
	} else {
		res = 0
	}

	s := 0
	for x := 1; x <= min(2*m, len(piles)-i); x++ {
		s += piles[i+x-1]
		if p == 0 {
			res = max(res, s+stoneGameIISolved(piles, dp, 1, i+x, max(m, x)))
		} else {
			res = min(res, stoneGameIISolved(piles, dp, 0, i+x, max(m, x)))
		}
	}

	dp[p][i][m] = res
	return dp[p][i][m]
}
