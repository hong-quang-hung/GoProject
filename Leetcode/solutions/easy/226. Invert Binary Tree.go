package easy

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/invert-binary-tree/
func Leetcode_Invert_Tree() {
	fmt.Println("Input: root = [4,2,7,1,3,6,9]")
	fmt.Println("Output:", invertTree(types.LazyNodeAll(4, types.LazyNodeValue(2, 1, 3), types.LazyNodeValue(7, 6, 9))))
}

func invertTree(root *types.TreeNode) *types.TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
