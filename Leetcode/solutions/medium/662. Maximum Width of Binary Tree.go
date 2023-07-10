package medium

import "fmt"

func init() {
	Solutions[662] = Leetcode_Width_Of_Binary_Tree
}

// Reference: https://leetcode.com/problems/maximum-width-of-binary-tree/
func Leetcode_Width_Of_Binary_Tree() {
	fmt.Println("Input: root = [1,3,2,5,3,null,9]")
	fmt.Println("Output:", widthOfBinaryTree(S2TreeNode("[1,3,2,5,3,null,9]")))
	fmt.Println("Input: root = [1,3,2,5,null,null,9,6,null,7]")
	fmt.Println("Output:", widthOfBinaryTree(S2TreeNode("[1,3,2,5,null,null,9,6,null,7]")))
	fmt.Println("Input: root = [1,3,2,5]")
	fmt.Println("Output:", widthOfBinaryTree(S2TreeNode("[1,3,2,5]")))
}

func widthOfBinaryTree(root *TreeNode) int {
	m := make(map[int]int)
	return widthOfBinaryTreeDFS(root, m, 0, 1)
}

func widthOfBinaryTreeDFS(root *TreeNode, m map[int]int, level int, width int) int {
	if root == nil {
		return 0
	}

	if _, ok := m[level]; !ok {
		m[level] = width
	}

	cur := width - m[level] + 1
	left := widthOfBinaryTreeDFS(root.Left, m, level+1, width*2-1)
	right := widthOfBinaryTreeDFS(root.Right, m, level+1, width*2)
	return max(cur, max(left, right))
}
