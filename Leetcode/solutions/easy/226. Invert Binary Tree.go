package easy

import "fmt"

// Reference: https://leetcode.com/problems/invert-binary-tree/
func Leetcode_Invert_Tree() {
	fmt.Println("Input: root = [4,2,7,1,3,6,9]")
	fmt.Println("Output:", invertTree(S2TreeNode("[4,2,7,1,3,6,9]")))
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
