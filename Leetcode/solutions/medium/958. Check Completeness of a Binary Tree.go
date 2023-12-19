package medium

import "fmt"

// Reference: https://leetcode.com/problems/check-completeness-of-a-binary-tree/
func init() {
	Solutions[958] = func() {
		fmt.Println(`Input: root = [1,2,3,5,null,7,8]`)
		fmt.Println(`Output:`, isCompleteTree(S2TreeNode(`[1,2,3,5,null,7,8]`)))
	}
}

func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
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
