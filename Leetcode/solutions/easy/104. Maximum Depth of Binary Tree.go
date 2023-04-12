package easy

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/
func LeetCode_Max_Depth() {
	var root *types.TreeNode

	root = types.LazyNodeAll(3, types.LazyNode(9), types.LazyNodeValue(20, 15, 7))
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", maxDepth(root))
	fmt.Println()

	root = types.LazyNodeAll(1, nil, types.LazyNode(9))
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output:", maxDepth(root))
}

func maxDepth(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	var l int = maxDepth(root.Left)
	var r int = maxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}
