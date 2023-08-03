package easy

import "fmt"

// Reference: https://leetcode.com/problems/binary-tree-inorder-traversal/
func init() {
	Solutions[94] = func() {
		fmt.Println("Input: root = [1,null,2,3]")
		fmt.Println("Output:", inorderTraversal(S2TreeNode("[1,null,2,3]")))
		fmt.Println("Input: root = [1]")
		fmt.Println("Output:", inorderTraversal(S2TreeNode("[1]")))
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
