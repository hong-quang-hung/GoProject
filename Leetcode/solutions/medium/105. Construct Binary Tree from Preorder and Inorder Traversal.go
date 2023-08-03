package medium

import "fmt"

// Reference: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func init() {
	Solutions[105] = func() {
		// fmt.Println("Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]")
		// fmt.Println("Output:", STreeNode(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
		// fmt.Println("Input: preorder = [-1], inorder = [-1]")
		// fmt.Println("Output:", STreeNode(buildTree([]int{-1}, []int{-1})))
		fmt.Println("Input: preorder = [1,2,3], inorder = [2,3,1]")
		fmt.Println("Output:", STreeNode(buildTree([]int{1, 2, 3}, []int{2, 3, 1})))
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeSolved(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTreeSolved(preorder []int, inorder []int, p1 int, p2 int, i1 int, i2 int) *TreeNode {
	mid, rootVal := i1, preorder[p1]
	for mid <= i2 && inorder[mid] != rootVal {
		mid++
	}

	root := &TreeNode{Val: rootVal}
	if mid-i1 == 1 {
		root.Left = &TreeNode{Val: inorder[i1]}
	} else if mid-i1 > 1 {
		root.Left = buildTreeSolved(preorder, inorder, p1+mid-i1-1, p2-1, i1, mid-1)
	}

	if i2-mid == 1 {
		root.Right = &TreeNode{Val: preorder[p2]}
	} else if i2-mid > 1 {
		root.Right = buildTreeSolved(preorder, inorder, p1+mid-i1+1, p2, mid+1, i2)
	}
	return root
}
