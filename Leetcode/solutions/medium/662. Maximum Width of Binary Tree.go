package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/maximum-width-of-binary-tree/
func Leetcode_Width_Of_Binary_Tree() {
	fmt.Println("Input: root = [1,3,2,5,3,null,9]")
	fmt.Println("Output:", widthOfBinaryTree(utils.S2TreeNode("[1,3,2,5,3,null,9]")))
	fmt.Println("Input: root = [1,3,2,5,null,null,9,6,null,7]")
	fmt.Println("Output:", widthOfBinaryTree(utils.S2TreeNode("[1,3,2,5,null,null,9,6,null,7]")))
	fmt.Println("Input: root = [1,3,2,5]")
	fmt.Println("Output:", widthOfBinaryTree(utils.S2TreeNode("[1,3,2,5]")))
}

func widthOfBinaryTree(root *types.TreeNode) int {
	res := 0
	fmt.Println(utils.STreeNode(root))
	return res
}
