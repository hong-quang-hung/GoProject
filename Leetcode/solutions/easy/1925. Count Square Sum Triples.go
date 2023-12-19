package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/count-square-sum-triples/
func init() {
	Solutions[1925] = func() {
		fmt.Println(`Input: n = 5`)
		fmt.Println(`Output:`, countTriples(5))
		fmt.Println(`Input: n = 10`)
		fmt.Println(`Output:`, countTriples(10))
		fmt.Println(`Input: n = 41`)
		fmt.Println(`Output:`, countTriples(41))
	}
}

func countTriples(n int) int {
	res, limit := 0, math.Pow(float64(n), 2)
	m := make(map[float64]bool)
	for i := 5; i <= n; i++ {
		m[math.Pow(float64(i), 2)] = true
	}

	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			sum := math.Pow(float64(i), 2) + math.Pow(float64(j), 2)
			if sum > limit {
				break
			}
			if m[sum] {
				if i == j {
					res++
				} else {
					res += 2
				}
			}
		}
	}
	return res
}
