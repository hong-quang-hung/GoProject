package medium

import "fmt"

// Reference: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
func init() {
	Solutions[236] = func() {
		fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1")
		root := S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]")
		p, q := root.Left, root.Right
		fmt.Println("Output:", STreeNode(lowestCommonAncestor(root, p, q)))

		fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4")
		root = S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]")
		p, q = root.Left, root.Left.Right.Right
		fmt.Println("Output:", STreeNode(lowestCommonAncestor(root, p, q)))

		fmt.Println("Input: root = [1,2], p = 1, q = 2")
		root = S2TreeNode("[1,2]")
		p, q = root, root.Left
		fmt.Println("Output:", STreeNode(lowestCommonAncestor(root, p, q)))
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
