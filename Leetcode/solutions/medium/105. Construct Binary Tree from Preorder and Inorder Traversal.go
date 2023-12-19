package medium

import "fmt"

// Reference: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func init() {
	Solutions[105] = func() {
		fmt.Println(`Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]`)
		fmt.Println(`Output:`, STreeNode(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
		fmt.Println(`Input: preorder = [-1], inorder = [-1]`)
		fmt.Println(`Output:`, STreeNode(buildTree([]int{-1}, []int{-1})))
		fmt.Println(`Input: preorder = [1,4,2,3], inorder = [1,2,3,4]`)
		fmt.Println(`Output:`, STreeNode(buildTree([]int{1, 4, 2, 3}, []int{1, 2, 3, 4})))
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}

	pv := preorder[0]
	pi := 0
	for pi < n && inorder[pi] != pv {
		pi++
	}

	res := new(TreeNode)
	res.Val = pv
	res.Left = buildTree(preorder[1:], inorder[:pi])
	res.Right = buildTree(preorder[1+pi:], inorder[pi+1:])
	return res
}
