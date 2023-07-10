package medium

import "fmt"

func init() {
	Solutions[129] = Leetcode_Sum_Numbers
}

// Reference: https://leetcode.com/problems/sum-root-to-leaf-numbers/
func Leetcode_Sum_Numbers() {
	fmt.Println("Input: root = [1,2,3]")
	fmt.Println("Output:", sumNumbers(S2TreeNode("[1,2,3]")))
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if root.Left != nil {
		root.Left.Val += 10 * root.Val
	}
	if root.Right != nil {
		root.Right.Val += 10 * root.Val
	}
	return sumNumbers(root.Left) + sumNumbers(root.Right)
}
