package easy

import "fmt"

// Reference:https://leetcode.com/problems/range-sum-of-bst/?envType=daily-question&envId=2024-01-08
func init() {
	Solutions[938] = func() {
		fmt.Println(`Input: root = [10,5,15,3,7,null,18], low = 7, high = 15`)
		fmt.Println(`Output:`, rangeSumBST(S2TreeNode("[10,5,15,3,7,null,18]"), 7, 15))
		fmt.Println(`Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10`)
		fmt.Println(`Output:`, rangeSumBST(S2TreeNode("[10,5,15,3,7,13,18,1,null,6]"), 6, 10))
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val >= low && root.Val <= high {
		res += root.Val
	}

	res += rangeSumBST(root.Left, low, high)
	res += rangeSumBST(root.Right, low, high)
	return res
}
