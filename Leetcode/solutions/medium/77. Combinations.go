package medium

import "fmt"

// Reference: https://leetcode.com/problems/combinations/
func init() {
	Solutions[77] = func() {
		fmt.Println(`Input: n = 4, k = 2`)
		fmt.Println(`Output:`, combine(4, 2))
		fmt.Println(`Input: n = 1, k = 1`)
		fmt.Println(`Output:`, combine(1, 1))
	}
}

func combine(n int, k int) [][]int {
	var res [][]int
	combineDFS(&res, n, k-1, make([]int, k))
	return res
}

func combineDFS(res *[][]int, n int, k int, current []int) {
	if k == -1 {
		dest := make([]int, len(current))
		copy(dest, current)
		*res = append(*res, dest)
	} else {
		for i := n; i > 0; i-- {
			current[k] = i
			combineDFS(res, i-1, k-1, current)
		}
	}
}
