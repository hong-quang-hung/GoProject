package medium

import "fmt"

// Reference: https://leetcode.com/problems/binary-tree-right-side-view/
func init() {
	Solutions[199] = func() {
		fmt.Println("Input: root = [1,2,3,null,5,null,4]")
		fmt.Println("Output:", rightSideView(S2TreeNode("[1,2,3,null,5,null,4]")))
		fmt.Println("Input: root = [1,null,3]")
		fmt.Println("Output:", rightSideView(S2TreeNode("[1,null,3]")))
	}
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		res = append(res, q[n-1].Val)
		for _, node := range q {
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
	}
	return res
}
