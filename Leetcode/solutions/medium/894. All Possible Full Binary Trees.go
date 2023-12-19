package medium

import "fmt"

// Reference: https://leetcode.com/problems/all-possible-full-binary-trees/
func init() {
	Solutions[894] = func() {
		fmt.Println(`Input: n = 7`)
		fmt.Println(`Output:`, allPossibleFBT(7))

		fmt.Println(`Input: n = 3`)
		fmt.Println(`Output:`, allPossibleFBT(3))
	}
}

func allPossibleFBT(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	dp[1] = []*TreeNode{{Val: 0}}
	return allPossibleFBTDFS(n, &dp)
}

func allPossibleFBTDFS(n int, dp *[][]*TreeNode) []*TreeNode {
	if n%2 == 0 {
		return nil
	} else if (*dp)[n] != nil {
		return (*dp)[n]
	}

	res := []*TreeNode{}
	for i := 1; i < n; i += 2 {
		left := allPossibleFBTDFS(i, dp)
		right := allPossibleFBTDFS(n-i-1, dp)
		for _, l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{Left: l, Right: r})
			}
		}
	}

	(*dp)[n] = res
	return res
}
