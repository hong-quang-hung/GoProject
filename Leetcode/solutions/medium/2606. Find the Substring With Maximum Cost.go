package medium

import "fmt"

func init() {
	Solutions[2606] = Leetcode_Maximum_Cost_Substring
}

// Reference: https://leetcode.com/problems/find-the-seubstring-with-maximum-cost/
func Leetcode_Maximum_Cost_Substring() {
	fmt.Println("Input: s = 'adaa', chars = 'd', vals = [-1000]")
	fmt.Println("Output:", maximumCostSubstring("adaa", "d", []int{-1000}))
	fmt.Println("Input: s = 'abc', chars = 'abc', vals = [-1,-1,-1]")
	fmt.Println("Output:", maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	m := make(map[rune]int)
	for i, ch := range chars {
		m[ch] = vals[i]
	}

	maxCost, cost := 0, 0
	for _, ch := range s {
		var val int
		if v, c := m[ch]; c {
			val = v
		} else {
			val = int(ch-'a') + 1
		}
		cost = max(0, cost+val)
		maxCost = max(maxCost, cost)
	}
	return maxCost
}
