package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
func init() {
	Solutions[2265] = func() {
		fmt.Println("Input: root = [4,8,5,0,1,null,6]")
		fmt.Println("Output:", averageOfSubtree(S2TreeNode("[4,8,5,0,1,null,6]")))
		fmt.Println("Input: root = [1]")
		fmt.Println("Output:", averageOfSubtree(S2TreeNode("[1]")))
	}
}

func averageOfSubtree(root *TreeNode) int {
	res := 0
	var f func(node *TreeNode) (int, int)
	f = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		l1, l2 := f(node.Left)
		r1, r2 := f(node.Right)

		nodeSum := l1 + r1 + node.Val
		nodeCount := l2 + r2 + 1

		if node.Val == nodeSum/nodeCount {
			res++
		}
		return nodeSum, nodeCount
	}

	f(root)
	return res
}
