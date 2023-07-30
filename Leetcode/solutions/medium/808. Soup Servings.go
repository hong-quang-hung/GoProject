package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/soup-servings/
func init() {
	Solutions[808] = func() {
		fmt.Println("Input: n = 50")
		fmt.Println("Output:", soupServings(50))
		fmt.Println("Input: n = 100")
		fmt.Println("Output:", soupServings(100))
	}
}

func soupServings(n int) float64 {
	m := int(math.Ceil(float64(n) / 25.0))
	memo := make(map[int]map[int]float64)
	for i := 1; i <= m; i++ {
		memo[i] = make(map[int]float64)
		if soupServingsSolved(i, i, memo) > 1-1e-5 {
			return 1
		}
	}
	return soupServingsSolved(m, m, memo)
}

func soupServingsSolved(a, b int, memo map[int]map[int]float64) float64 {
	if a <= 0 && b <= 0 {
		return 0.5
	} else if a <= 0 {
		return 1
	} else if b <= 0 {
		return 0
	}

	if val, ok := memo[a][b]; ok {
		return val
	}

	memo[a][b] = 0.25 * (soupServingsSolved(a-4, b, memo) + soupServingsSolved(a-3, b-1, memo) + soupServingsSolved(a-2, b-2, memo) + soupServingsSolved(a-1, b-3, memo))
	return memo[a][b]
}
