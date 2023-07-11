package medium

import "fmt"

// Reference: https://leetcode.com/problems/make-costs-of-paths-equal-in-a-binary-tree/
func init() {
	Solutions[2673] = func() {
		fmt.Println("Input : n = 7, cost = [1,5,2,2,3,3,1]")
		fmt.Println("Output:", minIncrements(7, []int{1, 5, 2, 2, 3, 3, 1}))
		fmt.Println("Input: n = 3, cost = [5,3,3]")
		fmt.Println("Output:", minIncrements(3, []int{5, 3, 3}))
	}
}

func minIncrements(n int, cost []int) int {
	res := 0
	for i := n/2 - 1; i >= 0; i-- {
		left, right := i*2+1, i*2+2
		res += abs(cost[left] - cost[right])
		cost[i] += max(cost[left], cost[right])
	}
	return res
}
