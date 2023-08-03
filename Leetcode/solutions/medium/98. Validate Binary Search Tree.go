package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/validate-binary-search-tree/
func init() {
	Solutions[98] = func() {
		fmt.Println("Input: root = [2,1,3]")
		fmt.Println("Output:", isValidBST(S2TreeNode("[2,1,3]")))
		fmt.Println("Input: root = [5,1,4,null,null,3,6]")
		fmt.Println("Output:", isValidBST(S2TreeNode("[5,1,4,null,null,3,6]")))
		fmt.Println("Input: root = [5,4,6,null,null,3,7]")
		fmt.Println("Output:", isValidBST(S2TreeNode("[5,4,6,null,null,3,7]")))
	}
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTCheck(root, math.MaxInt, math.MinInt)
}

func isValidBSTCheck(root *TreeNode, a int, b int) bool {
	if root == nil {
		return true
	}

	if root.Val >= a || root.Val <= b {
		return false
	}
	return isValidBSTCheck(root.Left, root.Val, b) && isValidBSTCheck(root.Right, a, root.Val)
}
