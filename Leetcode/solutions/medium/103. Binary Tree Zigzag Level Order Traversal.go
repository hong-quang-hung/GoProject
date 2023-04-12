package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func Leetcode_Zigzag_Level_Order() {
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", zigzagLevelOrder(types.LazyNodeAll(3, types.LazyNode(9), types.LazyNodeValue(20, 15, 7))))
}

func zigzagLevelOrder(root *types.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	nums, isNodeLeft := 1, false
	queue := []*types.TreeNode{root}
	nodes := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodes = append(nodes, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if len(nodes) == nums {
			if isNodeLeft {
				for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
					nodes[i], nodes[j] = nodes[j], nodes[i]
				}
			}
			res = append(res, nodes)
			nodes = nil
			nums = len(queue)
			isNodeLeft = !isNodeLeft
		}
	}
	return res
}
