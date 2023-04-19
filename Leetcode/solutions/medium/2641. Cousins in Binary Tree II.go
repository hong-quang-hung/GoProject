package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/cousins-in-binary-tree-ii/
func Leetcode_Replace_Value_In_Tree() {
	fmt.Println("Input: root = [5,4,9,1,10,null,7]")
	fmt.Println("Output:", utils.STreeNode(replaceValueInTree(types.LazyNodeAll(5, types.LazyNodeValue(4, 1, 10), types.LazyNodeAll(9, nil, types.LazyNode(7))))))
	fmt.Println("Input: root = [3,1,2]")
	fmt.Println("Output:", utils.STreeNode(replaceValueInTree(types.LazyNodeValue(3, 1, 2))))
}

func replaceValueInTree(root *types.TreeNode) *types.TreeNode {
	sum := make(map[int]int)
	replaceValueInTreeSum(root, sum, 0)
	replaceValueInTreeSolved(root, sum, 0, 0)
	return root
}

func replaceValueInTreeSum(root *types.TreeNode, sum map[int]int, depth int) {
	if root == nil {
		return
	}

	sum[depth] += root.Val
	if root.Left != nil {
		replaceValueInTreeSum(root.Left, sum, depth+1)
	}

	if root.Right != nil {
		replaceValueInTreeSum(root.Right, sum, depth+1)
	}
}

func replaceValueInTreeSolved(root *types.TreeNode, sum map[int]int, depth int, sameParentNode int) {
	if root == nil {
		return
	}

	leftNil, rightNil, leftVal, rightVal := root.Left != nil, root.Right != nil, 0, 0
	if leftNil {
		leftVal = root.Left.Val
	}

	if rightNil {
		rightVal = root.Right.Val
	}

	if leftNil {
		replaceValueInTreeSolved(root.Left, sum, depth+1, rightVal)
	}

	if rightNil {
		replaceValueInTreeSolved(root.Right, sum, depth+1, leftVal)
	}

	root.Val = sum[depth] - root.Val - sameParentNode
}
