package easy

import "fmt"

// Reference: https://leetcode.com/problems/count-of-matches-in-tournament/
func init() {
	Solutions[1688] = func() {
		fmt.Println(`Input: n = 7`)
		fmt.Println(`Output:`, numberOfMatches(7))
		fmt.Println(`Input: n = 14`)
		fmt.Println(`Output:`, numberOfMatches(14))
	}
}

func numberOfMatches(n int) int {
	return n - 1
}
