package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-distance-between-bst-nodes
func init() {
	Solutions[783] = func() {
		fmt.Println(`Input: root = [4,2,6,1,3]`)
		fmt.Println(`Output:`, minDiffInBST(S2TreeNode(`[4,2,6,1,3]`)))
	}
}

func minDiffInBST(root *TreeNode) int {
	minDiff, preValue := math.MaxInt, -1
	minDiffInBSTSolved(root, &preValue, &minDiff)
	return minDiff
}

func minDiffInBSTSolved(root *TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}

	minDiffInBSTSolved(root.Left, preValue, minDiff)
	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val
	minDiffInBSTSolved(root.Right, preValue, minDiff)
}
