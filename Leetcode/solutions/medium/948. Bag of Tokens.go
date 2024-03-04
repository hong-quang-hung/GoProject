package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/bag-of-tokens/
func init() {
	Solutions[948] = func() {
		fmt.Println(`Input: tokens = [100], power = 50`)
		fmt.Println(`Output:`, bagOfTokensScore([]int{100}, 50))
		fmt.Println(`Input: tokens = [200,100], power = 150`)
		fmt.Println(`Output:`, bagOfTokensScore([]int{200, 100}, 150))
		fmt.Println(`Input: tokens = [100,200,300,400], power = 200`)
		fmt.Println(`Output:`, bagOfTokensScore([]int{100, 200, 300, 400}, 200))
	}
}

func bagOfTokensScore(tokens []int, power int) int {
	n := len(tokens)
	sort.Ints(tokens)

	score := 0
	res := 0
	i, j := 0, n-1
	for i <= j {
		if power >= tokens[i] {
			power -= tokens[i]
			score++
		}

		res = max(res, score)
		if i+1 <= j && power < tokens[i+1] && score >= 1 {
			power += tokens[j]
			score--
			j--
		}
		i++
	}
	return res
}
