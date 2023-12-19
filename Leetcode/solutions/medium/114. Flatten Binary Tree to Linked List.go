package medium

import "fmt"

// Reference: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
func init() {
	Solutions[114] = func() {
		var root *TreeNode

		fmt.Println(`Input: root = [1,2,5,3,4,null,6]`)
		root = S2TreeNode(`[1,2,5,3,4,null,6]`)
		flatten(root)
		fmt.Println(`Output:`, STreeNode(root))

		fmt.Println(`Input: root = [0]`)
		root = S2TreeNode(`[0]`)
		flatten(root)
		fmt.Println(`Output:`, STreeNode(root))
	}
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Right)
	flatten(root.Left)

	temp := root.Right
	root.Right = root.Left
	root.Left = nil

	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	cur.Right = temp
}
