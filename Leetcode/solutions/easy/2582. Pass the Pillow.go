package easy

import "fmt"

// Reference: https://leetcode.com/problems/pass-the-pillow/
func init() {
	Solutions[2582] = func() {
		fmt.Println(`Input: n = 4, time = 5`)
		fmt.Println(`Output:`, passThePillow(4, 5))
		fmt.Println(`Input: n = 3, time = 2`)
		fmt.Println(`Output:`, passThePillow(3, 2))
	}
}

func passThePillow(n int, time int) int {
	div := time / (n - 1)
	mod := time % (n - 1)
	if div%2 == 0 {
		return 1 + mod
	}
	return n - mod
}
