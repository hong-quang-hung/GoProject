package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/find-bottom-left-tree-value/
func init() {
	Solutions[513] = func() {
		fmt.Println(`Input: root = [2,1,3]`)
		fmt.Println(`Output:`, findBottomLeftValue(S2TreeNode("[2,1,3]")))
		fmt.Println(`Input: root = [1,2,3,4,null,5,6,null,null,7]`)
		fmt.Println(`Output:`, findBottomLeftValue(S2TreeNode("[1,2,3,4,null,5,6,null,null,7]")))
	}
}

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	res := root
	for len(queue) > 0 {
		res = queue[0]
		queue = queue[1:]
		if res.Right != nil {
			queue = append(queue, res.Right)
		}
		if res.Left != nil {
			queue = append(queue, res.Left)
		}
	}
	return res.Val
}
