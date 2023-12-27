package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
func init() {
	Solutions[1578] = func() {
		fmt.Println(`Input: colors = "abaac", neededTime = [1,2,3,4,5]`)
		fmt.Println(`Output:`, minCost("abaac", []int{1, 2, 3, 4, 5}))
		fmt.Println(`Input: colors = "abc", neededTime = [1,2,3]`)
		fmt.Println(`Output:`, minCost("abc", []int{1, 2, 3}))
		fmt.Println(`Input: colors = "aabaa", neededTime = [1,2,3,4,1]`)
		fmt.Println(`Output:`, minCost("aabaa", []int{1, 2, 3, 4, 1}))
	}
}

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	res := 0
	for i := 0; i < n; {
		sum, maxTime := neededTime[i], neededTime[i]
		j := i + 1
		for j < n && colors[j] == colors[i] {
			sum += neededTime[j]
			maxTime = max(maxTime, neededTime[j])
			j++
		}
		i = j
		res += sum - maxTime
	}
	return res
}
