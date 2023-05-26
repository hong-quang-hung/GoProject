package easy

import "fmt"

// Reference: https://leetcode.com/problems/path-sum/
func Leetcode_Has_Path_Sum() {
	fmt.Println("Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22")
	fmt.Println("Output:", hasPathSum(S2TreeNode("[5,4,8,11,null,13,4,7,2,null,null,null,1]"), 22))
	fmt.Println("Input: root = [1,2,3], targetSum = 5")
	fmt.Println("Output:", hasPathSum(S2TreeNode("[1,2,3]"), 5))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
