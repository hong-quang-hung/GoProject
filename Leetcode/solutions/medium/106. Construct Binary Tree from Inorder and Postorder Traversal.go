package medium

import "fmt"

// Reference: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func init() {
	Solutions[106] = func() {
		fmt.Println(`Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]`)
		fmt.Println(`Output:`, STreeNode(buildTree_ii([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})))
	}
}

func buildTree_ii(inorder []int, postorder []int) *TreeNode {
	return buildTreeConstruct(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func buildTreeConstruct(inorder []int, postorder []int, i1 int, i2 int, p1 int, p2 int) *TreeNode {
	mid, rootVal := i1, postorder[p2]
	for mid <= i2 && inorder[mid] != rootVal {
		mid++
	}

	root := &TreeNode{Val: rootVal}
	if mid-i1 == 1 {
		root.Left = &TreeNode{Val: inorder[i1]}
	} else if mid-i1 > 1 {
		root.Left = buildTreeConstruct(inorder, postorder, i1, mid-1, p1, p1+mid-i1-1)
	}

	if i2-mid == 1 {
		root.Right = &TreeNode{Val: postorder[p2-1]}
	} else if i2-mid > 1 {
		root.Right = buildTreeConstruct(inorder, postorder, mid+1, i2, p1+mid-i1, p2-1)
	}
	return root
}
