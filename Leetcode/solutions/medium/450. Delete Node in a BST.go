package medium

import "fmt"

// Reference: https://leetcode.com/problems/search-in-a-binary-search-tree/
func init() {
	Solutions[450] = func() {
		fmt.Println(`Input: root = [5,3,6,2,4,null,7], key = 3`)
		fmt.Println(`Output:`, STreeNode(deleteNode(S2TreeNode(`[5,3,6,2,4,null,7]`), 3)))
		fmt.Println(`Input: root = [5,3,6,2,4,null,7], key = 0`)
		fmt.Println(`Output:`, STreeNode(deleteNode(S2TreeNode(`[5,3,6,2,4,null,7]`), 0)))
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return deleteNodeSolved(root.Left, root.Right)
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func deleteNodeSolved(left, right *TreeNode) *TreeNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	right.Left = deleteNodeSolved(left, right.Left)
	return right
}
