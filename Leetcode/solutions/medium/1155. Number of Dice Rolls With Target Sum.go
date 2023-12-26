package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
func init() {
	Solutions[1155] = func() {
		fmt.Println(`Input: n = 1, k = 6, target = 3`)
		fmt.Println(`Output:`, numRollsToTarget(1, 6, 3))
		fmt.Println(`Input: n = 2, k = 6, target = 7`)
		fmt.Println(`Output:`, numRollsToTarget(2, 6, 7))
		fmt.Println(`Input: n = 30, k = 30, target = 500`)
		fmt.Println(`Output:`, numRollsToTarget(30, 30, 500))
	}
}

func numRollsToTarget(n int, k int, target int) int {
	return 0
}
