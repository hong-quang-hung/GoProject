package contest

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference:
func Leetcode_Replace_Value_In_Tree() {
	fmt.Println("Input: root = [5,4,9,1,10,null,7]")
	fmt.Println("Output:", replaceValueInTree(types.LazyNodeAll(5, types.LazyNodeValue(4, 1, 10), types.LazyNodeAll(9, nil, types.LazyNode(7)))))
	fmt.Println("Input: root = [3,1,2]")
	fmt.Println("Output:", replaceValueInTree(types.LazyNodeValue(3, 1, 2)))
}

func replaceValueInTree(root *types.TreeNode) *types.TreeNode {
	sum := make(map[int]int)
	replaceValueInTreeSolve(0, root, sum, 0)
	fmt.Println(sum)
	root.Println()
	return root
}

func replaceValueInTreeSolve(other int, root *types.TreeNode, sum map[int]int, depth int) {
	if root == nil {
		return
	}

	sum[depth] += root.Val
	if root.Left != nil {
		if root.Right != nil {
			replaceValueInTreeSolve(root.Right.Val, root.Left, sum, depth+1)
		} else {
			replaceValueInTreeSolve(0, root.Left, sum, depth+1)
		}
	}

	if root.Right != nil {
		if root.Left != nil {
			replaceValueInTreeSolve(root.Left.Val, root.Right, sum, depth+1)
		} else {
			replaceValueInTreeSolve(0, root.Right, sum, depth+1)
		}
	}
	root.Val = sum[depth] - root.Val - other
}
