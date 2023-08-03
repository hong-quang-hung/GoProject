package medium

import "fmt"

// Reference: https://leetcode.com/problems/binary-tree-level-order-traversal/
func init() {
	Solutions[102] = func() {
		fmt.Println("Input: root = [3,9,20,null,null,15,7]")
		fmt.Println("Output:", levelOrder(S2TreeNode("[3,9,20,null,null,15,7]")))
		fmt.Println("Input: root = [1]")
		fmt.Println("Output:", levelOrder(S2TreeNode("[1]")))
	}
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		child := []int{}
		temp := []*TreeNode{}
		for _, node := range queue {
			child = append(child, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		res = append(res, child)
		queue = temp
	}
	return res
}
