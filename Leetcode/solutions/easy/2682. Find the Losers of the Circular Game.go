package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-the-losers-of-the-circular-game/
func init() {
	Solutions[2682] = func() {
		fmt.Println(`Input: n = 5, k = 2`)
		fmt.Println(`Output:`, circularGameLosers(5, 2))
		fmt.Println(`Input: n = 4, k = 4`)
		fmt.Println(`Output:`, circularGameLosers(4, 4))
	}
}

func circularGameLosers(n int, k int) []int {
	win := make(map[int]bool)
	lose := make(map[int]bool)
	for i := 2; i <= n; i++ {
		lose[i] = true
	}

	win[1] = true
	i, j := k+1, 2*k
	for !win[getValue(n, i)] {
		win[getValue(n, i)] = true
		delete(lose, getValue(n, i))
		i += j
		j += k
	}

	res := []int{}
	for key := range lose {
		res = append(res, key)
	}

	sort.Ints(res)
	return res
}

func getValue(n, i int) int {
	i = i % n
	if i == 0 {
		return n
	}
	return i
}
