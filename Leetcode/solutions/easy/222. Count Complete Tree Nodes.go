package easy

import "fmt"

// Reference: https://leetcode.com/problems/count-complete-tree-nodes/
func init() {
	Solutions[222] = func() {
		fmt.Println("Input: root = [1,2,3,4,5,6]")
		fmt.Println("Output:", countNodes(S2TreeNode("[1,2,3,4,5,6]")))
	}
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
