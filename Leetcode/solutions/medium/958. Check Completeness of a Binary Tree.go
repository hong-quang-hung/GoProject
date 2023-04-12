package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/check-completeness-of-a-binary-tree/
func Leetcode_Is_Complete_Tree() {
	fmt.Println("Input: root = [1,2,3,5,null,7,8]")
	fmt.Println("Output:", isCompleteTree(types.LazyNodeAll(1, types.LazyNodeAll(2, types.LazyNode(5), nil), types.LazyNodeValue(3, 7, 8))))
}

func isCompleteTree(root *types.TreeNode) bool {
	queue := make([]*types.TreeNode, 0)
	queue = append(queue, root)
	check := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			check = true
		} else {
			if check {
				return false
			}
			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}
