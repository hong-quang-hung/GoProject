package medium

import "fmt"

// Reference: https://leetcode.com/problems/domino-and-tromino-tiling/
func init() {
	Solutions[790] = func() {
		fmt.Println("Input: n = 3")
		fmt.Println("Output:", numTilings(3))
		fmt.Println("Input: n = 1")
		fmt.Println("Output:", numTilings(1))
	}
}

func numTilings(n int) int {
	one, two, three := 1, 2, 5
	if n == 1 {
		return one
	}

	if n == 2 {
		return two
	}

	if n == 3 {
		return three
	}

	res := 0
	for i := 4; i <= n; i++ {
		res = (2*three + one) % mod
		one, two, three = two, three, res
	}
	return res
}
