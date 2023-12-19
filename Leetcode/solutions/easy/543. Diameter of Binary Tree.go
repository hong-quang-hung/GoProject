package easy

import "fmt"

// Reference: https://leetcode.com/problems/diameter-of-binary-tree/
func init() {
	Solutions[543] = func() {
		fmt.Println(`Input: root = [1,2,3,4,5]`)
		fmt.Println(`Output:`, diameterOfBinaryTree(S2TreeNode(`[1,2,3,4,5]`)))
		fmt.Println(`Input: root = [1,2]`)
		fmt.Println(`Output:`, diameterOfBinaryTree(S2TreeNode(`[1,2]`)))
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	diameterOfBinaryTreeDFS(root, &res)
	return res
}

func diameterOfBinaryTreeDFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTreeDFS(root.Left, res)
	right := diameterOfBinaryTreeDFS(root.Right, res)
	*res = max(*res, left+right)
	return max(left, right) + 1
}
