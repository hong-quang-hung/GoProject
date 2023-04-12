package easy

import (
	"fmt"
	"math"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/minimum-distance-between-bst-nodes
func Leetcode_Min_Diff_In_BST() {
	fmt.Println("Input: root = [4,2,6,1,3]")
	fmt.Println("Output:", minDiffInBST(types.LazyNodeAll(4, types.LazyNodeValue(2, 1, 3), types.LazyNode(6))))
}

func minDiffInBST(root *types.TreeNode) int {
	minDiff, preValue := math.MaxInt, -1
	inorderTraversal(root, &preValue, &minDiff)
	return minDiff
}

func inorderTraversal(root *types.TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, preValue, minDiff)

	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val
	inorderTraversal(root.Right, preValue, minDiff)
}
