package medium

import "fmt"

// Reference: https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
func init() {
	Solutions[1457] = func() {
		fmt.Println(`Input: root = [2,3,1,3,1,null,1]`)
		fmt.Println(`Output:`, pseudoPalindromicPaths(S2TreeNode("[2,3,1,3,1,null,1]")))
		fmt.Println(`Input: root = [2,1,1,1,3,null,null,null,null,null,1]`)
		fmt.Println(`Output:`, pseudoPalindromicPaths(S2TreeNode("[2,1,1,1,3,null,null,null,null,null,1]")))
	}
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var f func(*TreeNode, int) int
	f = func(node *TreeNode, mask int) int {
		if node == nil {
			return 0
		}
		mask ^= 1 << node.Val
		if node.Left == node.Right {
			if mask&(mask-1) == 0 {
				return 1
			}
			return 0
		}

		left := f(node.Left, mask)
		right := f(node.Right, mask)
		return left + right
	}
	return f(root, 0)
}
