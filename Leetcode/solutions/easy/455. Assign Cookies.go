package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/assign-cookies/
func init() {
	Solutions[455] = func() {
		fmt.Println(`Input: g = [1,2,3], s = [1,1]`)
		fmt.Println(`Output:`, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
		fmt.Println(`Input: g = [1,2], s = [1,2,3]`)
		fmt.Println(`Output:`, findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	}
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i := 0
	for i < len(s) && res < len(g) {
		if s[i] >= g[res] {
			res++
		}
		i++
	}
	return res
}
