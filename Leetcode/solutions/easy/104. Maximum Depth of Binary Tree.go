package easy

import "fmt"

// Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/
func LeetCode_Max_Depth() {
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", maxDepth(S2TreeNode("[3,9,20,null,null,15,7]")))
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output:", maxDepth(S2TreeNode("[1,null,2]")))
}

func maxDepth(root *TreeNode) int {
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
