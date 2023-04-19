package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
func Leetcode_Longest_ZigZag() {
	fmt.Println("Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]")
	fmt.Println("Output:", longestZigZag(utils.S2TreeNode("[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]")))
	fmt.Println("Input: root = [1,1,1,null,1,null,null,1,1,null,1]")
	fmt.Println("Output:", longestZigZag(utils.S2TreeNode("[1,1,1,null,1,null,null,1,1,null,1]")))
}

func longestZigZag(root *types.TreeNode) int {
	res := 0
	longestZigZagDFS(root, &res, 0, true)
	longestZigZagDFS(root, &res, 0, false)
	return res
}

func longestZigZagDFS(root *types.TreeNode, maxPath *int, path int, isLeft bool) {
	if root == nil {
		return
	}

	*maxPath = max(*maxPath, path)
	if isLeft {
		longestZigZagDFS(root.Right, maxPath, 0, true)
		longestZigZagDFS(root.Right, maxPath, path+1, false)
	} else {
		longestZigZagDFS(root.Left, maxPath, path+1, true)
		longestZigZagDFS(root.Left, maxPath, 0, false)
	}
}
