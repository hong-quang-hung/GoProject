package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-good-nodes-in-binary-tree/
func Leetcode_Good_Nodes() {
	fmt.Println("Input: root = [3,1,4,3,null,1,5]")
	fmt.Println("Output:", goodNodes(S2TreeNode("[3,1,4,3,null,1,5]")))
	fmt.Println("Input: root = [3,3,null,4,2]")
	fmt.Println("Output:", goodNodes(S2TreeNode("[3,3,null,4,2]")))
	fmt.Println("Input: root = [1]")
	fmt.Println("Output:", goodNodes(S2TreeNode("[1]")))
}

func goodNodes(root *TreeNode) int {
	return goodNodesCount(root, -10_001)
}

func goodNodesCount(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	if root.Val >= maxVal {
		return 1 + goodNodesCount(root.Left, root.Val) + goodNodesCount(root.Right, root.Val)
	}
	return goodNodesCount(root.Left, maxVal) + goodNodesCount(root.Right, maxVal)
}
