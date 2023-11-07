package easy

import "fmt"

// https://leetcode.com/problems/same-tree/
func init() {
	Solutions[100] = func() {
		fmt.Println("Input: p = [1,2,3], q = [1,2,3]")
		fmt.Println("Output:", isSameTree(S2TreeNode("[1,2,3]"), S2TreeNode("[1,2,3]")))
		fmt.Println("Input: p = [1,2,1], q = [1,1,2]")
		fmt.Println("Output:", isSameTree(S2TreeNode("[1,2,1]"), S2TreeNode("[1,1,2]")))
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || p != nil && q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
