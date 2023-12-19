package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/reducing-dishes/
func init() {
	Solutions[1402] = func() {
		fmt.Println(`Input: satisfaction = [-1,-8,0,5,-9]`)
		fmt.Println(`Output:`, maxSatisfaction([]int{-1, -8, 0, 5, -9}))
		fmt.Println(`Input: satisfaction = [-1,-4,-5]`)
		fmt.Println(`Output:`, maxSatisfaction([]int{-1, -4, -5}))
	}
}

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool { return satisfaction[i] < satisfaction[j] })
	n, max, total := len(satisfaction), 0, 0

	for i := n - 1; i >= 0 && satisfaction[i]+total > 0; i-- {
		total += satisfaction[i]
		max += total
	}
	return max
}
