package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/stone-game-iii/
func Leetcode_Stone_Game_III() {
	fmt.Println("Input: values = [1,2,3,7]")
	fmt.Println("Output:", stoneGameIII([]int{1, 2, 3, 7}))
	fmt.Println("Input: values = [1,2,3,-9]")
	fmt.Println("Output:", stoneGameIII([]int{1, 2, 3, -9}))
	fmt.Println("Input: values = [1,2,3,6]")
	fmt.Println("Output:", stoneGameIII([]int{1, 2, 3, 6}))
}

func stoneGameIII(stoneValue []int) string {
	dp := make([]int, len(stoneValue))
	for i := range dp {
		dp[i] = math.MinInt
	}

	switch d0 := stoneGameIIISolved(stoneValue, dp, 0); {
	case d0 > 0:
		return "Alice"
	case d0 < 0:
		return "Bob"
	default:
		return "Tie"
	}
}

func stoneGameIIISolved(stoneValue []int, dp []int, i int) int {
	if i == len(stoneValue) {
		return 0
	}

	if dp[i] != math.MinInt {
		return dp[i]
	}

	s := 0
	for j := 0; j <= 2 && i+j < len(stoneValue); j++ {
		s += stoneValue[i+j]
		if h := s - stoneGameIIISolved(stoneValue, dp, i+j+1); h > dp[i] {
			dp[i] = h
		}
	}
	return dp[i]
}
