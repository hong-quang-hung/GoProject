package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
func init() {
	Solutions[1372] = func() {
		fmt.Println(`Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]`)
		fmt.Println(`Output:`, longestZigZag(S2TreeNode(`[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]`)))
		fmt.Println(`Input: root = [1,1,1,null,1,null,null,1,1,null,1]`)
		fmt.Println(`Output:`, longestZigZag(S2TreeNode(`[1,1,1,null,1,null,null,1,1,null,1]`)))
	}
}

func longestZigZag(root *TreeNode) int {
	res := 0
	longestZigZagDFS(root, &res, 0, true)
	longestZigZagDFS(root, &res, 0, false)
	return res
}

func longestZigZagDFS(root *TreeNode, maxPath *int, path int, isLeft bool) {
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
