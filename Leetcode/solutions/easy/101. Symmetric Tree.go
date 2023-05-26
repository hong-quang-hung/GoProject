package easy

import "fmt"

// Reference: https://leetcode.com/problems/symmetric-tree/
func Leetcode_Is_Symmetric() {
	fmt.Println("Input: root = [1,2,2,3,4,4,3]")
	fmt.Println("Output:", isSymmetric(S2TreeNode("[1,2,2,3,4,4,3]")))
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricCheck(root.Left, root.Right)
}

func isSymmetricCheck(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return isSymmetricCheck(left.Left, right.Right) && isSymmetricCheck(left.Right, right.Left)
}
