package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
func init() {
	Solutions[1026] = func() {
		fmt.Println(`Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]`)
		fmt.Println(`Output:`, maxAncestorDiff(S2TreeNode("[8,3,10,1,6,null,14,null,null,4,7,13]")))
		fmt.Println(`Input: root = [1,null,2,null,0,3]`)
		fmt.Println(`Output:`, maxAncestorDiff(S2TreeNode("[1,null,2,null,0,3]")))
	}
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	var f func(node *TreeNode, maxVal, minVal int)
	f = func(node *TreeNode, maxVal, minVal int) {
		if node == nil {
			return
		}

		res = max(res, max(abs(node.Val-maxVal), abs(node.Val-minVal)))
		maxVal, minVal = max(maxVal, node.Val), min(minVal, node.Val)

		f(node.Left, maxVal, minVal)
		f(node.Right, maxVal, minVal)
	}
	f(root, root.Val, root.Val)
	return res
}
