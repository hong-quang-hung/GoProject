package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/binary-tree-maximum-path-sum/
func init() {
	Solutions[124] = func() {
		fmt.Println("Input: root = [1,2,3]")
		fmt.Println("Output:", maxPathSum(S2TreeNode("[1,2,3]")))
		fmt.Println("Input: root = [-10,9,20,null,null,15,7]")
		fmt.Println("Output:", maxPathSum(S2TreeNode("[-10,9,20,null,null,15,7]")))
	}
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	maxPathSumDFS(root, &res)
	return res
}

func maxPathSumDFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := maxPathSumDFS(root.Left, res)
	right := maxPathSumDFS(root.Right, res)
	*res = max(*res, left+right+root.Val)
	return max(max(left+root.Val, right+root.Val), 0)
}
