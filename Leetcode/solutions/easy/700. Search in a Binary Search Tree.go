package easy

import "fmt"

// Reference: https://leetcode.com/problems/search-in-a-binary-search-tree/
func init() {
	Solutions[700] = func() {
		fmt.Println("Input: root = [4,2,7,1,3], val = 2")
		fmt.Println("Output:", STreeNode(searchBST(S2TreeNode("[4,2,7,1,3]"), 2)))
		fmt.Println("Input: root = [4,2,7,1,3], val = 5")
		fmt.Println("Output:", STreeNode(searchBST(S2TreeNode("[4,2,7,1,3]"), 5)))
	}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}
